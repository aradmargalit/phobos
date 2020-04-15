import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';

import MonthlyGraph from '../MonthlyGraph';

const average = data => _meanBy(data, 'miles');

const projection = data =>
  (moment().daysInMonth() * data[data.length - 1].miles) /
  moment(new Date()).date();

export default function MileageGraph({ loading, monthlyData }) {
  const data = monthlyData.map(({ month, miles }) => ({
    month,
    miles: parseFloat(miles.toFixed(2)),
  }));

  return (
    <MonthlyGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].month, y: projection(data) }}
      title="Monthly Workout Miles"
      color="#d4504f"
      stroke="#912827"
      xAxisKey="month"
      dataKey="miles"
      unit="Month"
      tooltipFormatter={value => [`${value} Miles`, '']}
    />
  );
}
