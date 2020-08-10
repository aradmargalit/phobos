import './IdentityButton.scss';

import { LogoutOutlined } from '@ant-design/icons';
import { Button, Spin } from 'antd';
import React, { useContext, useEffect } from 'react';

import { fetchUser } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';

export default function IdentityButton() {
  const { user, setUser } = useContext(UserContext);

  useEffect(() => {
    fetchUser(setUser);
  }, [setUser]);

  if (user.loading) return <Spin />;

  return user.payload ? (
    <div>
      <h1 className="ant-page-header-heading-title welcome">Welcome back,</h1>
      <h1 className="ant-page-header-heading-title primary">{user.payload.given_name}</h1>
      <Button icon={<LogoutOutlined />} href="/users/logout" type="danger" ghost>
        Logout
      </Button>
    </div>
  ) : null;
}
