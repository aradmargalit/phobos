import './Graphs.scss';

import { Empty, Select, Spin } from 'antd';
import moment from 'moment';
import React, { useContext, useEffect, useState } from 'react';

import { fetchMonthlySums, fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import DOWBarChart from '../DOWBarChart';
import DurationGraph from '../DurationGraph';
import MileageGraph from '../MileageGraph';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';
import SkippedGraph from '../SkippedGraph';

const { Option } = Select;

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

  const sortedMonthlyData = monthlyData.payload.sort(
    (a, b) => moment(new Date(a.month)) - moment(new Date(b.month))
  );

  return (
    <div className="graphs">
      <DOWBarChart dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph typeBreakdown={typeBreakdown} />
      <div className="granularity-selector">
        <h3>Summary Interval: </h3>

        <Select placeholder="Monthly">
          <Option value="weekly">Weekly</Option>
          <Option value="monthly">Monthly</Option>
          <Option value="yearly">Yearly</Option>
        </Select>
      </div>
      <DurationGraph
        loading={monthlyData.loading}
        monthlyData={sortedMonthlyData}
      />
      <MileageGraph
        loading={monthlyData.loading}
        monthlyData={sortedMonthlyData}
      />
      <SkippedGraph
        loading={monthlyData.loading}
        monthlyData={sortedMonthlyData}
      />
    </div>
  );
}
