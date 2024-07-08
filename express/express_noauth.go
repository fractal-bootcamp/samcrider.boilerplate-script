package express_boil

import (
	"fmt"
	"os"

	generated "sam.crider/boilerplate-script/file_generator/generated_files"
	"sam.crider/boilerplate-script/utils"
)

func Express_NoAuth(docker_port string) {
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

	utils.Work_wrapper(func() {

		// install cors, dotenv, express, nodemon, ts-node
		cmd_deps := utils.BoundCommand("npm", "install", "express", "cors", "dotenv", "nodemon", "ts-node")

		if err := cmd_deps.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Installing backend packages...")()

	utils.Work_wrapper(func() {

		// install dev deps: cors types, express types, prisma
		cmd_dev_deps := utils.BoundCommand("npm", "install", "--save-dev", "@types/cors", "@types/express", "prisma")

		if err := cmd_dev_deps.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Installing backend dev packages...")()

	// make app.ts
	utils.Create_File("app.ts", generated.File__noAuthApp)

	utils.Work_wrapper(func() {
		// make dockerfile
		if docker_port == "10009" {
			utils.Create_File("docker-compose.yml", generated.File__docker)
		} else {
			utils.Revise_File("docker-compose.yml", generated.File__docker, docker_port)

		}

		// get docker up
		cmd_docker := utils.BoundCommand("docker", "compose", "up", "-d")
		if err := cmd_docker.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}, "Starting Docker container...")()

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
	if docker_port == "10009" {
		utils.Create_File(".env", generated.File__firebaseEnv)
	} else {
		utils.Revise_File(".env", generated.File__firebaseEnv, docker_port)
	}

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

	utils.Work_wrapper(func() {
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

		// create service file and types file
		utils.Create_File("controller.ts", generated.File__noAuthController)
		utils.Create_File("types.ts", generated.File__firebaseAuthTypes)

	}, "Creating Utils and Library files...")()
}
