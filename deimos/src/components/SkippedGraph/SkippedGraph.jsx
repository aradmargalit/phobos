import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';

import MonthlyGraph from '../MonthlyGraph';

const average = data => _meanBy(data, 'skipped');

const projection = data =>
  (moment().daysInMonth() * data[data.length - 1].skipped) / moment(new Date()).date();

export default function SkippedGraph({ loading, intervalData, intervalType }) {
  const data = intervalData.map(({ interval, days_skipped: daysSkipped }) => ({
    interval,
    skipped: parseFloat(daysSkipped),
  }));

  console.log(projection(data));

  return (
    <MonthlyGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].interval, y: projection(data) }}
      title="Monthly Days Skipped"
      color="#9055A2"
      stroke="#2E294E"
      xAxisKey="interval"
      dataKey="skipped"
      unit={intervalType}
      tooltipFormatter={value => [`${value} Days Skipped`, '']}
    />
  );
}
