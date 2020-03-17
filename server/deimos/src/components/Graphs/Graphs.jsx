import './Graphs.scss';

import { Empty, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';

import { fetchMonthlySums, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import ActivityGraph from '../ActivityGraph';
import DOWBarChart from '../DOWBarChart';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';

const COLORS = ['#bd4946', '#d4524e', '#ec5b57', '#f07c79', '#f49d9a'];

export default function Graphs() {
  const { stats, setStats, statsLoading, setStatsLoading } = useContext(
    StatsContext
  );
  const [monthlyData, setMonthlyData] = useState([]);
  const [monthlyLoading, setMonthlyLoading] = useState(true);

  useEffect(() => {
    fetchStatistics(setStats, setStatsLoading);
    fetchMonthlySums(setMonthlyData, setMonthlyLoading);
  }, [setStats, setStatsLoading]);

  if (statsLoading || monthlyLoading) return <Spin />;
  if (!monthlyData || !monthlyData.length || !stats)
    return (
      <Empty
        style={{ marginTop: '50px' }}
        description="Come back after you've logged some activities!"
      />
    );

  const { day_breakdown: dayBreakdown, type_breakdown: typeBreakdown } = stats;

  return (
    <div className="graphs">
      <ActivityGraph loading={monthlyLoading} monthlyData={monthlyData} />
      <DOWBarChart colors={COLORS} dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph colors={COLORS} typeBreakdown={typeBreakdown} />
    </div>
  );
}
