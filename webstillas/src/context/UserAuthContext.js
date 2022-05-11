import { createContext, useContext, useEffect, useState } from "react";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  onAuthStateChanged,
  signOut,
} from "firebase/auth";
import { auth } from "../firebase";

const userAuthContext = createContext();
//Hentet fra https://github.com/WebDevSimplified/React-Firebase-Auth


/**
 *Function that handles firebase log in and sign up.
 *
 * @param children
 * @returns {JSX.Element}
 */
export function UserAuthContextProvider({ children }) {
  const [user, setUser] = useState({});

  /**
   * Function to handle login from firebase.
   *
   * @param email the registered email of the user
   * @param password the corresponding password that matches the users email.
   * @returns {Promise<UserCredential>}
   */
  function logIn(email, password) {
    return signInWithEmailAndPassword(auth, email, password);
  }

  /**
   * Function to sign up a new user to firebase.
   *
   * @param email of the user that is to be registered.
   * @param password of the user that is to be registered.
   * @returns {Promise<UserCredential>}
   */
  function signUp(email, password) {
    return createUserWithEmailAndPassword(auth, email, password);
  }

  /**
   * Function that will log a user out of the system.
   * @returns {Promise<void>}
   */
  function logOut() {
    return signOut(auth);
  }



  //todo kommenter
  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (currentuser) => {
      console.log("Auth", currentuser);
      setUser(currentuser);
    });
    return () => {
      unsubscribe();
    };
  }, []);

  return (
    <userAuthContext.Provider
      value={{ user, logIn, signUp, logOut }}
    >
      {children}
    </userAuthContext.Provider>
  );
}

export function useUserAuth() {
  return useContext(userAuthContext);
}
