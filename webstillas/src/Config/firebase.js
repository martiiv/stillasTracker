import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import {
  FirbaseAPIKEY, FirbaseAppID,
  FirbaseMessagingSenderId,
  FirbaseprojectId,
  FirbaseStorageBucket,
  FirebaseAuthDomain
} from "./firebaseConfig";

//Firebase configuration to database.
const firebaseConfig = {
  apiKey: FirbaseAPIKEY,
  authDomain: FirebaseAuthDomain,
  projectId: FirbaseprojectId,
  storageBucket: FirbaseStorageBucket,
  messagingSenderId: FirbaseMessagingSenderId,
  appId: FirbaseAppID
};

// Initialize Firebase.
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export default app;
