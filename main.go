package main

import (
	"example.com/app/linuxcmd"
	"fmt"
	"os"
)

// change pointer
func main() {
	cmdArgs := os.Args[1:]
	requiredLsArgs := 2
	requiredCpArgs := 3
	
	switch cmdArgs[0] {
	case "ls":
		if len(cmdArgs) >= requiredLsArgs {
			linuxcmd.Ls(cmdArgs[1])
		} else {
			linuxcmd.Ls("")
		}
	case "cp":
		if len(cmdArgs) >= requiredCpArgs {
			linuxcmd.Cp(cmdArgs[1],cmdArgs[2])
		}
	default:
		fmt.Println("command not found")
	}
}
