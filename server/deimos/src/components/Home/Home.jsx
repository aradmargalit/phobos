import './Home.scss';

import { Button, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchActivities, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext, UserContext } from '../../contexts';
import ActivityTable from '../ActivityTable';
import AddActivityForm from '../AddActivityForm';
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
        <h3 className="home-header">Add Activity</h3>
        <div className="input-form--contents">
          <div>
            <h3>Manual Add</h3>
            <AddActivityForm
              refetch={() => {
                fetchActivities(setActivities, setActivityLoading);
                fetchStatistics(setStats, setStatsLoading);
              }}
            />
          </div>
          <div>
            <h3>Quick-Add</h3>
            <div className="saved-workouts-list">
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
              <Button>Saved Workout 1</Button>
            </div>
          </div>
        </div>
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
  );
}
