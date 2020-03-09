import './Home.scss';

import { Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchActivities, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext, UserContext } from '../../contexts';
import ActivityTable from '../ActivityTable';
import AddActivityForm from '../AddActivityForm';
import Statistics from '../Statistics';

export default function Home() {
  const { user, loading } = useContext(UserContext);

  const [activities, setActivities] = useState(null);
  const [activityLoading, setActivityLoading] = useState(true);
  const [stats, setStats] = useState({
    workouts: 0, hours: 0, miles: 0, last_ten: [],
  });
  const [statsLoading, setStatsLoading] = useState(true);
  useEffect(() => {
    fetchActivities(setActivities, setActivityLoading);
  }, [setActivityLoading]);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <StatsContext.Provider value={{ stats, setStats }}>
      <div className="app-content">
        <div className="container input-form">
          <h3 className="home-header">Add Activity</h3>
          <AddActivityForm
            refetch={() => {
              fetchActivities(setActivities, setActivityLoading);
              fetchStatistics(setStats, setStatsLoading);
            }}
          />
        </div>
        <div className="container statistics">
          <h3 className="home-header">Your Statistics</h3>
          <Statistics loading={statsLoading} setLoading={setStatsLoading} />
        </div>
        <div className="container data-table">
          <h3 className="home-header">Your Activities</h3>
          <ActivityTable
            activities={activities}
            refetch={() => {
              fetchActivities(setActivities, setActivityLoading);
              fetchStatistics(setStats, setStatsLoading);
            }}
            loading={activityLoading}
          />
        </div>
      </div>
    </StatsContext.Provider>
  );
}
