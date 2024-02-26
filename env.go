package goborg

type BorgEnv struct {
	vars []string
}

func NewBorgEnvironment(variables ...string) BorgEnv {
	return BorgEnv{
		vars: variables,
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
