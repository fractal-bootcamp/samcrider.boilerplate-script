package main

// to build: go build -o scriptname

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	express_boil "sam.crider/boilerplate-script/express"
	vite_boil "sam.crider/boilerplate-script/vite"
)

func Checkboxes(label string, opts []string) string {
    res := ""
    prompt := &survey.Select{
        Message: label,
        Options: opts,
    }
    survey.AskOne(prompt, &res)

    return res
}

func main() {
    answer := Checkboxes(
        "Select Your Build Stack:",
        []string{
            "Vite + Express",
			"Next.js",
        },
    )

	if answer == "Vite + Express" {
		vite_boil.Vite()
		express_boil.Express()
		fmt.Println("Success! Boilerplate created.")
		return
	}
	if answer == "Next.js + " {
		fmt.Println("Coming Soon")
		fmt.Println("Success! Boilerplate created.")
		return
	}

	fmt.Println("Failure. Maybe you didn't select an option?")
}