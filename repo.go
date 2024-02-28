package goborg

import (
	"errors"
)

type EncryptionMode string

const (
	Repokey             EncryptionMode = "repokey"
	RepokeyBlake2                      = "repokey-blake2"
	KeyfileBlake2                      = "keyfile-blake2"
	Authenticated                      = "authenticated"
	AuthenticatedBlake2                = "authenticated-blake2"
	None                               = "none"
)

func (b *BorgEnv) InitRepo(encryption EncryptionMode, opts ...string) error {
	args := append([]string{b.repoPath, "--encryption=" + string(encryption)}, opts...)
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
