import './Graphs.scss';

import { Spin } from 'antd';
import React, { useContext, useEffect } from 'react';

import { fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import ActivityGraph from '../ActivityGraph';
import DOWBarChart from '../DOWBarChart';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';

const COLORS = ['#bd4946', '#d4524e', '#ec5b57', '#f07c79', '#f49d9a'];

export default function Graphs() {
  const { stats, setStats, statsLoading, setStatsLoading } = useContext(
    StatsContext
  );

  useEffect(() => {
    fetchStatistics(setStats, setStatsLoading);
  }, [setStats, setStatsLoading]);

  if (statsLoading) return <Spin />;

  const { day_breakdown: dayBreakdown, type_breakdown: typeBreakdown } = stats;

  return (
    <div className="graphs">
      <ActivityGraph />
      <DOWBarChart colors={COLORS} dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph colors={COLORS} typeBreakdown={typeBreakdown} />
    </div>
  );
}
