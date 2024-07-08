package utils

import (
	"fmt"
	"os"
	"strings"
)

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
