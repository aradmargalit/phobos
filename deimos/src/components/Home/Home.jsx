import React, { useContext } from 'react';
import { Spin } from 'antd';
import { Redirect } from 'react-router-dom';
import UserContext from '../../contexts';
import AddActivity from '../AddActivityModal';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  if (!user) return <Redirect to="/" />;
  if (loading) return <Spin />;

  return (
    <div>
      <h1><AddActivity /></h1>
    </div>
  );
}
