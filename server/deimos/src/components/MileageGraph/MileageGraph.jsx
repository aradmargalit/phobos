import './MileageGraph.scss';

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
    .map(({ month, miles }) => ({
      month,
      miles: parseFloat(miles.toFixed(2)),
    }));

const calculateAverage = data => _meanBy(data, 'miles');

export default function MileageGraph({ loading, monthlyData }) {
  if (loading) return <Spin />;
  const data = transform(monthlyData);

  return (
    <div className="activity-graph-wrapper">
      <div className="graph-header">
        <h2>M O N T H L Y</h2>
        <h2>M I L E A G E</h2>
      </div>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="activity-graph"
          data={data}
          margin={{ top: 10, right: 30, left: 0, bottom: 0 }}
          padding={{ top: 10, right: 30, left: 30, bottom: 10 }}
        >
          <defs>
            <linearGradient id="mileageColor" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#d4504f" stopOpacity={0.6} />
              <stop offset="95%" stopColor="#d4504f" stopOpacity={0} />
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
            stroke="red"
            strokeDasharray="3 3"
            label={{
              position: 'top',
              fontWeight: 600,
              value: 'Monthly Average',
            }}
          />
          <Tooltip
            separator={null}
            formatter={value => [`${value} miles`, '']}
            animationDuration={300}
          />
          <Area
            dataKey="miles"
            type="monotone"
            stroke="#912827"
            fillOpacity={1}
            fill="url(#mileageColor)"
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
