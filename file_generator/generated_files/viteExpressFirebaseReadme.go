// THIS FILE IS AUTOGENERATED, DO NOT MODIFY

package generated

var File__viteExpressFirebaseReadme = []string{
"// Vite + Express App with Firebase Integration",
"",
"## General",
"- You need to go to the firebase website and create a new project. Then add the project's serviceAccountKey to the serviceAccountKey file.",
"- Note that docker is running, it just isn't shown in the terminal",
"- If a file is in the gitignore but isn't greyed out, try the following:",
"    1. delete and then retype a letter of the filename in the gitignore, save the gitignore file.",
"    2. No luck? Reload the developer window.",
"    3. No luck? The .env is probably in the staging area of git. Unstage it.",
"    4. Still no luck? Look it up or ask for help. (sorry lol)",
"",
"- Note that a git repo has been initialized in the root directory of your project!",
"- I added a user to the request object in the backend by declaring it in the global.d.ts file. ",
"    You may need to go into the ./utils/global.d.ts file and delete a line, then add it back in and save the file in order for ",
"    the global type to be recognized.",
"",
"## Frontend",
"#### To run frontend:",
"1. In terminal: npm run dev",
"",
"#### Firebase",
"- You need to get your firebase config object from the Firebase console (your project on the firebase website) and add it to the .env file.",
"- Go look at the .env file to see what variables are needed.",
"",
"#### To connect to your backend:",
"1. Add your backend's url to the frontend .env file",
"2. Add this to your package.json: \"proxy\": \"backend url\",",
"",
"## Backend",
"#### To run backend: ",
"1. Go to package.json",
"2. Add this script: \"start\": \"nodemon index.ts\"",
"3. In terminal: npm run start",
"",
"#### To connect to your frontend:",
"1. In the app.ts file add your frontend's url as an origin in the cors object",
"",
"## Resources",
"- Firebase: https://firebase.google.com/docs",
"- Express: https://expressjs.com/en/guide/routing.html",
"- Docker: https://docs.docker.com/guides/",
"- Vite: https://vitejs.dev/guide/",
"- Prisma: https://www.prisma.io/docs",
"",
}
