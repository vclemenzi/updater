package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gitarchived/updater/data"
)

func GetLastCommit(r data.Repository, h data.Host) (string, error) {
	url := fmt.Sprintf("%s%s/%s.git", h.Prefix, r.Owner, r.Name)
	lsRemoteCmd := exec.Command("git", "ls-remote", url, "HEAD")
	lsRemoteCmd.Env = append(lsRemoteCmd.Env, "GIT_TERMINAL_PROMPT=0")

	out, err := lsRemoteCmd.Output()

	if err != nil {
		return "", err
	}

	commit := strings.Fields(string(out))[0]

	return commit, nil
}
