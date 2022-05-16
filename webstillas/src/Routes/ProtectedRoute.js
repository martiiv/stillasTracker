import React from "react";
import { Navigate } from "react-router-dom";
import { useUserAuth } from "../Config/UserAuthContext";

/**
 * Function that will check if the user is authenticated, before sending the user to a protected route
 *
 * @param children the element route.
 * @returns {JSX.Element|*}
 */
const ProtectedRoute = ({ children }) => {
  const { user } = useUserAuth();

  console.log("Check user in Private: ", user);
  if (!user) {
    return <Navigate to="/" />;
  }
  return children;
};

export default ProtectedRoute;
