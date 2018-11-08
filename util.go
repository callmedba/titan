package thanos

import (
	"fmt"
	"sync/atomic"

	"github.com/twinj/uuid"

	"go.uber.org/zap"
)

// Version information.
var (
	ReleaseVersion = "None"
	BuildTS        = "None"
	GitHash        = "None"
	GitBranch      = "None"
	GitLog         = "None"
	GolangVersion  = "None"
	ConfigFile     = "None"
)

func logVersionInfo() {
	zap.L().Info("Welcome to Titan.")
	zap.L().Info("Server info", zap.String("Release Version", ReleaseVersion))
	zap.L().Info("Server info", zap.String("Git Commit Hash", GitHash))
	zap.L().Info("Server info", zap.String("Git Commit Log", GitLog))
	zap.L().Info("Server info", zap.String("Git Branch", GitBranch))
	zap.L().Info("Server info", zap.String("UTC Build Time", BuildTS))
	zap.L().Info("Server info", zap.String("Golang compiler Version", GolangVersion))
}

// PrintVersionInfo print the server version info
func PrintVersionInfo() {
	fmt.Println("Welcome to Titan.")
	fmt.Println("Release Version: ", ReleaseVersion)
	fmt.Println("Git Commit Hash: ", GitHash)
	fmt.Println("Git Commit Log: ", GitLog)
	fmt.Println("Git Branch: ", GitBranch)
	fmt.Println("UTC Build Time:  ", BuildTS)
	fmt.Println("Golang compiler Version: ", GolangVersion)
}

func GetClientID() func() int64 {
	var id int64 = 1
	return func() int64 {
		return atomic.AddInt64(&id, 1)
	}
}

func GenerateTraceID() string { return uuid.NewV4().String() }