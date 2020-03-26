import './ActivityGraph.scss';

import { Spin } from 'antd';
import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
  ReferenceDot,
  ReferenceLine,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

import AngledGraphTick from '../AngledGraphTick';

const transform = data =>
  data
    .sort((a, b) => moment(new Date(a.month)) - moment(new Date(b.month)))
    .map(({ month, duration }) => ({
      month,
      rawDuration: duration,
      duration: parseFloat((duration / 60).toFixed(2)),
    }));

const calculateAverage = data => _meanBy(data, 'rawDuration') / 60;
const calculateCurrentMonth = currentMonth =>
  (moment().daysInMonth() * currentMonth.duration) / moment(new Date()).date();

export default function ActivityGraph({ loading, monthlyData }) {
  if (loading) return <Spin />;
  const data = transform(monthlyData);
  const currentMonth = data[data.length - 1];

  return (
    <div className="activity-graph-wrapper">
      <div className="graph-header">
        <h2>M O N T H L Y</h2>
        <h2>W O R K O U T</h2>
        <h2>H O U R S</h2>
      </div>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="activity-graph"
          data={data}
          margin={{ top: 10, right: 30, left: 0, bottom: 0 }}
          padding={{ top: 10, right: 30, left: 30, bottom: 10 }}
        >
          <defs>
            <linearGradient id="durationColor" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#117088" stopOpacity={0.6} />
              <stop offset="95%" stopColor="#117088" stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis
            interval={3}
            dataKey="month"
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis />
          <CartesianGrid strokeDasharray="3 3" />
          <ReferenceLine
            y={calculateAverage(data)}
            stroke="red"
            strokeDasharray="3 3"
            label={{
              position: 'top',
              fontWeight: 600,
              value: 'Monthly Average',
            }}
          />
          <ReferenceDot
            x={currentMonth.month}
            y={calculateCurrentMonth(currentMonth)}
            stroke="blue"
            strokeDasharray="3 3"
            label={{
              position: 'left',
              fontWeight: 600,
              value: "This Month's Projection",
            }}
          />
          <Tooltip
            separator={null}
            formatter={value => [`${value} Hours`, '']}
            animationDuration={300}
          />
          <Area
            dataKey="duration"
            type="monotone"
            stroke="#0e5a6d"
            fillOpacity={1}
            fill="url(#durationColor)"
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
