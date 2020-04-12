import './Graphs.scss';

import { Empty, Spin } from 'antd';
import React, { useContext, useEffect, useState } from 'react';

import { fetchMonthlySums, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import ActivityGraph from '../ActivityGraph';
import DOWBarChart from '../DOWBarChart';
import MileageGraph from '../MileageGraph';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';
import SkippedGraph from '../SkippedGraph';

export default function Graphs() {
  const { stats, setStats } = useContext(StatsContext);

  const [monthlyData, setMonthlyData] = useState({
    payload: [],
    loading: true,
    errors: null,
  });

  useEffect(() => {
    fetchStatistics(setStats);
    fetchMonthlySums(setMonthlyData);
  }, [setStats]);

  if (stats.loading || monthlyData.loading) return <Spin />;
  if (!monthlyData.payload || !monthlyData.payload.length || !stats.payload)
    return (
      <Empty
        style={{ marginTop: '50px' }}
        description="Come back after you've logged some activities!"
      />
    );

  const {
    day_breakdown: dayBreakdown,
    type_breakdown: typeBreakdown,
  } = stats.payload;

  return (
    <div className="graphs">
      <ActivityGraph
        loading={monthlyData.loading}
        monthlyData={monthlyData.payload}
      />
      <DOWBarChart dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph typeBreakdown={typeBreakdown} />
      <MileageGraph
        loading={monthlyData.loading}
        monthlyData={monthlyData.payload}
      />
      <SkippedGraph
        loading={monthlyData.loading}
        monthlyData={monthlyData.payload}
      />
    </div>
  );
}
