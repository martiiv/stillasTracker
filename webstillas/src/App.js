import './App.css';
import { Routes, Route} from "react-router-dom";
import Projects from "./components/projects/projects";
import MapPage from "./components/mapPage/mapPage";
import Scaffolding from "./components/scaffolding/scaffolding";
import TopBar from "./components/topBar/topBar";


function App() {
  return (
          <div>
              <TopBar />
                  <Routes>
                      <Route path="/" element={<TopBar />} />
                      <Route path="/prosjekt" element={<Projects />} />
                      <Route path="/kart" element={ <MapPage />} />
                      <Route path="/stillas" element={ <Scaffolding />} />
                  </Routes>
          </div>
  );
}

export default App;
