/* eslint global-require: 0 */

import React, { useContext } from 'react';
import { Redirect } from 'react-router-dom';
import { Spin, Button } from 'antd';
import UserContext from '../../contexts';
import './Landing.scss';

export default function Landing() {
  const BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

  const { user, loading } = useContext(UserContext);
  if (loading) return <Spin />;
  if (user) return <Redirect to="/home" />;

  return (
    <div className="landing-container">
      <h1>P H O B O S</h1>
      <h3>a smarter fitness tracker</h3>
      <div className="landing-container__tray">
        <img
          className="landing-container__hero"
          src={require('./moon.png')}
          alt="rocket"
        />
        <div className="landing-container__actions">
          <Button href={`${BACKEND_URL}/auth/google`} className="ant-btn big-button" icon="google" type="primary">
            Sign Up with Google
          </Button>
          <i>already have an account?</i>
          <Button href={`${BACKEND_URL}/auth/google`} className="ant-btn small-button" icon="login" ghost type="primary">
            Sign In
          </Button>
        </div>
      </div>
    </div>
  );
}
