import React from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { MantineProvider } from '@mantine/core'
import { Notifications } from '@mantine/notifications';


import CreateRoom from 'src/components/CreateRoom.tsx'
import Header from 'src/components/Header.tsx'
import Home from 'src/components/Home.tsx'
import Room from 'src/components/Room.tsx'

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <MantineProvider>
        <Notifications />
        <Header />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/room/create" element={<CreateRoom />} />
          <Route path="/room/:id" element={<Room />} />
        </Routes>
      </MantineProvider>
    </BrowserRouter>
  )
}

export default App
