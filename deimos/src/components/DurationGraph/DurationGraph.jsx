import { meanBy as _meanBy, startCase as _startCase } from 'lodash';
import moment from 'moment';
import React from 'react';

import IntervalGraph from '../IntervalGraph';

const average = data => _meanBy(data, 'rawDuration') / 60;
const adjustedWeekNumber = ((moment().day() + 6) % 7) + 1;

const projection = (data, intervalType) => {
  const running = data[data.length - 1].duration;
  switch (intervalType) {
    case 'month':
      return moment().daysInMonth() * (running / moment(new Date()).date());
    case 'year':
      return (365 * running) / moment().dayOfYear();
    case 'week':
      return 7 * (running / adjustedWeekNumber);
    default:
      return 0;
  }
};

export default function DurationGraph({ loading, intervalData, intervalType, goals, setGoals }) {
  const data = intervalData.map(({ interval, duration }) => ({
    interval,
    rawDuration: duration,
    duration: parseFloat((duration / 60).toFixed(2)),
  }));

  const goal = goals
    ? goals.find(g => g.period === intervalType.toLowerCase() && g.metric === 'hours')
    : null;

  const startCaseIntervalType = _startCase(intervalType);

  return (
    <IntervalGraph
      loading={loading}
      data={data}
      average={average(data)}
      projection={{ x: data[data.length - 1].interval, y: projection(data, intervalType) }}
      title={`${startCaseIntervalType}ly Workout Hours`}
      color="#117088"
      stroke="#0e5a6d"
      xAxisKey="interval"
      dataKey="duration"
      metricName="Hours"
      goalDot={goal && { x: data[data.length - 1].interval, y: goal.goal }}
      setGoals={setGoals}
      currentGoal={goal}
      unit={startCaseIntervalType}
    />
  );
}
