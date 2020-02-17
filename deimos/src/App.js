import React, { useState } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Header from './components/Header';
import './App.css';
import Home from './components/Home';
import Landing from './components/Landing';
import { UserContext } from './contexts';

export default function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  const userValue = {
    user,
    setUser,
    loading,
    setLoading
  };

  return (
    <UserContext.Provider value={userValue}>
      <div className="App">
        <BrowserRouter>
          <div>
            <Header />
            <Switch>
              <Route exact path="/home">
                <Home />
              </Route>
              <Route exact path="/">
                <Landing />
              </Route>
            </Switch>
          </div>
        </BrowserRouter>
      </div>
    </UserContext.Provider>
  );
}
