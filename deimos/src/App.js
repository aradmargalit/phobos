import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import Header from './components/Header';
import './App.css';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <div>
          <Header />
          <Switch>
            <Route path="/"></Route>
          </Switch>
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
