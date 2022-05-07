import './App.css';
import React from "react";
import { Routes, Route} from "react-router-dom";
import {Project} from "./components/projects/projects";
import {MapPage} from "./components/mapPage/mapPage";
import {Scaffolding} from "./components/scaffolding/scaffolding";
import TopBar from "./components/topBar/topBar";
import {PreView} from "./components/projects/elements/preView";
import Logistic from "./components/logistics/logistic";
import { QueryClientProvider, QueryClient } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'
import ProtectedRoute from "./components/ProtectedRoute";
import Login from "./components/Login";
import Signup from "./components/Signup";
import {UserAuthContextProvider, useUserAuth} from "./context/UserAuthContext";
import auth from "./firebase";



const queryClient = new QueryClient()

function App() {
    return (
        //Authorisation of user
        <UserAuthContextProvider>
            {//Caching provider client
            }
            <QueryClientProvider client={queryClient}>
                <TopBar/> {/*Topbar for the user to navigate throughout the webpage}*/}
                <Routes> {/*Router that creates the routes the user is able to navigate*/}
                    <Route path="/prosjekt/*" element={<ProtectedRoute> <Project/></ProtectedRoute>}/>
                    <Route path="/kart" element={<ProtectedRoute> <MapPage/></ProtectedRoute>}/>
                    <Route path="/stillas" element={<ProtectedRoute> <Scaffolding/></ProtectedRoute>}/>
                    <Route path="/project/:id" element={<ProtectedRoute> <PreView/></ProtectedRoute>}/>
                    <Route path="/logistics" element={<ProtectedRoute> <Logistic/></ProtectedRoute>}/>
                    <Route path="/" element={<Login/>}/>
                    <Route path="/signup" element={<Signup/>}/>
                </Routes>
                <ReactQueryDevtools initialIsOpen={true} />
            </QueryClientProvider>
        </UserAuthContextProvider>

    );

}

export default App;
