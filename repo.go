package goborg

import "errors"

const (
	Repokey             string = "repokey"
	RepokeyBlake2              = "repokey-blake2"
	KeyfileBlake2              = "keyfile-blake2"
	Authenticated              = "authenticated"
	AuthenticatedBlake2        = "authenticated-blake2"
	None                       = "none"
)

func (b *BorgEnv) InitRepo(path string, encryption string, opts ...string) error {
	args := append([]string{path, "--encryption=" + encryption}, opts...)
	output, err := b.Exec("init", args...)
	if err != nil {
		return err
	}

	if len(output) == 0 {
		return nil
	}

	if output["levelname"] == "ERROR" {
		return errors.New(output["message"].(string))
	} else {
		return errors.New("error initializing the repo")
	}
}
