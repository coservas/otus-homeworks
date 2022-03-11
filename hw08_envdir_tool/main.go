package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("args is empty\n")
		return
	}

	envs, err := ReadDir(args[0])
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	os.Exit(RunCmd(args[1:], envs))
}
