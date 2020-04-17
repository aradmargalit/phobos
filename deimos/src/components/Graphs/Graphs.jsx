import './Graphs.scss';

import { Empty, Select, Spin } from 'antd';
import moment from 'moment';
import React, { useContext, useEffect, useState } from 'react';

import { fetchStatistics, fetchSummariesByInterval } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import DOWBarChart from '../DOWBarChart';
import DurationGraph from '../DurationGraph';
import MileageGraph from '../MileageGraph';
import RadialActivityTypesGraph from '../RadialActivityTypesGraph';
import SkippedGraph from '../SkippedGraph';

const { Option } = Select;

export default function Graphs() {
  const { stats, setStats } = useContext(StatsContext);
  const [interval, setInterval] = useState('month');

  const [intervalData, setIntervalData] = useState({
    payload: [],
    loading: true,
    errors: null,
  });

  useEffect(() => {
    fetchStatistics(setStats);
    fetchSummariesByInterval(setIntervalData, interval);
  }, [setStats, interval]);

  const loading = stats.loading || intervalData.loading;

  if (loading) return <Spin />;
  if (!intervalData.payload || !intervalData.payload.length || !stats.payload)
    return (
      <Empty
        style={{ marginTop: '50px' }}
        description="Come back after you've logged some activities!"
      />
    );

  const { day_breakdown: dayBreakdown, type_breakdown: typeBreakdown } = stats.payload;

  let data;
  if (interval === 'year' || interval === 'month') {
    data = intervalData.payload.sort(
      (a, b) => moment(new Date(a.interval)) - moment(new Date(b.interval))
    );
  } else {
    data = intervalData.payload.sort((a, b) => a.interval.localeCompare(b.interval));
  }

  return (
    <div className="graphs">
      <DOWBarChart dayBreakdown={dayBreakdown} />
      <RadialActivityTypesGraph typeBreakdown={typeBreakdown} />
      <div className="interval-selector">
        <h3>Summary Interval: </h3>

        <Select placeholder="Monthly" onChange={value => setInterval(value)}>
          <Option value="week">Weekly</Option>
          <Option value="month">Monthly</Option>
          <Option value="year">Yearly</Option>
        </Select>
      </div>
      <DurationGraph loading={loading} intervalData={data} intervalType={interval} />
      <MileageGraph loading={loading} intervalData={data} intervalType={interval} />
      <SkippedGraph loading={loading} intervalData={data} intervalType={interval} />
    </div>
  );
}
