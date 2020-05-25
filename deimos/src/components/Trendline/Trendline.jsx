import './Trendline.scss';

import { Tooltip } from 'antd';
import React from 'react';
import CountUp from 'react-countup';
import { Line, LineChart, ResponsiveContainer } from 'recharts';

const makeGraphData = arr => arr.map((datum, idx) => ({ idx, datum }));
const sumHours = arr => arr.reduce((accum, curr) => accum + curr, 0) / 60;

export default function Trendline({ trendData }) {
  return (
    <div className="trendline">
      <Tooltip title="Each day's workout hours for the past 10 days.">
        <div className="trendline__label">
          <h3>Last 10 Days:</h3>
          <CountUp end={sumHours(trendData)} decimals={2} decimal="." duration={2.5} />
          <h3>hours total</h3>
        </div>
      </Tooltip>
      <ResponsiveContainer id="trendline-wrapper" width="100%">
        <LineChart data={makeGraphData(trendData)}>
          <Line type="monotone" dataKey="datum" stroke="#0e5a6d" strokeWidth={2} dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
