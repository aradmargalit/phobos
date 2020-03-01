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
      <div className="container input-form">
        <h3>Add Activity</h3>
        <AddActivityForm />
      </div>
      <br />
      <div className="container data-table">
        <h3>Your Activities</h3>
        <ActivityTable />
      </div>
    </div>
  );
}
