/* eslint-disable react/no-array-index-key */
import React from 'react';
import {
  Bar,
  BarChart,
  Cell,
  LabelList,
  ResponsiveContainer,
  XAxis,
  YAxis,
} from 'recharts';

import AngledGraphTick from '../AngledGraphTick';

export default function DOWBarChart({ colors, dayBreakdown }) {
  const total = dayBreakdown.reduce((acc, current) => acc + current.count, 0);

  // Find the proportion of each day, and give it a darker color if it's higher
  const calculateColor = day =>
    colors[Math.floor(((day.count / total) * 100) / (100 / colors.length))];

  return (
    <div className="statistics--dow">
      <h3>Daily Breakdown</h3>
      <ResponsiveContainer width="100%" height={300}>
        <BarChart data={dayBreakdown}>
          <XAxis
            dataKey="day_of_week"
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis />
          <Bar dataKey="count" fill="#0e5a6d">
            <LabelList dataKey="count" position="top" />
            {/* This is exclusively for alternating colors in the graph */}
            {dayBreakdown.map(day => (
              <Cell key={day.day_of_week} fill={calculateColor(day)} />
            ))}
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
}
