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
            {dayBreakdown.map((day, index) => (
              <Cell
                key={day.day_of_week}
                fill={colors[index % colors.length]}
              />
            ))}
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
}
