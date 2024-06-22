package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"

	express_boil "sam.crider/boilerplate-script/express"
	"sam.crider/boilerplate-script/utils"
	vite_boil "sam.crider/boilerplate-script/vite"
)

func main() {
	express_boil.Express_FirebaseAuth()
}

func _main() {
	// get the user's selected stack
	stack := utils.Select(
		"Select Your Build Stack:",
		[]string{
			"Vite + Express",
			"Next.js",
		},
	)

	// get the user's project name
	project_name := utils.Name_Project(
		"What's the name of this project?",
	)

	auth_integration := utils.Select(
		"Pick an auth integration:",
		[]string{
			"Firebase",
			"Clerk",
			"None",
		},
	)

	if stack == "Vite + Express" {
		// create a directory for the project, 0755 is the permission bits
		err := os.Mkdir(project_name, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into the new project
		err = os.Chdir(project_name)
		if err != nil {
			fmt.Println(err)
			return
		}

		// initialize git for the project
		cmd := utils.BoundCommand("git", "init")
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		if auth_integration == "Firebase" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressFirebaseReadme)

			// create the frontend
			// vite_boil.Vite_FirebaseAuth()

			// create the backend
			// express_boil.Express_FirebaseAuth()

			fmt.Println("Success! Boilerplate created.")

		} else if auth_integration == "Clerk" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressClerkReadme)

			// create the frontend
			// vite_boil.Vite_ClerkAuth()

			// create the backend
			// express_boil.Express_ClerkAuth()

			fmt.Println("Success! Boilerplate created.")

		} else if auth_integration == "None" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressNoAuthReadme)

			// create the frontend
			vite_boil.Vite_NoAuth()

			// create the backend
			express_boil.Express_NoAuth()

			fmt.Println("Success! Boilerplate created.")
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return

	}
	if stack == "Next.js" {
		fmt.Println("Coming Soon")
		fmt.Println("Success! Boilerplate created.")
		return
	}

	fmt.Println("Failure. Maybe you didn't select an option?")
}
