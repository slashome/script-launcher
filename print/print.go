package print

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

var version = "0"
var commentMargin int = 4

// SetVersion set the application version to be printed (default 0)
func SetVersion(v string) {
	version = v
}

// Error print the error message
func Error(message string) {
	printHeader()
	errorMessage(message)
}

// Error print the error message and usage
func ErrorWithUsage(message string) {
	Error(message)
	printUsage()
}

// Error print the error message, usage and list
func ErrorWithUsageAndList(message string, path string, group string) {
	Error(message)
	printUsage()
	printList(path, group)
}

// List print the list of available commands
// If group paramater is not empty string,
// it will print only the commands this group
func List(path string, group string) {
	printHeader()
	printList(path, group)
}

func printList(path string, group string) {
	fmt.Println("")
	fmt.Println(color.YellowString("Available commands:"))
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		errorMessage("Impossible to read directory")
	}
	// Iterate one time throught dir and file to check data and count longest
	// comment space
	var errors []string
	var textMaxLength int
	for _, d := range dirs {
		if !d.IsDir() {
			errors = append(errors, "\""+d.Name()+"\" is not a directory. Please move it in a group.")
			break
		}
		files, err := ioutil.ReadDir(path + "/" + d.Name())
		if err != nil {
			errors = append(errors, "Impossible to read directory \""+d.Name()+"\"")
		}
		for _, f := range files {
			if f.IsDir() {
				errors = append(errors, "\""+f.Name()+"\" is a directory. Please move it to group list.")
				break
			}
			textLength := utf8.RuneCountInString(d.Name()+":"+f.Name()) + commentMargin
			if textLength > textMaxLength {
				textMaxLength = textLength
			}
		}
	}

	for _, d := range dirs {
		if group != "" {
			if d.Name() != group {
				continue
			}
		}
		fmt.Println(color.YellowString(d.Name()))
		files, _ := ioutil.ReadDir(path + "/" + d.Name())
		for _, f := range files {
			filepath := path + "/" + d.Name() + "/" + f.Name()
			textLength := utf8.RuneCountInString(d.Name() + ":" + f.Name())
			comment := getComment(filepath)
			fmt.Printf(" %s:%s%s%s\n", color.GreenString(d.Name()), color.GreenString(f.Name()), getNSpaces(textMaxLength-textLength), comment)
		}
	}
}

func printHeader() {
	fmt.Println("Script Launcher", color.GreenString(version))
}

func getNSpaces(n int) string {
	var spaces string
	for i := 0; i < n; i++ {
		spaces = spaces + " "
	}
	return spaces
}

func printUsage() {
	fmt.Println("")
	fmt.Println(color.YellowString("Usage:"))
	fmt.Println("  command [path] [group:command|option] [arguments]")
}

func getComment(path string) string {
	fileHandle, _ := os.Open(path)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var count int
	for fileScanner.Scan() {
		count++
		if strings.HasPrefix(fileScanner.Text(), "# ") {
			return strings.TrimPrefix(fileScanner.Text(), "# ")
		}
		if count == 2 {
			break
		}
	}
	return "No comment"
}

func errorMessage(message string) {
	fmt.Println(color.RedString(message))
}
