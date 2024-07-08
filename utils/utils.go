package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
)

func BoundCommand(name string, arg ...string) *exec.Cmd {
	command := exec.Command(name, arg...)

	// bind command to terminal
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command
}

func Select(label string, opts []string) string {
	res := ""
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func Input(label string) string {
	res := ""
	prompt := &survey.Input{
		Message: label,
	}
	survey.AskOne(prompt, &res)

	return res
}

func CloseFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Close_and_Remove_File(f *os.File) {
	// close and remove the file
	CloseFile(f)

	os.Remove(f.Name())
}

func Create_File(name string, file_content []string) {
	// create file
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// makes sure the file closes when function finishes execution
	defer CloseFile(file)

	// loop through data and write lines
	for _, v := range file_content {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

/* NOTE: this function is currently hardcoded to only replace the docker port number */
func Revise_File(name string, file_content []string, new string) {
	// create file
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// makes sure the file closes when function finishes execution
	defer CloseFile(file)

	// loop through data and write lines, if line contains the word "10009", replace it with the user's input
	for _, v := range file_content {
		_, err := fmt.Fprintln(file, strings.Replace(v, "10009", new, -1))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func Work_wrapper(wrapped func(), suffix string) func() {
	return func() {
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Color("magenta")
		s.Suffix = " " + suffix
		s.Start()

		// call internal function
		wrapped()

		s.Stop()
	}

}
