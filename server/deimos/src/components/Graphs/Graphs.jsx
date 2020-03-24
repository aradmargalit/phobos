import './Graphs.scss';

import { Empty, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';

import { fetchMonthlySums, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import ActivityGraph from '../ActivityGraph';
import DOWBarChart from '../DOWBarChart';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';

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
      <DOWBarChart dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph typeBreakdown={typeBreakdown} />
    </div>
  );
}
