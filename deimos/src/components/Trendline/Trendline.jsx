import './Trendline.scss';

import { Tooltip } from 'antd';
import React from 'react';
import { Line, LineChart, ResponsiveContainer } from 'recharts';

const makeGraphData = arr => arr.map((datum, idx) => ({ idx, datum }));

export default function Trendline({ trendData }) {
  return (
    <div className="trendline">
      <Tooltip title="Each day's workout total minutes for the past 10 days.">
        <h3>Last 10 Days</h3>
      </Tooltip>
      <ResponsiveContainer id="trendline-wrapper" width="100%">
        <LineChart data={makeGraphData(trendData)}>
          <Line type="monotone" dataKey="datum" stroke="#0e5a6d" strokeWidth={2} dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
