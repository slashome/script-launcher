package main

import (
	"github.com/slashome/script-launcher/print"
	"os"
)

const version string = "0.0.1"

func main() {
	print.SetVersion(version)
	// get arguments
	args := os.Args[1:]
	// test if scripts path as been given
	if len(args) == 0 {
		print.Error("Please entrer the scripts path as a first argument", false)
		os.Exit(0)
	}
	// get script path
	path := args[0]
	// test if folder exists
	_, err := os.Stat(path)
	if err != nil {
		print.Error("Folder "+path+" does not exist", false)
		os.Exit(0)
	}
	// test if options have been given
	if len(args) < 2 {
		print.List(path)
		os.Exit(0)
	}
	// group := args[1]
	// action := args[2]
}
