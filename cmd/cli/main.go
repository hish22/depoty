package main

import (
	"depoty/cmd/cli/build"

	"golang.org/x/sys/windows"
)

func checkProcessPrivileges() bool {
	Token := windows.GetCurrentProcessToken()

	return Token.IsElevated()
}

func main() {
	// Check if user started the process with privileges or not.
	if !checkProcessPrivileges() {
		panic("Please run the application with administrator privileges," +
			" as Chocolatey requires elevated permissions to perform installations and other tasks.")
	}
	// Start the Process
	build.CreateCommands()
}
