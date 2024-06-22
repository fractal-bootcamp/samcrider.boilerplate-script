import admin from "firebase-admin";

const serviceAccount = require("../../serviceAccountKey.json");

const app = admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
});

const auth = app.auth();

export default auth;
