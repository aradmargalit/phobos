import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';

import MonthlyGraph from '../MonthlyGraph';

const average = data => _meanBy(data, 'skipped');

const projection = data =>
  (moment().daysInMonth() * data[data.length - 1].skipped) /
  moment(new Date()).date();

export default function SkippedGraph({ loading, intervalData }) {
  const data = intervalData.map(({ month, days_skipped: daysSkipped }) => ({
    month,
    skipped: parseFloat(daysSkipped),
  }));

  return (
    <MonthlyGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].month, y: projection(data) }}
      title="Monthly Days Skipped"
      color="#9055A2"
      stroke="#2E294E"
      xAxisKey="month"
      dataKey="skipped"
      unit="Month"
      tooltipFormatter={value => [`${value} Days Skipped`, '']}
    />
  );
}
