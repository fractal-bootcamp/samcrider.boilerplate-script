package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)


func BoundCommand(name string, arg ...string) (*exec.Cmd) {
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

func Name_Project(label string) string {
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
