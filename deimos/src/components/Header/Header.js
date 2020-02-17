import React from 'react';
import { PageHeader } from 'antd';
import IdentityButton from '../IdentityButton';
import { UserContext } from '../../contexts/UserContext';
import './Header.scss';

export default function Header() {
  return (
    <UserContext.Consumer>
      {({ user, setUser, loading, setLoading }) => (
        <PageHeader
          className="header"
          title="Phobos"
          subTitle="A Fitness Tracker"
          extra={[
            <IdentityButton
              key="idb"
              user={user}
              setUser={setUser}
              loading={loading}
              setLoading={setLoading}
            />
          ]}
        />
      )}
    </UserContext.Consumer>
  );
}
