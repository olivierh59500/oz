package ozinit

import (
	"bufio"
	"fmt"
	"github.com/kr/pty"
	"github.com/op/go-logging"
	"github.com/subgraph/oz"
	"github.com/subgraph/oz/fs"
	"github.com/subgraph/oz/ipc"
	"github.com/subgraph/oz/xpra"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

const SocketAddress = "/tmp/oz-init-control"

type initState struct {
	log     *logging.Logger
	profile *oz.Profile
	config  *oz.Config
	uid     int
	gid     int
	user    *user.User
	display int
	fs      *fs.Filesystem
}

// By convention oz-init writes log messages to stderr with a single character
// prefix indicating the logging level.  These messages are read one line at a time
// over a pipe by oz-daemon and translated into appropriate log events.
func createLogger() *logging.Logger {
	l := logging.MustGetLogger("oz-init")
	be := logging.NewLogBackend(os.Stderr, "", 0)
	f := logging.MustStringFormatter("%{level:.1s} %{message}")
	fbe := logging.NewBackendFormatter(be, f)
	logging.SetBackend(fbe)
	return l
}

func Main() {
	parseArgs().runInit()
}

func parseArgs() *initState {
	log := createLogger()
	getvar := func(name string) string {
		val := os.Getenv(name)
		if val == "" {
			log.Error("Error: missing required '%s' argument", name)
			os.Exit(1)
		}
		return val
	}
	pname := getvar("INIT_PROFILE")
	uidval := getvar("INIT_UID")
	dispval := os.Getenv("INIT_DISPLAY")

	var config *oz.Config
	config, err := oz.LoadConfig(oz.DefaultConfigPath)
	if err != nil {
		log.Info("Could not load config file (%s), using default config", oz.DefaultConfigPath)
		config = oz.NewDefaultConfig()
	}

	p, err := loadProfile(config.ProfileDir, pname)
	if err != nil {
		log.Error("Could not load profile %s: %v", pname, err)
		os.Exit(1)
	}
	uid, err := strconv.Atoi(uidval)
	if err != nil {
		log.Error("Could not parse INIT_UID argument (%s) into an integer: %v", uidval, err)
		os.Exit(1)
	}
	u, err := user.LookupId(uidval)
	if err != nil {
		log.Error("Failed to look up user with uid=%s: %v", uidval, err)
		os.Exit(1)
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		log.Error("Failed to parse gid value (%s) from user struct: %v", u.Gid, err)
		os.Exit(1)
	}
	display := 0
	if dispval != "" {
		d, err := strconv.Atoi(dispval)
		if err != nil {
			log.Error("Unable to parse display (%s) into an integer: %v", dispval, err)
			os.Exit(1)
		}
		display = d
	}

	return &initState{
		log:     log,
		config:  config,
		profile: p,
		uid:     uid,
		gid:     gid,
		user:    u,
		display: display,
		fs:      fs.NewFromProfile(p, u, log),
	}
}

func (st *initState) runInit() {
	st.log.Info("Starting oz-init for profile: %s", st.profile.Name)
	if syscall.Sethostname([]byte(st.profile.Name)) != nil {
		st.log.Error("Failed to set hostname to (%s)", st.profile.Name)
	}
	st.log.Info("Hostname set to (%s)", st.profile.Name)

	if err := st.fs.OzInit(); err != nil {
		st.log.Error("Error: setting up filesystem failed: %v\n", err)
		os.Exit(1)
	}
	if st.profile.XServer.Enabled {
		if st.display == 0 {
			st.log.Error("Cannot start xpra because no display number was passed to oz-init")
			os.Exit(1)
		}
		st.startXpraServer()
	}

	oz.ReapChildProcs(st.log, st.handleChildExit)

	s, err := ipc.NewServer(SocketAddress, messageFactory, st.log,
		handlePing,
		st.handleRunShell,
	)
	if err != nil {
		st.log.Error("NewServer failed: %v", err)
		os.Exit(1)
	}
	if err := os.Chown(SocketAddress, st.uid, st.gid); err != nil {
		st.log.Warning("Failed to chown oz-init control socket: %v", err)
	}

	if err := s.Run(); err != nil {
		st.log.Warning("MsgServer.Run() return err: %v", err)
	}
	st.log.Info("oz-init exiting...")
}

func (st *initState) startXpraServer() {
	workdir := st.fs.Xpra()
	if workdir == "" {
		st.log.Warning("Xpra work directory not set")
		return
	}
	xpra := xpra.NewServer(&st.profile.XServer, uint64(st.display), workdir)
	p, err := xpra.Process.StderrPipe()
	if err != nil {
		st.log.Warning("Error creating stderr pipe for xpra output: %v", err)
		os.Exit(1)
	}
	go st.readXpraOutput(p)
	xpra.Process.Env = []string{
		"HOME=" + st.user.HomeDir,
	}
	xpra.Process.SysProcAttr = &syscall.SysProcAttr{}
	xpra.Process.SysProcAttr.Credential = &syscall.Credential{
		Uid: uint32(st.uid),
		Gid: uint32(st.gid),
	}
	st.log.Info("Starting xpra server")
	if err := xpra.Process.Start(); err != nil {
		st.log.Warning("Failed to start xpra server: %v", err)
	}
}

func (st *initState) readXpraOutput(r io.ReadCloser) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		if len(line) > 0 {
			if strings.Contains(line, "xpra is ready.") {
				os.Stderr.WriteString("XPRA READY\n")
				if !st.config.LogXpra {
					r.Close()
					return
				}
			}
			if st.config.LogXpra {
				st.log.Debug("(xpra) %s", line)
			}
		}
	}
}

