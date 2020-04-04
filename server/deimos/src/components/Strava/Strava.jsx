import { Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchStravaStats } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';

export default function Strava() {
  const { user } = useContext(UserContext);

  const [stravaStats, setStravaStats] = useState(null);
  const [stravaLoading, setStravaLoading] = useState(true);

  useEffect(() => {
    fetchStravaStats(setStravaStats, setStravaLoading);
  }, [setStravaLoading]);

  if (user.loading || stravaLoading) return <Spin />;
  if (!user.payload) return <Redirect to="/" />;

  return <div>{stravaStats.all_swim_totals.count}</div>;
}
