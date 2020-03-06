import './Home.scss';

import { Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchActivities, fetchActivityTypes } from '../../apis/phobos-api';
import UserContext from '../../contexts';
import ActivityGraph from '../ActivityGraph';
import ActivityTable from '../ActivityTable';
import AddActivityForm from '../AddActivityForm';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  const [activities, setActivities] = useState(null);
  const [activityTypes, setActivityTypes] = useState([]);
  const [activityLoading, setActivityLoading] = useState(true);

  useEffect(() => {
    fetchActivities(setActivities, setActivityLoading);
    fetchActivityTypes(setActivityTypes, setActivityLoading);
  }, [setActivityLoading]);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <div className="app-content">
      <div className="container input-form">
        <h3 className="home-header">Add Activity</h3>
        <AddActivityForm
          activityTypes={activityTypes}
          loading={activityLoading}
          refetch={() => fetchActivities(setActivities, setActivityLoading)}
        />
      </div>
      <div className="container activity-graph">
        <h3 className="home-header">Chorts</h3>
        <ActivityGraph />
      </div>
      <div className="container data-table">
        <h3 className="home-header">Your Activities</h3>
        <ActivityTable
          activityTypes={activityTypes}
          activities={activities}
          refetch={() => fetchActivities(setActivities, setActivityLoading)}
          loading={activityLoading}
        />
      </div>
    </div>
  );
}
