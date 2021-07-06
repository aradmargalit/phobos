import './App.scss';

import React, { useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import ErrorPage from './components/ErrorPage';
import Graphs from './components/Graphs';
import Header from './components/Header';
import Home from './components/Home';
import Landing from './components/Landing';
import { StatsContext, UserContext } from './contexts';
import { FetchedData, Stats, User } from './types';
import { defaultState } from './utils/stateUtils';
import { initialUserState } from './contexts/UserContext';
import { initialStatsState } from './contexts/StatsContext';

export default function App(): JSX.Element {
  const [user, setUser] = useState<FetchedData<User>>(initialUserState);
  const [stats, setStats] = useState<FetchedData<Stats>>(initialStatsState);

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
