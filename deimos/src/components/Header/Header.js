import React from 'react';
import { PageHeader } from 'antd';
import IdentityButton from '../IdentityButton';
import './Header.scss';

export default function Header() {
  return (
    <PageHeader
      className="header"
      title="Phobos"
      subTitle="A Fitness Tracker"
      extra={[<IdentityButton key="idb" />]}
    />
  );
}
