import './Header.scss';

import { PageHeader } from 'antd';
import React from 'react';
import { withRouter } from 'react-router-dom';

import IdentityButton from '../IdentityButton';

const Header = ({ history, showBack }) => {
  return (
    <PageHeader
      className="header"
      onBack={showBack ? () => history.push('/home') : null}
      title="PHOBOS"
      subTitle="A Fitness Tracker"
      extra={[<IdentityButton key="idb" />]}
    />
  );
};

export default withRouter(Header);
