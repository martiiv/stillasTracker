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



const queryClient = new QueryClient()

function App() {
    return (
        <UserAuthContextProvider>
            <QueryClientProvider client={queryClient}>
                <TopBar/>
                <Routes>
                    <Route path="/prosjekt/*" element={<ProtectedRoute> <Project /></ProtectedRoute>} />
                    <Route path="/kart" element={<ProtectedRoute> <MapPage /></ProtectedRoute>} />
                    <Route path="/stillas" element={<ProtectedRoute> <Scaffolding /></ProtectedRoute>} />
                    <Route path="/project/:id" element={<ProtectedRoute> <PreView /></ProtectedRoute>} />
                    <Route path="/logistics" element={<ProtectedRoute> <Logistic /></ProtectedRoute>} />
                    <Route path="/" element={<Login />} />
                    <Route path="/signup" element={<Signup />} />
                </Routes>
                <ReactQueryDevtools initialIsOpen={false} position='bottom-right' />
            </QueryClientProvider>
        </UserAuthContextProvider>

    );
}

export default App;
