import { Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';
import { Redirect } from 'react-router-dom';

import { fetchStravaStats } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';

export default function Strava() {
  const { user, loading } = useContext(UserContext);

  const [stravaStats, setStravaStats] = useState(null);
  const [stravaLoading, setStravaLoading] = useState(true);

  useEffect(() => {
    fetchStravaStats(setStravaStats, setStravaLoading);
  }, [setStravaLoading]);

  if (loading || stravaLoading) return <Spin />;
  if (!user) return <Redirect to="/" />;

  return <div>{stravaStats.all_swim_totals.count}</div>;
}
