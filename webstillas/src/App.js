import './App.css';
import React from "react";
import {Routes, Route} from "react-router-dom";
import {Project} from "./components/projects/projects";
import {MapPage} from "./components/mapPage/mapPage";
import {Scaffolding} from "./components/scaffolding/scaffolding";
import TopBar from "./components/topBar/topBar";
import {PreView} from "./components/projects/elements/preView";
import Logistic from "./components/logistics/logistic";
import {QueryClientProvider, QueryClient} from 'react-query'
import {ReactQueryDevtools} from 'react-query/devtools'
import ProtectedRoute from "./components/ProtectedRoute";
import Login from "./components/Login";
import Signup from "./components/Signup";
import {UserAuthContextProvider} from "./context/UserAuthContext";
import AddProjectFunc from "./components/logistics/project/addProject";
import AddScaffolding from "./components/logistics/scaffold/addScaffolding";
import {UserInfo} from "./components/userinformation/userInfo";
import {NotFound} from "./components/error/error";
import {
    ADD_PROJECT_URL, ADD_SCAFFOLDING_URL, LOGIN,
    LOGISTICS_URL,
    MAP_URL, NOTFOUND,
    PROJECT_URL,
    PROJECT_URL_ID,
    SCAFFOLDING_URL, SIGNUP, USERINFO_URL
} from "./components/constants";


const queryClient = new QueryClient()

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
                    <Route path={PROJECT_URL_ID} exact={true} element={<ProtectedRoute> <PreView/></ProtectedRoute>}/>
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
