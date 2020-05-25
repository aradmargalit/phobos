import './Trendline.scss';

import { Select, Spin } from 'antd';
import React, { useEffect, useState } from 'react';
import CountUp from 'react-countup';
import { Line, LineChart, ResponsiveContainer } from 'recharts';

import { fetchTrendPoints } from '../../apis/phobos-api';
import { defaultState } from '../../utils/stateUtils';

const { Option } = Select;

const makeGraphData = arr => arr.map((datum, idx) => ({ idx, datum }));
const sumHours = arr => arr.reduce((accum, curr) => accum + curr, 0) / 60;

export default function Trendline({ activityTimtestamp }) {
  const [trendPoints, setTrendPoints] = useState(defaultState());
  const [lookback, setLookback] = useState('l10');

  useEffect(() => {
    fetchTrendPoints(setTrendPoints, lookback);
  }, [lookback, activityTimtestamp]);

  if (trendPoints.loading) return <Spin />;

  return (
    <div className="trendline">
      <div className="trendline__label">
        <Select defaultValue="l10" onChange={v => setLookback(v)}>
          <Option value="l10">Last 10 Days</Option>
          <Option value="l7">Last 7 Days</Option>
          {/* Disable if it's the first day, since a single point isn't a trend */}
          <Option value="lw" disabled={new Date().getDay() === 1}>
            This Week
          </Option>
          <Option value="lm" disabled={new Date().getDate() === 1}>
            This Month
          </Option>
        </Select>
        <CountUp end={sumHours(trendPoints.payload)} decimals={2} decimal="." duration={2.5} />
        <h3>hours total</h3>
      </div>
      <ResponsiveContainer id="trendline-wrapper" width="100%">
        <LineChart data={makeGraphData(trendPoints.payload)}>
          <Line type="monotone" dataKey="datum" stroke="#0e5a6d" strokeWidth={2} dot={false} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
