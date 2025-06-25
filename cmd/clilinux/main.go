//go:build linux

package main

import "depoty/cmd/clilinux/build"

func main() {
	// Start the Process
	build.CreateCommands()
}
