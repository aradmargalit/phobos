import { meanBy as _meanBy, startCase as _startCase } from 'lodash';
import moment from 'moment';
import React from 'react';

import IntervalGraph from '../IntervalGraph';

const average = data => _meanBy(data, 'skipped');

const projection = (data, intervalType) => {
  const running = data[data.length - 1].skipped;
  switch (intervalType) {
    case 'month':
      return moment().daysInMonth() * (running / moment(new Date()).date());
    case 'year':
      return 365 * (running / moment().dayOfYear());
    case 'week':
      return 7 * (running / (moment().day() + 1));
    default:
      return 0;
  }
};

export default function SkippedGraph({ loading, intervalData, intervalType }) {
  const data = intervalData.map(({ interval, days_skipped: daysSkipped }) => ({
    interval,
    skipped: parseFloat(daysSkipped),
  }));

  const startCaseIntervalType = _startCase(intervalType);

  return (
    <IntervalGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].interval, y: projection(data, intervalType) }}
      title={`${startCaseIntervalType}ly Days Skipped`}
      color="#9055A2"
      stroke="#2E294E"
      xAxisKey="interval"
      dataKey="skipped"
      unit={startCaseIntervalType}
      tooltipFormatter={value => [`${value} Days Skipped`, '']}
    />
  );
}
