import './Home.scss';

import { Form, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect, withRouter } from 'react-router-dom';

import { fetchActivities, fetchQuickAdds, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext, UserContext } from '../../contexts';
import { makeDurationBreakdown } from '../../utils/durationUtils';
import { defaultState } from '../../utils/stateUtils';
import ActivityTable from '../ActivityTable';
import CreateActivity from '../CreateActivity';
import QuickAdd from '../QuickAdd';
import Statistics from '../Statistics';

function Home({ history }) {
  const { user } = useContext(UserContext);
  const { setStats } = useContext(StatsContext);

  const [activity, setActivity] = useState(null);

  const [activities, setActivities] = useState(defaultState());
  const [quickAdds, setQuickAdds] = useState(defaultState());

  const refetch = () => {
    fetchActivities(setActivities);
    fetchStatistics(setStats);
    fetchQuickAdds(setQuickAdds);
  };

  const [form] = Form.useForm();
  const setFormValues = values => {
    form.setFieldsValue({
      ...values,
      duration: makeDurationBreakdown(values.duration),
    });
    setActivity(form.getFieldsValue());
  };

  useEffect(() => {
    fetchActivities(setActivities);
    fetchQuickAdds(setQuickAdds);
  }, []);

  // If there are hard errors, our server is probably down, so render an error page
  if (user.errors) history.push('/error');

  // If either item is loading, show a spinner
  if (user.loading || activities.loading) return <Spin />;

  // If there's no user, but no errors, they're just not logged in
  if (!user.payload) return <Redirect to="/" />;

  return (
    <div className="app-content">
      <div className="container input-form">
        <h3>Add Activity</h3>
        <div className="input-form--contents">
          <div>
            <h4>Manual Add</h4>
            <CreateActivity
              form={form}
              refetch={refetch}
              activity={activity}
              setActivity={setActivity}
            />
          </div>
          <div>
            <h4>Quick-Add</h4>
            <QuickAdd
              quickAdds={quickAdds}
              setQuickAdds={setQuickAdds}
              setQuickAdd={setFormValues}
            />
          </div>
        </div>
      </div>
      <div className="container statistics">
        <h3>Your Statistics</h3>
        <Statistics activities={activities} />
      </div>
      <div className="container data-table">
        <h3>Your Activities</h3>
        <ActivityTable
          activities={activities.payload}
          refetch={() => {
            fetchActivities(setActivities);
            fetchStatistics(setStats);
          }}
          loading={activities.loading}
        />
      </div>
    </div>
  );
}

export default withRouter(Home);
