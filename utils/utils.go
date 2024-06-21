package utils

import (
	"os"
	"os/exec"
)


func BoundCommand(name string, arg ...string) (*exec.Cmd) {
	command := exec.Command(name, arg...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	
	
	return command
}

func GenerateConstantsFromFiles (){
	// crawl through the directory and transform *.constant.txt files to constant string arrays
	// parse those files and write them to a new .generated.go file
}