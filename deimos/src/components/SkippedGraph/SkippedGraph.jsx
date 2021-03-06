import { meanBy as _meanBy, startCase as _startCase } from 'lodash';
import React from 'react';

import IntervalGraph from '../IntervalGraph';

const average = data => _meanBy(data, 'percentage');

export default function SkippedGraph({ loading, intervalData, intervalType, goals, setGoals }) {
  const data = intervalData.map(({ interval, percentage_active: percentageActive }) => ({
    interval,
    percentage: parseFloat(percentageActive),
  }));

  const goal = goals
    ? goals.find(
        g => g.period === intervalType.toLowerCase() && g.metric === '% of days with a workout'
      )
    : null;

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
      metricName="% of Days with a Workout"
      goalDot={goal && { x: data[data.length - 1].interval, y: goal.goal }}
      setGoals={setGoals}
      currentGoal={goal}
      unit={startCaseIntervalType}
      fixedTop={100}
    />
  );
}
