package goborg

import "errors"

func (b *BorgEnv) Create(name, archivePath string, opts ...string) error {
	args := append([]string{b.repoPath + "::" + name, archivePath}, opts...)
	output, err := b.Exec("create", args...)
	if err != nil {
		return err
	}

	if len(output) == 0 {
		return nil
	}

	if output["levelname"] == "ERROR" {
		return errors.New(output["message"].(string))
	} else {
		return errors.New("error creating the archive")
	}
}

func (b *BorgEnv) Extract(path string) error {

}
