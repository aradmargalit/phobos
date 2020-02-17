import React from 'react';
import { UserContext } from '../../contexts/UserContext';
import { Redirect } from 'react-router-dom';

export default function Home() {
  return (
    <UserContext.Consumer>
      {({ user, loading }) =>
        loading ? (
          <p>Loading...</p>
        ) : !isEmpty(user) ? (
          <div>
            <h1>{`Welcome Home ${user.given_name}`}</h1>
          </div>
        ) : (
          <Redirect to="/" />
        )
      }
    </UserContext.Consumer>
  );
}

const isEmpty = obj => {
  const ie = Object.entries(obj).length === 0 && obj.constructor === Object;
  return ie;
};
