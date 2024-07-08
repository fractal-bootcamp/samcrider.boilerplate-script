package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	next_boil "sam.crider/boilerplate-script/next"

	express_boil "sam.crider/boilerplate-script/express"
	"sam.crider/boilerplate-script/utils"
	vite_boil "sam.crider/boilerplate-script/vite"
)

// func main() {
// 	next_boil.Next_ClerkAuth("test")
// }

func main() {
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

	// get users auth preference
	auth_integration := utils.Select(
		"Pick an auth integration:",
		[]string{
			"Firebase",
			"Clerk",
			"None",
		},
	)

	// TODO: make this a switch case
	if stack == "Vite + Express" {
		// create a directory for the project, 0755 is the permission bits
		err := os.Mkdir(project_name, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("creating project directory")

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

		fmt.Print("installing auth", auth_integration)

		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressFirebaseReadme)

			// create the frontend
			vite_boil.Vite_FirebaseAuth()

			// create the backend
			express_boil.Express_FirebaseAuth()

			fmt.Println("Success! Boilerplate created.")
			return

		} else if auth_integration == "Clerk" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressClerkReadme)

			// create the frontend
			// vite_boil.Vite_ClerkAuth()

			// create the backend
			// express_boil.Express_ClerkAuth()

			fmt.Println("Success! Boilerplate created.")
			return

		} else if auth_integration == "None" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressNoAuthReadme)

			// create the frontend
			vite_boil.Vite_NoAuth()

			// create the backend
			express_boil.Express_NoAuth()

			fmt.Println("Success! Boilerplate created.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return

	}
	if stack == "Next.js" {
		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressFirebaseReadme)

			// create the app
			// next_boil.Next_FirebaseAuth(project_name)

			fmt.Println("Success! Boilerplate created.")
			return

		} else if auth_integration == "Clerk" {
			// create the app
			next_boil.Next_ClerkAuth(project_name)

			fmt.Println("Success! Boilerplate created.")
			return

		} else if auth_integration == "None" {
			// create the app
			next_boil.Next_NoAuth(project_name)

			fmt.Println("Success! Boilerplate created.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return

	}

	fmt.Println("Failure. Maybe you didn't select an option?")
}
