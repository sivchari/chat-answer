import React from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<h1>Home</h1>} />
        <Route path="/room/create" element={<h1>Create Room</h1>} />
        <Route path="/room/:id" element={<h1>Room</h1>} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
