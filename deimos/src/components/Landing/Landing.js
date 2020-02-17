import React, { useContext } from 'react';
import { UserContext } from '../../contexts';
import { Redirect } from 'react-router-dom';
import { Spin } from 'antd';

export default function Landing() {
  const { user, loading } = useContext(UserContext);
  if (loading) return <Spin />;

  return user ? (
    <div>
      <Redirect to="/home" />
    </div>
  ) : (
    <h1>Oh, you're not logged in</h1>
  );
}
