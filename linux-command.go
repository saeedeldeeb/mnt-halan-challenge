package main

import (
	"os/exec"
)

type File struct {
	Name string
}

func (f File) Rename(s string) error {
	cmd := exec.Command("mv", f.Name, s)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
