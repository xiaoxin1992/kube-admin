package version

import "fmt"

var (
	GitTag    string
	GitCommit string
	GitBranch string
	BuildTime string
	GoVersion string
)

func FullVersion() string {
	version := fmt.Sprintf("Version   : %s\nBuild Time: %s\nGit Branch: %s\nGit Commit: %s\nGo Version: %s\n",
		GitTag, BuildTime, GitBranch, GitCommit, GoVersion)
	return version
}

func AppName() string {
	return "kube-admin"
}

func AppDescribe() string {
	return "kube-admin服务"
}
