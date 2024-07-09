package next_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Next_Firebase(project_name string, docker_port string) {
	// create next app
	cmd := utils.BoundCommand("npx", "create-next-app", project_name, "--typescript")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// cd into project
	err := os.Chdir(project_name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// remove readme and replace with firebase readme
	err = os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__nextFirebaseReadme)

	utils.Work_wrapper(func() {

		// install deps
		cmd_deps := utils.BoundCommand("npm", "install", "firebase")

		if err := cmd_deps.Run(); err != nil {
			fmt.Println(err)
			return
		}

	}, "Installing Firebase package...")()

	utils.Work_wrapper(func() {

		// install dev deps (prisma)
		cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "prisma")

		if err := cmd_dev_deps.Run(); err != nil {
			fmt.Println(err)
			return
		}

	}, "Installing project dev packages...")()

	utils.Work_wrapper(func() {
		// make dockerfile
		utils.Revise_File("docker-compose.yml", generated.File__docker, docker_port)

		// get docker up
		cmd_docker := utils.BoundCommand("docker", "compose", "up", "-d")
		if err := cmd_docker.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Starting Docker container...")()

	utils.Work_wrapper(func() {
		// initialize prisma
		cmd_prisma := utils.BoundCommand("npx", "prisma", "init", "--datasource-provider", "postgreSQL")
		if err := cmd_prisma.Run(); err != nil {
			fmt.Println(err)
			return
		}

		// replace the .env file
		err = os.Remove(".env")
		if err != nil {
			fmt.Println(err)
			return
		}
		utils.Revise_File(".env", generated.File__nextFirebaseEnv, docker_port)

		// replace the gitignore file
		err = os.Remove(".gitignore")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File(".gitignore", generated.File__nextGitignore)

		// cd into prisma
		err = os.Chdir("prisma")
		if err != nil {
			fmt.Println(err)
			return
		}

		// remove the schema and create a new one
		err = os.Remove("schema.prisma")
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.Create_File("schema.prisma", generated.File__firebasePrismaSchema)

		// cd out of prisma
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}
	}, "Setting up Prisma...")()

	// run a db migration
	cmd_migration := utils.BoundCommand("npx", "prisma", "migrate", "dev")
	if err := cmd_migration.Run(); err != nil {
		fmt.Println(err)
		return
	}

	utils.Work_wrapper(func() {
		// cd src directory
		err = os.Chdir("src")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make utils folder, cd into it
		err = os.Mkdir("utils", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = os.Chdir("utils")
		if err != nil {
			fmt.Println(err)
			return
		}

		// create client.ts file
		utils.Create_File("client.ts", generated.File__client)

		// cd out of utils
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into app
		err = os.Chdir("app")
		if err != nil {
			fmt.Println(err)
			return
		}

		// remove page.tsx file
		err = os.Remove("page.tsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		// replace page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebasePage)

		// remove the layout file
		err = os.Remove("layout.tsx")
		if err != nil {
			fmt.Println(err)
			return
		}

		// replace the layout file
		utils.Create_File("layout.tsx", generated.File__nextFirebaseLayout)

		// mkdir login
		err = os.Mkdir("login", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into login
		err = os.Chdir("login")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebaseLoginPage)

		// cd out of login
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir dashboard
		err = os.Mkdir("dashboard", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into dashboard
		err = os.Chdir("dashboard")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make page.tsx file
		utils.Create_File("page.tsx", generated.File__nextFirebaseDashboardPage)

		// cd out of app
		err = os.Chdir("../../")

		// mkdir hooks
		err = os.Mkdir("hooks", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into hooks
		err = os.Chdir("hooks")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make useAuth hook
		utils.Create_File("useAuth.ts", generated.File__nextFirebaseHook)

		// cd out of hooks
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir components
		err = os.Mkdir("components", 0755)
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

		// mkdir providers
		err = os.Mkdir("providers", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into providers
		err = os.Chdir("providers")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make AuthProvider file
		utils.Create_File("AuthProvider.tsx", generated.File__nextFirebaseProvider)

		// cd out of providers
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

		// mkdir middleware
		err = os.Mkdir("middleware", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd into middleware
		err = os.Chdir("middleware")
		if err != nil {
			fmt.Println(err)
			return
		}

		// make middleware file
		utils.Create_File("middleware.ts", generated.File__nextFirebaseMiddleware)

		// cd out of middleware
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
		}

		// mkdir firebase
		err = os.Mkdir("firebase", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// cd firebase
		err = os.Chdir("firebase")
		if err != nil {
			fmt.Println(err)
			return
		}

		// create firebase config file
		utils.Create_File("config.ts", generated.File__nextFirebaseConfig)

	}, "Creating Utils, Components, and Library folders...")()

}
