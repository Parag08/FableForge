import React from "react";
import Login, { LandingPage } from "./pages/Landing";
import Home from "./pages/Home";
import { HashRouter, Route, Routes } from "react-router-dom";


function App() {
  return (
   <HashRouter>
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/home" element={<Home />} />
    </Routes>
   </HashRouter>
  );
}

export default App;