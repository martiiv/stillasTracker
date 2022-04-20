import './App.css';
import { Routes, Route} from "react-router-dom";
import Projects from "./components/projects/projects";
import MapPage from "./components/mapPage/mapPage";
import Scaffolding from "./components/scaffolding/scaffolding";
import TopBar from "./components/topBar/topBar";
import React from "react";
import PreView from "./components/projects/elements/preView";
import Logistic from "./components/logistics/logistic";


function App() {

    return (
          <div className={"maintodo"}>
              <TopBar/>
                  <Routes>
                      <Route path="/prosjekt/*" element={<Projects />} />
                      <Route path="/kart" element={ <MapPage />} />
                      <Route path="/stillas" element={ <Scaffolding />} />
                      <Route path="/project/:id" element={<PreView />} />
                      <Route path="/logistics" element={<Logistic />} />
                  </Routes>
          </div>
  );
}

export default App;
