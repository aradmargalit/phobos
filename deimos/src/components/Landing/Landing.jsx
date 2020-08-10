/* eslint-disable global-require */

import './Landing.scss';

import { Spin } from 'antd';
import React, { useContext } from 'react';
import GoogleButton from 'react-google-button';
import { Redirect } from 'react-router-dom';

import { UserContext } from '../../contexts';

export default function Landing() {
  const { user } = useContext(UserContext);
  if (user.loading) return <Spin />;
  if (user.payload) return <Redirect to="/home" />;

  return (
    <div className="landing-container">
      <h1>P H O B O S</h1>
      <h3>a smarter fitness tracker</h3>
      <div className="landing-container__tray">
        <img className="landing-container__hero" src={require('./phobos.png')} alt="phobos" />
        <div className="landing-container__actions">
          <a href="/auth/google">
            {/* Use an off-the-shelf styled component to be compliant with Google's styling */}
            <GoogleButton href="/auth/google">Sign In with Google</GoogleButton>
          </a>
        </div>
      </div>
    </div>
  );
}
