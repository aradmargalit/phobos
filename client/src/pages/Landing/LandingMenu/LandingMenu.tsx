import React from 'react';
import './LandingMenu.scss';

function LandingMenu(): JSX.Element {
  return (
    <div className="landing-menu">
      <button type="button" className="home-button">
        phobos
      </button>
      <button type="button" className="login-button outlined">
        Log In
      </button>
    </div>
  );
}

export default LandingMenu;
