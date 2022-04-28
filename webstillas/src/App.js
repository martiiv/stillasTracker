import './App.css';
import { Routes, Route} from "react-router-dom";
import {Project} from "./components/projects/projects";
import {MapPage} from "./components/mapPage/mapPage";
import {Scaffolding} from "./components/scaffolding/scaffolding";
import TopBar from "./components/topBar/topBar";
import React from "react";
import {PreView} from "./components/projects/elements/preView";
import Logistic from "./components/logistics/logistic";
import { QueryClientProvider, QueryClient } from 'react-query'
import { ReactQueryDevtools } from 'react-query/devtools'


const queryClient = new QueryClient()

function App() {
    return (
        <QueryClientProvider client={queryClient}>
            <div className={"maintodo"}>
                <TopBar/>
                <Routes>
                    <Route path="/prosjekt/*" element={<Project />} />
                    <Route path="/kart" element={ <MapPage />} />
                    <Route path="/stillas" element={ <Scaffolding />} />
                    <Route path="/project/:id" element={<PreView />} />
                    <Route path="/logistics" element={<Logistic />} />
                </Routes>
            </div>
            <ReactQueryDevtools initialIsOpen={false} position='bottom-right' />
        </QueryClientProvider>

    );
}

export default App;
