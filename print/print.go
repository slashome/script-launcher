package print

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/fatih/color"
)

var version string = "0";

func SetVersion(v string) {
	version = v
}

func Error(message string, listPath bool) {
	Header()
	errorMessage(message)
	Usage()
}

func errorMessage(message string) {
	fmt.Println(color.RedString(message))
}

func Header() {
	fmt.Println("Script Launcher", color.GreenString(version))
}

func Usage() {
	fmt.Println("")
	fmt.Println(color.YellowString("Usage:"))
	fmt.Println("  command [path] [group:command|option] [arguments]")
}

func List(path string) {
	Header()
	Usage()
	fmt.Println("")
	fmt.Println(color.YellowString("Available commands:"))
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		errorMessage("Impossible to read directory")
	}
	for _, d := range dirs {
		if !d.IsDir() {
			errorMessage(d.Name() + " is not a directory. Please move it in a group.")
			os.Exit(0)
		}
	}

	for _, d := range dirs {
		fmt.Println(color.YellowString(d.Name()))
		files, err := ioutil.ReadDir(path+"/"+d.Name())
		if err != nil {
			errorMessage("Impossible to read directory")
		}
		for _, f := range files {
			if f.IsDir() {
				errorMessage(f.Name() + " is a directory. Please move it to group list.")
				os.Exit(0)
			}
			fmt.Printf(" %s:%s\n", color.GreenString(d.Name()), color.GreenString(f.Name()))
		}
	}
}