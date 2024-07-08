package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
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

func Create_Dynamic_Dockerfile(name string, file_content []string, port int) {
	// open the docker file as read only
	source_file, err := os.OpenFile("./file_generator/source_files/docker.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ensure the file is closed once the function finishes execution
	defer CloseFile(source_file)

	// initialize the file scanner
	file_scanner := bufio.NewScanner(source_file)

	// create and open a new docker file
	docker_file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ensure the file is closed once the function finishes execution
	defer CloseFile(docker_file)

	// read each line of the current file
	for file_scanner.Scan() {

		// get the current line
		line := file_scanner.Text()

		// if line contains the word "10009", replace it with the user's input
		line = strings.Replace(line, "10009", strconv.Itoa(port), -1)

		// write the current line into the file
		_, err := fmt.Fprintln(docker_file, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := file_scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

}
