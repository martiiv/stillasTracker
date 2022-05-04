import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyBsymxf23p52Pe0pB9ZuD_7HI6yBMMWN8g",
  authDomain: "cryptocurrency-7d594.firebaseapp.com",
  projectId: "cryptocurrency-7d594",
  storageBucket: "cryptocurrency-7d594.appspot.com",
  messagingSenderId: "796346602164",
  appId: "1:796346602164:web:e10a4f566951b0fd18fde1",
  measurementId: "G-5SZ9MHB1DM"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export default app;
