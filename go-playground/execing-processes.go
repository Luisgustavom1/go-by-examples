package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	bin, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()
	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
