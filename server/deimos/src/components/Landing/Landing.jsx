/* eslint-disable global-require */

import './Landing.scss';

import { GoogleOutlined, LoginOutlined } from '@ant-design/icons';
import { Button, Spin } from 'antd';
import React, { useContext } from 'react';
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
        <img
          className="landing-container__hero"
          src={require('./moon.png')}
          alt="rocket"
        />
        <div className="landing-container__actions">
          <Button
            href="/auth/google"
            className="ant-btn big-button"
            icon={<GoogleOutlined />}
            type="primary"
          >
            Sign Up with Google
          </Button>
          <i>already have an account?</i>
          <Button
            href="/auth/google"
            className="ant-btn small-button"
            icon={<LoginOutlined />}
            ghost
            type="primary"
          >
            Sign In
          </Button>
        </div>
      </div>
    </div>
  );
}
