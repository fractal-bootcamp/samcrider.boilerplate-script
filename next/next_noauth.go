package next_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Next_NoAuth(project_name string, docker_port string) {
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

	// remove readme and replace with no auth readme
	err = os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__nextNoAuthReadme)

	// install dev deps (prisma)
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "prisma")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

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
		utils.Revise_File(".env", generated.File__firebaseEnv, docker_port)

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

		utils.Create_File("schema.prisma", generated.File__noAuthPrismaSchema)

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
		utils.Mkdir_chdir("utils")

		// create client.ts file
		utils.Create_File("client.ts", generated.File__client)

		// cd out of utils
		err = os.Chdir("..")
		if err != nil {
			fmt.Println(err)
			return
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

		// mkdir controllers
		utils.Mkdir_chdir("controllers")

		// mkdir users
		utils.Mkdir_chdir("users")

		// create controller and types files
		utils.Create_File("controller.ts", generated.File__nextNoAuthController)
		utils.Create_File("types.ts", generated.File__firebaseFrontTypes)

	}, "Creating Utils, Components, and Library folders...")()

}
