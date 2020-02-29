import React, { useContext } from 'react';
import { Spin } from 'antd';
import { Redirect } from 'react-router-dom';
import UserContext from '../../contexts';
import AddActivityForm from '../AddActivityForm';
import './Home.scss';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <div className="form-container">
      <h2>Add Activity</h2>
      <AddActivityForm />
    </div>
  );
}
