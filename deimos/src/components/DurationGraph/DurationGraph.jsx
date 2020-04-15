import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';

import MonthlyGraph from '../MonthlyGraph';

const average = data => _meanBy(data, 'rawDuration') / 60;

const projection = data =>
  (moment().daysInMonth() * data[data.length - 1].duration) /
  moment(new Date()).date();

export default function DurationGraph({ loading, monthlyData }) {
  const data = monthlyData.map(({ month, duration }) => ({
    month,
    rawDuration: duration,
    duration: parseFloat((duration / 60).toFixed(2)),
  }));

  return (
    <MonthlyGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].month, y: projection(data) }}
      title="Monthly Workout Hours"
      color="#117088"
      stroke="#0e5a6d"
      xAxisKey="month"
      dataKey="duration"
      unit="Month"
      tooltipFormatter={value => [`${value} Hours`, '']}
    />
  );
}
