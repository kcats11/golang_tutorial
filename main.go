package main

import (
	"fmt"

	"example.com/app/app/models"
	"example.com/app/app/models/controllers"
)

// func main() {
// 	cmdArgs := os.Args[1:]
// 	requiredLsArgs := 2
// 	requiredCpArgs := 3

// 	switch cmdArgs[0] {
// 	case "ls":
// 		if len(cmdArgs) >= requiredLsArgs {
// 			linuxcmd.Ls(cmdArgs[1])
// 		} else {
// 			linuxcmd.Ls("")
// 		}
// 	case "cp":
// 		if len(cmdArgs) >= requiredCpArgs {
// 			linuxcmd.Cp(cmdArgs[1],cmdArgs[2])
// 		}
// 	default:
// 		fmt.Println("command not found")
// 	}
// }

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}