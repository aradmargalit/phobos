import './Header.scss';

import { PageHeader } from 'antd';
import React, { useContext } from 'react';
import { withRouter } from 'react-router-dom';

import { UserContext } from '../../contexts';
import GitHubLink from '../GitHubLink';
import IdentityButton from '../IdentityButton';
import StravaButton from '../StravaButton';

const Header = ({ history, showBack, showStrava }) => {
  const { user } = useContext(UserContext);
  return (
    <PageHeader
      className="header"
      onBack={showBack ? () => history.push('/home') : null}
      title="PHOBOS"
      subTitle="A Fitness Tracker"
      extra={
        <div className="header__extra">
          <GitHubLink />
          {showStrava && (
            <StravaButton
              registered={user.payload && user.payload.strava_token}
              loading={user.loading}
            />
          )}
          <IdentityButton />
        </div>
      }
    />
  );
};

export default withRouter(Header);
