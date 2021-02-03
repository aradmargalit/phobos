import React from 'react';
import './Landing.scss';
import LandingMenu from './LandingMenu';

function Landing(): JSX.Element {
  return (
    <div className="landing">
      <LandingMenu />
    </div>
  );
}

export default Landing;
