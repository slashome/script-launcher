package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/slashome/script-launcher/print"
)

const version string = "0.0.1"

func main() {
	print.SetVersion(version)
	// get arguments
	args := os.Args[1:]
	// test if scripts path as been given
	if len(args) == 0 {
		print.Error("Please entrer the scripts path as a first argument")
		os.Exit(0)
	}
	// get script path
	path := args[0]
	// test if folder exists
	_, err := os.Stat(path)
	if err != nil {
		print.Error("Folder " + path + " does not exist")
		os.Exit(0)
	}
	// test if options have been given
	if len(args) < 2 {
		print.List(path, "")
		os.Exit(0)
	}
	optionsStr := args[1]
	options := strings.Split(optionsStr, ":")
	group := options[0]
	dir, err := os.Open(path + "/" + group)
	if err != nil {
		print.ErrorWithUsageAndList("Unable to find group "+"\""+group+"\"", path, "")
		os.Exit(0)
	}
	defer dir.Close()

	if len(options) < 2 {
		print.List(path, group)
		os.Exit(0)
	}

	command := options[1]
	file, err := os.Open(path + "/" + group + "/" + command)
	if err != nil {
		print.ErrorWithUsageAndList("Unable to find command "+"\""+command+"\""+" for group "+"\""+group+"\"", path, group)
		os.Exit(0)
		os.Exit(0)
	}
	defer file.Close()

	cmd := exec.Command(path + "/" + group + "/" + command)
	fmt.Println("Running command and waiting for it to finish...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
