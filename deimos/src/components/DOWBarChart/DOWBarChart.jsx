/* eslint-disable react/no-array-index-key */
import { sortBy as _sortBy } from 'lodash';
import React from 'react';
import { Bar, BarChart, Cell, LabelList, XAxis, YAxis } from 'recharts';

import AngledGraphTick from '../AngledGraphTick';

const COLORS = ['#f7bdbc', '#f49d9a', '#f07c79', '#ec5b57', '#d4524e', '#bd4946', '#8e3734'];

const abbreviate = (dayBreakdown) =>
  dayBreakdown.map((day) => ({ ...day, day_of_week: day.day_of_week.slice(0, 3) }));

export default function DOWBarChart({ dayBreakdown }) {
  const sortedDays = _sortBy(dayBreakdown, 'count');

  // Find the proportion of each day, and give it a darker color if it's higher
  const calculateColor = (day) => COLORS[sortedDays.indexOf(day)];

  return (
    <div className="statistics--dow">
      <h3>Daily Breakdown</h3>
      <BarChart data={abbreviate(dayBreakdown)} width={500} height={300} margin={{ top: 20 }}>
        <XAxis dataKey="day_of_week" height={80} tick={<AngledGraphTick />} />
        <YAxis />
        <Bar dataKey="count" fill="#0e5a6d">
          <LabelList dataKey="count" position="top" />
          {/* This is exclusively for alternating colors in the graph */}
          {dayBreakdown.map((day) => (
            <Cell key={day.day_of_week} fill={calculateColor(day)} />
          ))}
        </Bar>
      </BarChart>
    </div>
  );
}
