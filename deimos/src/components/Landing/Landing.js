import React from 'react';
import { UserContext } from '../../contexts/UserContext';
import { Redirect } from 'react-router-dom';

export default function Landing() {
  return (
    <UserContext.Consumer>
      {({ user, loading }) =>
        loading ? (
          <p>Loading</p>
        ) : user ? (
          <div>
            <Redirect to="/home" />
          </div>
        ) : (
          <h1>Oh, you're not logged in</h1>
        )
      }
    </UserContext.Consumer>
  );
}
