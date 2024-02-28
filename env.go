package goborg

type BorgEnv struct {
	repoPath string
	vars     []string
}

func NewBorgEnvironment(repoPath string, variables ...string) BorgEnv {
	return BorgEnv{
		repoPath: repoPath,
		vars:     variables,
	}
}

func WithPassphrase(passphrase string) string {
	return "BORG_PASSPHRASE=" + passphrase
}

func WithRSH(command string) string {
	return "BORG_RSH=" + command
}

func WithNewVariable(key, value string) string {
	return key + "=" + value
}
