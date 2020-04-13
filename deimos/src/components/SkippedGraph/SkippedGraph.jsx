import './SkippedGraph.scss';

import { Spin } from 'antd';
import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
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
    .map(({ month, days_skipped: daysSkipped }) => ({
      month,
      skipped: parseFloat(daysSkipped),
    }));

const calculateAverage = data => _meanBy(data, 'skipped');

export default function SkippedGraph({ loading, monthlyData }) {
  if (loading) return <Spin />;
  const data = transform(monthlyData);

  return (
    <div className="activity-graph-wrapper">
      <div className="graph-header">
        <h2>M O N T H L Y</h2>
        <h2>D A Y S</h2>
        <h2>S K I P P E D</h2>
      </div>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="activity-graph"
          data={data}
          margin={{ top: 10, right: 30, left: 0, bottom: 0 }}
          padding={{ top: 10, right: 30, left: 30, bottom: 10 }}
        >
          <defs>
            <linearGradient id="skipColor" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#043927" stopOpacity={0.6} />
              <stop offset="95%" stopColor="#043927" stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis
            interval={2}
            dataKey="month"
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis />
          <CartesianGrid strokeDasharray="3 3" />
          <ReferenceLine
            y={calculateAverage(data)}
            stroke="#043927"
            strokeDasharray="3 3"
            label={{
              position: 'top',
              fontWeight: 600,
              value: 'Monthly Average',
            }}
          />
          <Tooltip
            separator={null}
            formatter={value => [`${value} days skipped`, '']}
            animationDuration={300}
          />
          <Area
            dataKey="skipped"
            type="monotone"
            stroke="#043927"
            fillOpacity={1}
            fill="url(#skipColor)"
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
