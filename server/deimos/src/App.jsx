import './App.scss';

import React, { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Graphs from './components/Graphs';
import Header from './components/Header';
import Home from './components/Home';
import Landing from './components/Landing';
import Strava from './components/Strava';
import { StatsContext, UserContext } from './contexts';

export default function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    workouts: 0,
    hours: 0,
    miles: 0,
    last_ten: [],
  });
  const [statsLoading, setStatsLoading] = useState(true);

  const userValue = {
    user,
    setUser,
    loading,
    setLoading,
  };

  return (
    <UserContext.Provider value={userValue}>
      <StatsContext.Provider
        value={{ stats, setStats, statsLoading, setStatsLoading }}
      >
        <div className="App">
          <BrowserRouter>
            <Switch>
              <Route exact path="/home">
                <Header />
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
