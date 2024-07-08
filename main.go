package main

// to build: go build -o scriptname

import (
	"fmt"
	"os"
	"strconv"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	next_boil "sam.crider/boilerplate-script/next"

	express_boil "sam.crider/boilerplate-script/express"
	"sam.crider/boilerplate-script/utils"
	vite_boil "sam.crider/boilerplate-script/vite"
)

func main() {
	vite_boil.Vite_FirebaseAuth()
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
	project_name := utils.Input(
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

	// get users docker port preference
	docker_port := utils.Input(
		"What docker port should the backend be on? (default: 10009)",
	)
	if docker_port == "" {
		docker_port = "10009"
	} else {
		// make sure the port is a number
		docker_port_int, err := strconv.Atoi(docker_port)
		if err != nil {
			fmt.Println("Invalid docker port number")
			return
		}
		// make sure the port is between 1 and 65535
		if docker_port_int < 1 || docker_port_int > 65535 {
			fmt.Println("Docker port number must be between 1 and 65535")
			return
		}
	}

	// TODO: make this a switch case
	if stack == "Vite + Express" {
		utils.Work_wrapper(func() {
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
		}, "Initializing project...")()

		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressFirebaseReadme)

			// create the frontend
			vite_boil.Vite_FirebaseAuth()

			// create the backend
			express_boil.Express_FirebaseAuth(docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "Clerk" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressClerkReadme)

			// create the frontend
			// vite_boil.Vite_ClerkAuth()

			// create the backend
			// express_boil.Express_ClerkAuth(docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "None" {
			// add readme
			utils.Create_File("README.md", generated.File__viteExpressNoAuthReadme)

			// create the frontend
			vite_boil.Vite_NoAuth()

			// create the backend
			express_boil.Express_NoAuth(docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return

	}
	if stack == "Next.js" {
		// TODO: make this a switch case
		if auth_integration == "Firebase" {
			// create the app
			// next_boil.Next_FirebaseAuth(project_name, docker_port_int)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "Clerk" {
			// create the app
			next_boil.Next_ClerkAuth(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return

		} else if auth_integration == "None" {
			// create the app
			next_boil.Next_NoAuth(project_name, docker_port)

			fmt.Println("Success! Boilerplate created. Check the root directory README.md for further instructions.")
			return
		}

		fmt.Println("Failure. Maybe you didn't select an option?")
		return

	}

	fmt.Println("Failure. Maybe you didn't select an option?")
}
