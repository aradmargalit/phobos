import './Header.scss';

import { Button, PageHeader } from 'antd';
import React from 'react';
import { withRouter } from 'react-router-dom';

import IdentityButton from '../IdentityButton';

const extra = () => (
  <div className="header__extra">
    <Button className="header__strava-button" href="/strava/auth">
      Connect with Strava
    </Button>
    <IdentityButton />
  </div>
);

const Header = ({ history, showBack, showStrava }) => {
  return (
    <PageHeader
      className="header"
      onBack={showBack ? () => history.push('/home') : null}
      title="PHOBOS"
      subTitle="A Fitness Tracker"
      extra={showStrava ? extra() : <IdentityButton />}
    />
  );
};

export default withRouter(Header);
