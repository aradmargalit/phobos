import './Header.scss';

import { PageHeader } from 'antd';
import React, { useContext } from 'react';
import { withRouter } from 'react-router-dom';

import { UserContext } from '../../contexts';
import GitHubLink from '../GitHubLink';
import IdentityButton from '../IdentityButton';
import StravaButton from '../StravaButton';

const extra = (showStrava, registeredWithStrava, loading) => (
  <div className="header__extra">
    <GitHubLink />
    {showStrava && <StravaButton registered={registeredWithStrava} loading={loading} />}
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
      extra={extra(showStrava, user.payload && user.payload.strava_token, user.loading)}
    />
  );
};

export default withRouter(Header);
