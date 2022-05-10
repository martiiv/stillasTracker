import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyB5XxJ-AIC_Bm38oOH4TjdeIBA0eNLRl7w",
  authDomain: "stillas-16563.firebaseapp.com",
  projectId: "stillas-16563",
  storageBucket: "stillas-16563.appspot.com",
  messagingSenderId: "586975019426",
  appId: "1:586975019426:web:83a11475b6ae32ffbc32fb"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export default app;
