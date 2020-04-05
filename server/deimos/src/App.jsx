import './App.scss';

import React, { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import ErrorPage from './components/ErrorPage';
import Graphs from './components/Graphs';
import Header from './components/Header';
import Home from './components/Home';
import Landing from './components/Landing';
import Strava from './components/Strava';
import { StatsContext, UserContext } from './contexts';

export default function App() {
  const [user, setUser] = useState({
    payload: null,
    loading: true,
    errors: null,
  });

  const [stats, setStats] = useState({
    payload: {
      workouts: 0,
      hours: 0,
      miles: 0,
      last_ten: [],
    },
    loading: true,
    errors: null,
  });

  return (
    <UserContext.Provider value={{ user, setUser }}>
      <StatsContext.Provider value={{ stats, setStats }}>
        <div className="App">
          <BrowserRouter>
            <Switch>
              <Route exact path="/home">
                <Header showStrava />
                <Home />
              </Route>
              <Route exact path="/graph">
                <Header showBack />
                <Graphs />
              </Route>
              <Route exact path="/dashboard/strava">
                <Header showBack />
                <Strava />
              </Route>
              <Route path="/error">
                <Header showBack />
                <ErrorPage />
              </Route>
              <Route>
                <Header />
                <Landing />
              </Route>
            </Switch>
          </BrowserRouter>
        </div>
      </StatsContext.Provider>
    </UserContext.Provider>
  );
}
