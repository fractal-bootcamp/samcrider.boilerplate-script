// Vite + Express App with Firebase Integration

## General
- You need to go to the firebase website and create a new project. Then add the project's serviceAccountKey to the serviceAccountKey file.
- Note that docker is running, it just isn't shown in the terminal


## Frontend
#### To run frontend:
1. In terminal: npm run dev

#### To connect to your backend:
1. Add your backend's url to the frontend .env file
2. Add this to your package.json: "proxy": "backend url",

## Backend
#### To run backend: 
1. Go to package.json
2. Add this script: "start": "nodemon index.ts"
3. In terminal: npm run start

#### To connect to your frontend:
1. In the app.ts file add your frontend's url as an origin in the cors object

