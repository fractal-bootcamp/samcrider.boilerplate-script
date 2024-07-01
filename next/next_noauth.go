package next_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"

	"sam.crider/boilerplate-script/utils"
)

func Next_NoAuth(project_name string) {
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

	// install dev deps (prisma)
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "prisma")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make dockerfile
	utils.Create_File("docker-compose.yml", generated.File__docker)

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

	utils.Create_File(".env", generated.File__firebaseEnv)

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

	// docker compose up
	cmd_dockerUp := utils.BoundCommand("docker", "compose", "up", "-d")
	if err := cmd_dockerUp.Run(); err != nil {
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
