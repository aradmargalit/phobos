import './Home.scss';

import { Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchActivities, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext, UserContext } from '../../contexts';
import ActivityTable from '../ActivityTable';
import AddActivityForm from '../AddActivityForm';
import QuickAdd from '../QuickAdd';
import Statistics from '../Statistics';

export default function Home() {
  const { user, loading } = useContext(UserContext);
  const { statsLoading, setStats, setStatsLoading } = useContext(StatsContext);

  const [activities, setActivities] = useState(null);
  const [activityLoading, setActivityLoading] = useState(true);

  useEffect(() => {
    fetchActivities(setActivities, setActivityLoading);
  }, [setActivityLoading]);

  if (loading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return (
    <div className="app-content">
      <div className="container input-form">
        <h3>Add Activity</h3>
        <div className="input-form--contents">
          <div>
            <h4>Manual Add</h4>
            <AddActivityForm
              refetch={() => {
                fetchActivities(setActivities, setActivityLoading);
                fetchStatistics(setStats, setStatsLoading);
              }}
            />
          </div>
          <div>
            <h4>Quick-Add</h4>
            <QuickAdd />
          </div>
        </div>
      </div>
      <div className="container statistics">
        <h3>Your Statistics</h3>
        <Statistics loading={statsLoading} setLoading={setStatsLoading} />
      </div>
      <div className="container data-table">
        <h3>Your Activities</h3>
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
  );
}
