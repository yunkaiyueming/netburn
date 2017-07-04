package analyze

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

type AnalyzeModel struct{}

var DefaultAnalyzeModel = new(AnalyzeModel)

func (a *AnalyzeModel) GetSysInfo() {
	runtime.CPUProfile()
}

func (a *AnalyzeModel) GoEnv() map[string]interface{} {
	host_name, _ := os.Hostname()
	groups, _ := os.Getgroups()
	pwd, _ := os.Getwd()

	return map[string]interface{}{
		"GOARCH":    runtime.GOARCH,
		"GOOS":      runtime.GOOS,
		"GOROOT":    runtime.GOROOT(),
		"GOPATH":    os.Getenv("GOPATH"),
		"Compiler":  runtime.Compiler,
		"GOVERSION": runtime.Version(),
		"CPU":       runtime.NumCPU(),
		"GOROUTINE": runtime.NumGoroutine(),
		"HOST_NAME": host_name,
		"PAGE_SIZE": os.Getpagesize(),
		"ENVS":      os.Environ(),
		"UID":       os.Getuid(),
		"Getgid":    os.Getgid(),
		"groups":    groups,
		"GROUP_ID":  os.Getegid(),
		"PPID":      os.Getppid(),
		"PID":       os.Getpid(),
		"pwd":       pwd,
	}
}

func (a *AnalyzeModel) Pprof() {
	pfs := pprof.Profiles()
	pprof.StartCPUProfile(os.Stdout)
	for _, pf := range pfs {
		fmt.Println("====", pf.Name(), pf.Count())
	}
	pprof.StopCPUProfile()
}
