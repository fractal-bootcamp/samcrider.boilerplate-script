package file_generator

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func scrub(line string) string {
	cleaner := func(r rune) rune {
		if r == '"' {
			return '\''
		}
		return r
	}

	// call cleaner using map function
	clean_line := strings.Map(cleaner, line)

	return clean_line
}

func serialize_lines(file fs.DirEntry) {
	// open the current file as read only
	source_file, err := os.OpenFile("./file_generator/source_files/" + file.Name(), os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		source_file.Close()
		return
	}

	// initialize the file scanner
	file_scanner := bufio.NewScanner(source_file)

	// create and open a new file in /generated
	generated_file, err := os.Create("./file_generator/generated_files/" + file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	// read each line of the current file
	for file_scanner.Scan() {
		// change any double quotes to single quotes
		dirtyText := file_scanner.Text()
		cleanText := scrub(dirtyText)

		// write the current line into the generated file with quotes and a comma 
		_, err := fmt.Fprintln(generated_file, `"` + cleanText + `",`)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := file_scanner.Err(); err != nil {
		fmt.Println(err)
		source_file.Close()
		generated_file.Close()
		return
	}
	// close the files
	source_file.Close()
	generated_file.Close()
}

func Generate_Files() {
	// read from the source directory
	items, err := os.ReadDir("./file_generator/source_files")
	if err != nil {
		fmt.Println(err)
		return
	}
	// serialize each file
    for _, item := range items {
        serialize_lines(item)
    }
}