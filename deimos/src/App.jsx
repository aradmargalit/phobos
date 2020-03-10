import './App.scss';

import React, { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import ActivityGraph from './components/ActivityGraph';
import Header from './components/Header';
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
    setLoading,
  };

  return (
    <UserContext.Provider value={userValue}>
      <div className="App">
        <BrowserRouter>
          <Switch>
            <Route exact path="/home">
              <Header />
              <Home />
            </Route>
            <Route exact path="/graph">
              <Header showBack />
              <ActivityGraph />
            </Route>
            <Route exact path="/">
              <Header />
              <Landing />
            </Route>
          </Switch>
        </BrowserRouter>
      </div>
    </UserContext.Provider>
  );
}
