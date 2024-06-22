package vite_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Vite_NoAuth() {
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

	// npm install all the vite packages
	cmd = utils.BoundCommand("npm", "install")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// import deps
	cmd = utils.BoundCommand("npm", "install", "axios")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// create .env file
	utils.Create_File(".env", generated.File__noAuthFrontEnv)

	// replace the gitignore file
	err = os.Remove(".gitignore")
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

	// ask if user wants tailwind
	tailwind_check := utils.Select(
		"Do you need tailwind for this project?",
		[]string{
			"Yes",
			"No",
		},
	)

	if tailwind_check == "Yes" {
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
		err = os.Remove("tailwind.config.js")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("tailwind.config.js", generated.File__firebaseFrontTailwindConfig)

		// replace index.css file
		// cd into src
		err := os.Chdir("src")
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

		// cd out of src
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	// cd into src
	err = os.Chdir("src")
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

	// mkdir services
	err = os.Mkdir("services", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd into services
	err = os.Chdir("services")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir users
	err = os.Mkdir("users", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd into users
	err = os.Chdir("users")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create service file and types file
	utils.Create_File("service.ts", generated.File__noAuthService)
	utils.Create_File("types.ts", generated.File__firebaseFrontTypes)

}
