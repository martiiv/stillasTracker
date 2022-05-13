import '../Assets/Styles/App.css';
import React from "react";
import {Routes, Route} from "react-router-dom";
import {Project} from "../Pages/projects";
import {MapPage} from "../Pages/mapPage";
import {Scaffolding} from "../Pages/scaffolding";
import TopBar from "../Layout/topBar/topBar";
import {PreViewSite} from "../components/projects/preViewSite";
import {QueryClientProvider, QueryClient} from 'react-query'
import {ReactQueryDevtools} from 'react-query/devtools'
import ProtectedRoute from "./ProtectedRoute";
import Login from "../Pages/Login";
import Signup from "../Pages/Signup";
import {UserAuthContextProvider} from "../Config/UserAuthContext";
import AddProjectFunc from "../components/addproject/addProject";
import AddScaffolding from "../Pages/addScaffolding";
import {UserInfo} from "../Pages/userInfo";
import {NotFound} from "../components/Indicators/error";
import {
    ADD_PROJECT_URL, ADD_SCAFFOLDING_URL, LOGIN,
    MAP_URL, NOTFOUND,
    PROJECT_URL,
    PROJECT_URL_ID,
    SCAFFOLDING_URL, SIGNUP, USERINFO_URL
} from "../Constants/webURL";


const queryClient = new QueryClient()


/**
 * Function that will route the
 * @returns {JSX.Element}
 * @constructor
 */
function App() {
    return (
        //Authorisation of user
        <UserAuthContextProvider>
            {/*Caching provider client*/}
            <QueryClientProvider client={queryClient}>
                <TopBar/> {/*Topbar for the user to navigate throughout the webpage}*/}
                <Routes> {/*Router that creates the routes the user is able to navigate*/}
                    <Route path={PROJECT_URL} exact={true} element={<ProtectedRoute> <Project/></ProtectedRoute>}/>
                    <Route path={MAP_URL} exact={true} element={<ProtectedRoute> <MapPage/></ProtectedRoute>}/>
                    <Route path={SCAFFOLDING_URL} exact={true} element={<ProtectedRoute> <Scaffolding/></ProtectedRoute>}/>
                    <Route path={PROJECT_URL_ID} exact={true} element={<ProtectedRoute> <PreViewSite/></ProtectedRoute>}/>
                    <Route path={LOGIN} exact={true} element={<Login/>}/>
                    <Route path={SIGNUP} exact={true} element={<Signup/>}/>
                    <Route path={ADD_PROJECT_URL} exact={true}
                           element={<ProtectedRoute> <AddProjectFunc/></ProtectedRoute>}/>
                    <Route path={ADD_SCAFFOLDING_URL} exact={true}
                           element={<ProtectedRoute> <AddScaffolding/></ProtectedRoute>}/>
                    <Route path={USERINFO_URL} exact={true} element={<ProtectedRoute> <UserInfo/></ProtectedRoute>}/>
                    <Route path={NOTFOUND} element={<NotFound/>}/>
                </Routes>
                <ReactQueryDevtools initialIsOpen={true}/>
            </QueryClientProvider>
        </UserAuthContextProvider>

    );
}

export default App;
