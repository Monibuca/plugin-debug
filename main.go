package debug

import (
	"io"
	"net/http"
	"net/http/pprof"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	. "m7s.live/engine/v4"
)

type DebugConfig struct {
}

type WriteToFile struct {
	header http.Header
	io.Writer
}

func (w *WriteToFile) Header() http.Header {
	// return w.w.Header()
	return w.header
}

// func (w *WriteToFile) Write(p []byte) (int, error) {
// 	// w.w.Write(p)
// 	return w.Writer.Write(p)
// }
func (w *WriteToFile) WriteHeader(statusCode int) {
	// w.w.WriteHeader(statusCode)
}
func (p *DebugConfig) OnEvent(event any) {
	switch event.(type) {
	case FirstConfig:

	}
}

func (p *DebugConfig) Trace(w http.ResponseWriter, r *http.Request) {
	pprof.Trace(w, r)
}

func (p *DebugConfig) Profile(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("cpu.profile", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		plugin.Error("profile", zap.Error(err))
	}
	pprof.Profile(&WriteToFile{make(http.Header), file}, r)
	file.Close()
	if r.Host == "" {
		r.Host = "localhost"
	}
	w.Write([]byte(strings.Join([]string{"go", "tool", "pprof", "-http :6060", ExecPath, filepath.Join(ExecDir, "cpu.profile")}, " ")))
}

func (p *DebugConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}

var plugin = InstallPlugin(&DebugConfig{})
