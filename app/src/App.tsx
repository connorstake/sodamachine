import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Routes, Route, BrowserRouter as Router ,Link, Navigate } from "react-router-dom";
import Register from './components/Register';
import Login from './components/Login';
import Profile from './components/Profile';
import HOCForRouteProps from './HOCForRouteProps';

function App() {
  return (
    <Router>
      <div className="app-container">
      <Routes>
        <Route path='/' element={<HOCForRouteProps Component={Login} />} />
        <Route path='/register' element={<HOCForRouteProps Component={Register} />}/>
        <Route path='/login' element={<HOCForRouteProps Component={Login} />} />
        <Route path='/profile' element={<HOCForRouteProps Component={Profile} />} />
      </Routes>
      </div>
  </Router>
  );
}

export default App;
