package vite_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Vite_FirebaseAuth() {

	// create vite app
	utils.Work_wrapper(func() {
		cmd := utils.BoundCommand("npm", "create", "vite@latest", "frontend")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// cd into frontend
		err := os.Chdir("frontend")
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "Creating frontend...")()

	utils.Work_wrapper(func() {
		// npm install all the vite packages
		cmd := utils.BoundCommand("npm", "install")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}

	}, "Installing frontend packages...")()

	utils.Work_wrapper(func() {
		// import deps
		cmd := utils.BoundCommand("npm", "install", "axios", "firebase")

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Installing Firebase and Axios packages...")()

	utils.Work_wrapper(func() {
		// create .env file
		utils.Create_File(".env", generated.File__firebaseFrontEnv)

		// replace the gitignore file
		err := os.Remove(".gitignore")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File(".gitignore", generated.File__firebaseFrontGitignore)

		// replace the vite.config file
		err = os.Remove("vite.config.ts")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("vite.config.ts", generated.File__firebaseFrontViteConfig)

	}, "Restructuring Vite boilerplate...")()

	// ask if user wants tailwind
	tailwind_check := utils.Select(
		"Do you need tailwind for this project?",
		[]string{
			"Yes",
			"No",
		},
	)

	if tailwind_check == "Yes" {
		utils.Work_wrapper(func() {
			// install tailwind
			cmd := utils.BoundCommand("npm", "install", "-D", "tailwindcss", "postcss", "autoprefixer")

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// initialize tailwind
			cmd = utils.BoundCommand("npx", "tailwindcss", "init", "-p")

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// replace tailwind config file
			err := os.Remove("tailwind.config.js")
			if err != nil {
				fmt.Println(err)
				return
			}

			utils.Create_File("tailwind.config.js", generated.File__firebaseFrontTailwindConfig)

			// replace index.css file
			// cd into src
			err = os.Chdir("src")
			if err != nil {
				fmt.Println(err)
				return
			}

			err = os.Remove("index.css")
			if err != nil {
				fmt.Println(err)
				return
			}

			utils.Create_File("index.css", generated.File__firebaseFrontIndexCss)

		}, "Adding Tailwind...")()

	}

	utils.Work_wrapper(func() {
		// mkdir components
		err := os.Mkdir("components", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into components
		err = os.Chdir("components")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir pages
		err = os.Mkdir("pages", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir compound
		err = os.Mkdir("compound", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir base
		err = os.Mkdir("base", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd out of components
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir lib
		err = os.Mkdir("lib", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd lib
		err = os.Chdir("lib")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir's firebase and services
		err = os.Mkdir("firebase", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = os.Mkdir("services", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into firebase
		err = os.Chdir("firebase")
		if err != nil {
			fmt.Println(err)
			return
		}

		// created firebase config
		utils.Create_File("config.ts", generated.File__firebaseFrontConfig)

		// cd into services
		err = os.Chdir("../services")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir auth
		err = os.Mkdir("auth", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into auth
		err = os.Chdir("auth")
		if err != nil {
			fmt.Println(err)
			return
		}

		// create service file and types file
		utils.Create_File("service.ts", generated.File__firebaseFrontService)
		utils.Create_File("types.ts", generated.File__firebaseFrontTypes)

		// cd back to project root in preparation for creating the backend
		err = os.Chdir("../../../../")
		if err != nil {
			fmt.Println(err)
			return
		}
	}, "Creating Library files...")()
}