func loadProfile(dir, name string) (*oz.Profile, error) {
	ps, err := oz.LoadProfiles(dir)
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		if name == p.Name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("no profile named '%s'", name)
}

func handlePing(ping *PingMsg, msg *ipc.Message) error {
	return msg.Respond(&PingMsg{Data: ping.Data})
}

func (st *initState) handleRunShell(rs *RunShellMsg, msg *ipc.Message) error {
	if msg.Ucred == nil {
		return msg.Respond(&ErrorMsg{"No credentials received for RunShell command"})
	}
	if msg.Ucred.Uid == 0 || msg.Ucred.Gid == 0 && !st.config.AllowRootShell {
		return msg.Respond(&ErrorMsg{"Cannot open shell because allowRootShell is disabled"})
	}
	st.log.Info("Starting shell with uid = %d, gid = %d", msg.Ucred.Uid, msg.Ucred.Gid)
	cmd := exec.Command(st.config.ShellPath, "-i")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid: msg.Ucred.Uid,
		Gid: msg.Ucred.Gid,
	}
	if rs.Term != "" {
		cmd.Env = append(cmd.Env, "TERM="+rs.Term)
	}
	cmd.Env = append(cmd.Env, "PATH=/usr/bin:/bin")
	cmd.Env = append(cmd.Env, fmt.Sprintf("PS1=[%s] $ ", st.profile.Name))
	st.log.Info("Executing shell...")
	f, err := ptyStart(cmd)
	defer f.Close()
	if err != nil {
		return msg.Respond(&ErrorMsg{err.Error()})
	}
	err = msg.Respond(&OkMsg{}, int(f.Fd()))
	return err
}

func ptyStart(c *exec.Cmd) (ptty *os.File, err error) {
	ptty, tty, err := pty.Open()
	if err != nil {
		return nil, err
	}
	defer tty.Close()
	c.Stdin = tty
	c.Stdout = tty
	c.Stderr = tty
	if c.SysProcAttr == nil {
		c.SysProcAttr = &syscall.SysProcAttr{}
	}
	c.SysProcAttr.Setctty = true
	c.SysProcAttr.Setsid = true
	if err := c.Start(); err != nil {
		ptty.Close()
		return nil, err
	}
	return ptty, nil
}

func (is *initState) handleChildExit(pid int, wstatus syscall.WaitStatus) {
	is.log.Debug("Child process pid=%d exited with status %d", pid, wstatus.ExitStatus())
}
