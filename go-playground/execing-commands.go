package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	c := exec.Command("ls", "-a")
	var buffer strings.Builder
	c.Stdout = &buffer
	must(c.Run())
	fmt.Println(buffer.String())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
