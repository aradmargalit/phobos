import './Graphs.scss';

import { Empty, Select, Spin } from 'antd';
import moment from 'moment';
import React, { useContext, useEffect, useState } from 'react';

import {
  fetchStatistics,
  fetchSummariesByInterval,
} from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import DOWBarChart from '../DOWBarChart';
import DurationGraph from '../DurationGraph';
import MileageGraph from '../MileageGraph';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';
import SkippedGraph from '../SkippedGraph';

const { Option } = Select;

export default function Graphs() {
  const { stats, setStats } = useContext(StatsContext);
  const [interval, setInterval] = useState('monthly');

  const [intervalData, setIntervalData] = useState({
    payload: [],
    loading: true,
    errors: null,
  });

  useEffect(() => {
    fetchStatistics(setStats);
    fetchSummariesByInterval(setIntervalData, interval);
  }, [setStats, interval]);

  if (stats.loading || intervalData.loading) return <Spin />;
  if (!intervalData.payload || !intervalData.payload.length || !stats.payload)
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

  const sortedData = intervalData.payload.sort(
    (a, b) => moment(new Date(a.month)) - moment(new Date(b.month))
  );

  return (
    <div className="graphs">
      <DOWBarChart dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph typeBreakdown={typeBreakdown} />
      <div className="interval-selector">
        <h3>Summary Interval: </h3>

        <Select placeholder="Monthly" onChange={value => setInterval(value)}>
          <Option value="weekly">Weekly</Option>
          <Option value="monthly">Monthly</Option>
          <Option value="yearly">Yearly</Option>
        </Select>
      </div>
      <DurationGraph loading={intervalData.loading} intervalData={sortedData} />
      <MileageGraph loading={intervalData.loading} intervalData={sortedData} />
      <SkippedGraph loading={intervalData.loading} intervalData={sortedData} />
    </div>
  );
}
