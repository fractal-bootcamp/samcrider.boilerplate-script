// THIS FILE IS AUTOGENERATED, DO NOT MODIFY

package generated

var File__viteExpressFirebaseReadme = []string{
	"// Vite + Express App with Firebase Integration",
	"",
	"## General",
	"- You need to go to the firebase website and create a new project. Then add the project's serviceAccountKey to the serviceAccountKey file.",
	"- Note that docker is running, it just isn't shown in the terminal",
	"- If a file is in the gitignore but isn't greyed out:",
	"    1. delete and then retype a letter of the filename in the gitignore",
	"    2. save the gitignore file, now the file should be greyed out",
	"    3. if that doesn't work, reload the developer window",
	"- Note that a git repo has been initialized in the root directory of your project!",
	"",
	"## Frontend",
	"#### To run frontend:",
	"1. In terminal: npm run dev",
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
}
