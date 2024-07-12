package utils

import (
	"fmt"
	"os"
)

const helpText = `

Chiks: Your one-stop solution for hatching modern web projects with ease!

Usage: chiks [options]

Options:
  -h, --help           Display this help message

Currently Supported Stacks:
  - Next.js
  - Vite + Express
  - NOTE: Don't see your stack? Please add it!! Check out the CONTRIBUTING.md file! (https://github.com/SamuelRCrider/chiks/CONTRIBUTING.md)

Features:
  - Docker Integration
  - Prisma Database Setup
  - Auth Integration Options
  - Tailwind CSS Integration (optional)

Prerequisites (things you need to have installed):
  - Docker
  - PostgreSQL

For more information, visit: https://github.com/SamuelRCrider/chiks

To update Chiks:
  $ npm update -g chiks

To uninstall Chiks:
  1. $ npm uninstall -g chiks
  2. $ rm $(which chiks)

For bug reports or feature requests, please visit our GitHub repository.
`

func PrintHelp() {
	fmt.Fprint(os.Stderr, helpText)
}
