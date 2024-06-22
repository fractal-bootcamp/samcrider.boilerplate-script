package express_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Express_NoAuth() {
	// mkdir for backend, 0755 is the permission bits
	err := os.Mkdir("backend", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cd into the new directory
	err = os.Chdir("backend")
	if err != nil {
		fmt.Println(err)
		return
	}

	// initialize npm project
	cmd_npm := utils.BoundCommand("npm", "init")
	if err := cmd_npm.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// create index.ts file in new project
	utils.Create_File("index.ts", generated.File__index)

	// install cors, dotenv, express, nodemon, ts-node
	cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node")

	if err := cmd_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// install dev deps: cors types, express types, prisma
	cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express", "prisma")

	if err := cmd_dev_deps.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// make app.ts
	utils.Create_File("app.ts", generated.File__noAuthApp)

	// make dockerfile
	utils.Create_File("docker-compose.yml", generated.File__docker)

	// initialize primsa
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

	// // replace the gitignore file
	// err = os.Remove(".gitignore")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// utils.Create_File(".gitignore", generated.File__firebaseGitignore)

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
}
