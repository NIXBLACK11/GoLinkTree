import React from 'react';
import { ChakraProvider } from '@chakra-ui/react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Home from "./Routes/Home";
// import Home from './routes/Home';
// import Signup from './routes/Signup'
// import User from './routes/User';

const App = () => {
  return (
    <ChakraProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home/>}/>
          {/* <Route path="/home" element={<Home/>}/>
          <Route path="/signup" element={<Signup/>}/>
          <Route path="/user/:userName" element={<User/>}/> */}
        </Routes>
      </BrowserRouter>
    </ChakraProvider>
  );
};

export default App;
