import React, { useContext } from 'react';
import { Spin } from 'antd';
import { Redirect } from 'react-router-dom';
import UserContext from '../../contexts';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <div>
      <h1>{`Welcome Home ${user.given_name}`}</h1>
    </div>
  );
}
