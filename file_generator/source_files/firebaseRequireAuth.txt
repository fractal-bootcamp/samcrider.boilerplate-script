import { NextFunction, Request, Response } from "express";
import auth from "../lib/firebase/config";
import { client } from "./client";

export const requireAuth = async (
  req: Request,
  res: Response,
  next: NextFunction
) => {
  // headers = { ..., authorization: "Bearer ${token}" }
  const authToken = req.headers.authorization?.split(" ")[1];

  // if no token return error
  if (!authToken) {
    return res
      .status(401)
      .json({ error: "You don't belong here. Who do you know here?" });
  }

  // verify user in firebase
  try {
    const firebaseUser = await auth.verifyIdToken(authToken);
    // weird check here, I guess firebase auto configures email to be optional
    // TODO: fix this so email is required
    if (!firebaseUser || !firebaseUser.email) {
      throw new Error("authToken doesn't match any users in firebase");
    }

    // find user in db
    const dbUser = await client.user.findUnique({
      where: {
        firebaseId: firebaseUser.uid,
      },
    });
    // if user not in db but they are in firebase, they need added to db
    // this isn't properly type narrowing dbUser! TODO: fix
    if (!dbUser) {
      const newDbUser = await client.user.create({
        data: {
          email: firebaseUser.email,
          firebaseId: firebaseUser.uid,
        },
      });
      // set req["user"] = db user
      req["user"] = newDbUser;

      return next();
    }
    // if user in db, set req["user"] = db user
    req["user"] = { id: dbUser.id, email: dbUser.email };

    next();
  } catch (e) {
    console.error(e);
    return res.status(400).json({
      error: "Either couldn't verify user in firebase",
    });
  }
};
