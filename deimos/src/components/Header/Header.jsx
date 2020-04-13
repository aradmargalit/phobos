import './Header.scss';

import { PageHeader } from 'antd';
import React, { useContext } from 'react';
import { withRouter } from 'react-router-dom';

import { UserContext } from '../../contexts';
import IdentityButton from '../IdentityButton';
import StravaButton from '../StravaButton';

const extra = (registeredWithStrava, loading) => (
  <div className="header__extra">
    <StravaButton registered={registeredWithStrava} loading={loading} />
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
