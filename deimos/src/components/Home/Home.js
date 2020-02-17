import React, { useContext } from 'react';
import { Spin } from 'antd';
import { UserContext } from '../../contexts';
import { Redirect } from 'react-router-dom';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  if (loading) return <Spin />;
  return user ? (
    <div>
      <h1>{`Welcome Home ${user.given_name}`}</h1>
    </div>
  ) : (
    <Redirect to="/" />
  );
}
