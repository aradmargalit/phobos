import './Header.scss';

import { PageHeader } from 'antd';
import React, { useContext } from 'react';
import { withRouter, RouteComponentProps } from 'react-router-dom';

import { UserContext } from '../../contexts';
import GitHubLink from '../GitHubLink';
import IdentityButton from '../IdentityButton';
import StravaButton from '../StravaButton';

interface HeaderProps extends RouteComponentProps {
  showBack?: boolean;
  showStrava?: boolean;
}

const Header = ({ history, showBack = false, showStrava = false }: HeaderProps) => {
  const { user } = useContext(UserContext);
  return (
    <PageHeader
      className="header"
      onBack={showBack ? () => history.push('/home') : undefined}
      title="PHOBOS"
      subTitle="A Fitness Tracker"
      extra={
        <div className="header__extra">
          <GitHubLink />
          {showStrava && (
            <StravaButton
              registered={user.payload && user.payload.hasStravaToken}
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
