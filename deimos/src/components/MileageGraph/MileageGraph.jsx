import { meanBy as _meanBy, startCase as _startCase } from 'lodash';
import moment from 'moment';
import React from 'react';

import IntervalGraph from '../IntervalGraph';

const average = data => _meanBy(data, 'miles');
const adjustedWeekNumber = ((moment().day() + 6) % 7) + 1;

const projection = (data, intervalType) => {
  const running = data[data.length - 1].miles;
  switch (intervalType) {
    case 'month':
      return moment().daysInMonth() * (running / moment(new Date()).date());
    case 'year':
      return 365 * (running / moment().dayOfYear());
    case 'week':
      return 7 * (running / adjustedWeekNumber);
    default:
      return 0;
  }
};

export default function MileageGraph({ loading, intervalData, intervalType }) {
  const data = intervalData.map(({ interval, miles }) => ({
    interval,
    miles: parseFloat(miles.toFixed(2)),
  }));

  const startCaseIntervalType = _startCase(intervalType);

  return (
    <IntervalGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].interval, y: projection(data, intervalType) }}
      title={`${startCaseIntervalType}ly Workout Miles`}
      color="#d4504f"
      stroke="#912827"
      xAxisKey="interval"
      dataKey="miles"
      unit={startCaseIntervalType}
      tooltipFormatter={value => [`${value} Miles`, '']}
    />
  );
}
