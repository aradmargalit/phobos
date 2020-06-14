import { meanBy as _meanBy, startCase as _startCase } from 'lodash';
import React from 'react';

import IntervalGraph from '../IntervalGraph';

const average = data => _meanBy(data, 'rawPercentage');

export default function SkippedGraph({ loading, intervalData, intervalType }) {
  const data = intervalData.map(({ interval, days_skipped: daysSkipped }) => ({
    interval,
    rawPercentage: parseFloat(daysSkipped),
    percentage: parseFloat(daysSkipped).toFixed(2),
  }));

  const startCaseIntervalType = _startCase(intervalType);

  return (
    <IntervalGraph
      loading={loading}
      data={data}
      average={average(data)}
      title={`Percentage of Days per ${startCaseIntervalType} Active`}
      color="#9055A2"
      stroke="#2E294E"
      xAxisKey="interval"
      dataKey="percentage"
      unit={startCaseIntervalType}
      tooltipFormatter={value => [`${value}% of Days with a Workout`, '']}
      fixedTop={100}
    />
  );
}
