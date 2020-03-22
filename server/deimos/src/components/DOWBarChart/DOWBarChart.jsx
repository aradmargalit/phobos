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
      <ResponsiveContainer width={500} height={300}>
        <BarChart
          data={dayBreakdown}
          margin={{
            top: 5,
            right: 30,
            left: 20,
            bottom: 5,
          }}
        >
          <XAxis
            dataKey="day_of_week"
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis />
          <Bar dataKey="count" fill="#0e5a6d">
            <LabelList dataKey="count" position="top" />
            {dayBreakdown.map((entry, index) => (
              <Cell
                key={`cell-${index}`}
                fill={colors[index % colors.length]}
              />
            ))}
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
}
