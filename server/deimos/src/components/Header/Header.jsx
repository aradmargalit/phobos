import './Header.scss';

import { CheckCircleOutlined } from '@ant-design/icons';
import { Button, PageHeader } from 'antd';
import React, { useContext } from 'react';
import { withRouter } from 'react-router-dom';

import { UserContext } from '../../contexts';
import IdentityButton from '../IdentityButton';

const extra = (registeredWithStrava, loading) => (
  <div className="header__extra">
    <Button
      disabled={loading || registeredWithStrava}
      className={`header__strava-button${
        registeredWithStrava ? '--registered' : ''
      }`}
      href="/strava/auth"
    >
      {registeredWithStrava ? 'Connected with Strava' : 'Connect with Strava'}
      {registeredWithStrava && <CheckCircleOutlined />}
    </Button>
    <IdentityButton />
  </div>
);

const Header = ({ history, showBack, showStrava }) => {
  const { user } = useContext(UserContext);
  return (
    <PageHeader
      className="header"
      onBack={showBack ? () => history.push('/home') : null}
      title="PHOBOS"
      subTitle="A Fitness Tracker"
      extra={
        showStrava ? (
          extra(user.payload && user.payload.strava_token, user.loading)
        ) : (
          <IdentityButton />
        )
      }
    />
  );
};

export default withRouter(Header);
