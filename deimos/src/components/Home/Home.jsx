import './Home.scss';

import { Spin } from 'antd';
import React, { useContext } from 'react';
import { Redirect } from 'react-router-dom';

import UserContext from '../../contexts';
import ActivityTable from '../ActivityTable';
import AddActivityForm from '../AddActivityForm';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <div>
      <div className="container">
        <h2>Add Activity</h2>
        <AddActivityForm />
      </div>
      <br />
      <div className="container data-table">
        <h2>Your Activities</h2>
        <ActivityTable />
      </div>
    </div>
  );
}
