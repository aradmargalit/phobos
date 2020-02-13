import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Login from './components/Login';
import Header from './components/Header';
import './App.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <div>
          <Header />
          <Switch>
            <Route path="/">
              <Login />
            </Route>
          </Switch>
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
