package express_boil

import (
	"fmt"
	"os"

	"sam.crider/boilerplate-script/utils"
)

func Express() {
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
	data := []string{
		"import app from './app';",
		"const PORT = [your port]",
		"app.listen(PORT, () => {",
		"console.log(`Server running on port ${PORT}`);",
		"});"}
	utils.Create_File("index.ts", data)
	

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
	data = []string{
		"import express from 'express';",
		"import cors from 'cors';",
		"const app = express();",
		"app.use(",
		"cors({",
		"origin: ['origin(s)'],",
		"allowedHeaders: ['headers'],",
		"})",
		");",
		"app.use(express.json());",
		"app.use('[path], [router]');",
		"export default app;",
	}
	utils.Create_File("app.ts", data)

	// make dockerfile
	data = []string{
		"version: '3.9'",
		"services:",
		"postgres:",
		"image: postgres:13",
		"environment:",
		"POSTGRES_USER: postgres",
		"POSTGRES_PASSWORD: postgres",
		"command: -c fsync=off -c full_page_writes=off -c synchronous_commit=off -c max_connections=500",
		"ports:",
		" - ['example: 10002']:5432",
	}
	utils.Create_File("docker-compose.yml", data)

	// initialize primsa
	cmd_prisma := utils.BoundCommand("npx", "prisma", "init", "--datasource-provider", "postgreSQL")
	if err := cmd_prisma.Run(); err != nil {
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
	data = []string{
		"import { PrismaClient } from '@prisma/client';",
		"const client = new PrismaClient();",
		"export default client;",
	}
	utils.Create_File("client.ts", data)
}
