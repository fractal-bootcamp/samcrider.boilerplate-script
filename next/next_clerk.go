package next_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Next_ClerkAuth(project_name string, docker_port string) {
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

	// remove readme and replace with clerk readme
	err = os.Remove("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("README.md", generated.File__nextClerkReadme)

	// install deps
	cmd_deps := utils.BoundCommand("npm", "install", "@clerk/nextjs")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make .env.local file
	utils.Create_File(".env.local", generated.File__nextClerkEnvLocal)

	// install dev deps (prisma)
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "prisma")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make dockerfile
	if docker_port == "10009" {
		utils.Create_File("docker-compose.yml", generated.File__docker)
	} else {
		utils.Create_Dynamic_Port_File("docker-compose.yml", "./file_generator/source_files/docker.txt", docker_port)
	}

	// get docker up
	cmd_docker := utils.BoundCommand("docker", "compose", "up", "-d")
	if err := cmd_docker.Run(); err != nil {
		fmt.Println(err)
		return
	}

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

	if docker_port == "10009" {
		utils.Create_File(".env", generated.File__firebaseEnv)
	} else {
		utils.Create_Dynamic_Port_File(".env", "./file_generator/source_files/firebaseEnv.txt", docker_port)
	}

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

	// run a db migration
	cmd_migration := utils.BoundCommand("npx", "prisma", "migrate", "dev")
	if err := cmd_migration.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// cd src directory
	err = os.Chdir("src")
	if err != nil {
		fmt.Println(err)
		return
	}

	// create clerk middleware file
	utils.Create_File("middleware.ts", generated.File__nextClerkMiddleware)

	// cd app directory
	err = os.Chdir("app")
	if err != nil {
		fmt.Println(err)
		return
	}

	// remove and replace the layout file
	err = os.Remove("layout.tsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.Create_File("layout.tsx", generated.File__nextClerkLayout)

	// cd out of app
	err = os.Chdir("..")
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

	// mkdir controllers
	err = os.Mkdir("controllers", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd into controllers
	err = os.Chdir("controllers")
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

	// create controller and types files
	utils.Create_File("controller.ts", generated.File__nextNoAuthController)
	utils.Create_File("types.ts", generated.File__firebaseAuthTypes)

}
