package vite_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Vite_ClerkAuth() {

	cmd := utils.BoundCommand("npx", "create-vite@latest", "frontend", "--", "--template", "react-ts")

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
	cmd = utils.BoundCommand("npm", "install", "axios", "@clerk/clerk-js", "@clerk/clerk-react")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// create .env file
	utils.Create_File(".env", generated.File__noAuthFrontEnv)

	// create .env.local file
	utils.Create_File(".env.local", generated.File__viteClerkEnvLocal)

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

	// cd into src
	err = os.Chdir("src")
	if err != nil {
		fmt.Println(err)
		return
	}

	// remove the main.tsx file
	err = os.Remove("main.tsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// replace the main.tsx file
	utils.Create_File("main.tsx", generated.File__viteClerkMain)

	// remove the app.tsx file
	err = os.Remove("app.tsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// replace the App.tsx file
	utils.Create_File("App.tsx", generated.File__viteClerkApp)

	// ask if user wants tailwind
	tailwind_check := utils.Select(
		"Do you need tailwind for this project?",
		[]string{
			"Yes",
			"No",
		},
	)

	if tailwind_check == "Yes" {
		// ask if user would like to add daisyUI, Shadcn UI, or just Tailwind
		ui_check := utils.Select(
			"Which UI framework would you like to use?",
			[]string{
				"Shadcn UI",
				"DaisyUI",
				"None (base Tailwind)",
			},
		)

		/* install tailwind */

		// cd out of src
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

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

		// TODO: make this a switch case if we want to add more UI frameworks
		if ui_check == "Shadcn UI" {
			// install tailwind with shadcn ui

			// cd out of src
			err = os.Chdir("..")
			if err != nil {
				fmt.Println(err)
				return
			}

			// remove the tsconfig.json file
			err = os.Remove("tsconfig.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			// replace the tsconfig.json file
			utils.Create_File("tsconfig.json", generated.File__viteTsconfig)

			// install node types
			cmd = utils.BoundCommand("npm", "i", "-D", "@types/node")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// remove the vite.config.ts file
			err = os.Remove("vite.config.ts")
			if err != nil {
				fmt.Println(err)
				return
			}

			// replace the vite.config.ts file
			utils.Create_File("vite.config.ts", generated.File__viteShadConfig)

			// run shadcn ui init
			cmd = utils.BoundCommand("npx", "shadcn-ui@latest", "init")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// cd into src
			err = os.Chdir("src")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if ui_check == "DaisyUI" {
			// install tailwind with daisy ui

			// cd out of src
			err = os.Chdir("..")
			if err != nil {
				fmt.Println(err)
				return
			}

			// install daisy ui
			cmd = utils.BoundCommand("npm", "i", "-D", "daisyui")
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
				return
			}

			// remove the tailwind.config.js file
			err = os.Remove("tailwind.config.js")
			if err != nil {
				fmt.Println(err)
				return
			}

			// replace the tailwind.config.js file
			utils.Create_File("tailwind.config.js", generated.File__viteDaisyTconfig)

			// cd into src
			err = os.Chdir("src")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

	// mkdir components
	utils.Mkdir_chdir("components")

	// mkdir pages
	utils.Mkdir_chdir("pages")

	// create example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of pages
	err = os.Chdir("..")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir compound
	utils.Mkdir_chdir("compound")

	// create example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of compound
	err = os.Chdir("..")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir base
	utils.Mkdir_chdir("base")

	// create example component
	utils.Create_File("Example.tsx", generated.File__exampleComponent)

	// cd out of components
	err = os.Chdir("../../")
	if err != nil {
		fmt.Println(err)
		return
	}

	// mkdir lib
	utils.Mkdir_chdir("lib")

	// mkdir services
	utils.Mkdir_chdir("services")

	// mkdir users
	utils.Mkdir_chdir("users")

	// create service file and types file
	utils.Create_File("service.ts", generated.File__viteClerkService)
	utils.Create_File("types.ts", generated.File__firebaseAuthTypes)

	// cd back to project root in preparation for creating the backend
	err = os.Chdir("../../../../../")
	if err != nil {
		fmt.Println(err)
		return
	}
}
