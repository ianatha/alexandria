// !build darwin,linux
package ostools

import (
	"fmt"
	"os"
	"os/exec"
)

func IsAdmin() bool {
	return os.Geteuid() == 0
}

const SUDO = "/usr/bin/sudo"

func runSelfWithSudo() error {
	fmt.Printf("elevating privileges by running sudo %v\n", os.Args)
	cmd := exec.Command(SUDO, os.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func EnsureAdministratorRights() {
	if !IsAdmin() {
		runSelfWithSudo()
		os.Exit(0)
	}
}
