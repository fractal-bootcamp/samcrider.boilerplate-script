# Hatch Your Next Project With Chiks!

<img src="./public/chick.svg" width="100" height="100" alt="SVG Image"><img src="./public/chick.svg" width="100" height="100" alt="SVG Image"><img src="./public/chick.svg" width="100" height="100" alt="SVG Image"><img src="./public/chick.svg" width="100" height="100" alt="SVG Image"><img src="./public/chick.svg" width="100" height="100" alt="SVG Image"><img src="./public/chick.svg" width="100" height="100" alt="SVG Image">

### Requirements

- Docker
- PostgreSQL

### Features

- Currently
  - you can use this script to create a Next.js project or a Vite + Express project
- Spins up
  - a docker container for you
  - a Prisma database for you and runs your first migration
- You get to choose
  - the auth integration you want to use
  - the docker port you want to use
  - whether you want to use Tailwind or not

### How to use

1. npm install -g chiks
2. cd into you project directory or wherever you keep your projects
3. chiks
4. Thats it!

### How to update

1. npm update -g chiks
2. Boom, done!

### How to uninstall (if you're a meanie)

1. npm uninstall -g chiks
2. rm $(which chicks) - NOTE #1
3. Boom, done.

#### Notes

1. Due to npm v7+, uninstall scripts won't work so the package's preuninstall script will not remove the binary from your system. That is why you need to run rm $(which chicks) to remove the binary. Once you've done that (and npm uninstall -g chiks), you will have successfully uninstalled chiks.
