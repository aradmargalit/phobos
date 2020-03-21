import './Home.scss';

import { Form, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import {
  fetchActivities,
  fetchQuickAdds,
  fetchStatistics,
} from '../../apis/phobos-api';
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

  const [quickAdds, setQuickAdds] = useState(null);
  const [quickAddsLoading, setQuickAddsLoading] = useState(true);

  const [form] = Form.useForm();
  const setFormValues = values => form.setFieldsValue(values);

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
                fetchQuickAdds(setQuickAdds, setQuickAddsLoading);
              }}
              setQuickAdds={setQuickAdds}
              form={form}
            />
          </div>
          <div>
            <h4>Quick-Add</h4>
            <QuickAdd
              quickAdds={quickAdds}
              setQuickAdds={setQuickAdds}
              loading={quickAddsLoading}
              setLoading={setQuickAddsLoading}
              setQuickAdd={setFormValues}
            />
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
