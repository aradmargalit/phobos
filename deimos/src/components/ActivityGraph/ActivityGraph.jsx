import './ActivityGraph.scss';

import { Spin } from 'antd';
import { meanBy as _meanBy } from 'lodash';
import moment from 'moment';
import React, { useEffect, useState } from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
  ReferenceLine,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

import { fetchMonthlySums } from '../../apis/phobos-api';

const transform = data =>
  data
    // Todo, convert these to a moment-compliant format
    .sort((a, b) => moment(a.month) - moment(b.month))
    .map(({ month, duration }) => ({
      month,
      rawDuration: duration,
      duration: parseFloat((duration / 60).toFixed(2)),
    }));

const calculateAverage = data => _meanBy(data, 'rawDuration') / 60;
const calculateCurrentMonth = data =>
  (moment().daysInMonth() * data[data.length - 1].duration) /
  moment(new Date()).date();

export default function ActivityGraph() {
  const [monthlyData, setMonthlyData] = useState([]);
  const [loading, setLoading] = useState(true);

  const data = transform(monthlyData);

  useEffect(() => {
    fetchMonthlySums(setMonthlyData, setLoading);
  }, [setLoading]);

  if (loading) return <Spin />;
  return (
    <div>
      <div className="graph-header">
        <h2>M O N T H L Y</h2>
        <h2>W O R K O U T</h2>
        <h2>H O U R S</h2>
      </div>

      <AreaChart
        className="activity-graph"
        width={1250}
        height={500}
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
          tick={<CustomizedAxisTick />}
        />
        <YAxis />
        <CartesianGrid strokeDasharray="3 3" />
        <ReferenceLine
          y={calculateAverage(data)}
          stroke="red"
          strokeDasharray="3 3"
        />
        <ReferenceLine
          y={calculateCurrentMonth(data)}
          stroke="blue"
          strokeDasharray="3 3"
          label="Current Month Projection"
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
          label={({ x, y, stroke, value }) => (
            <text
              x={x}
              y={y}
              dy={-15}
              stroke={stroke}
              fontSize={12}
              textAnchor="middle"
              fill="gray"
            >
              {value}
            </text>
          )}
        />
      </AreaChart>
    </div>
  );
}

const CustomizedAxisTick = ({ x, y, payload }) => {
  return (
    <g transform={`translate(${x},${y})`}>
      <text
        x={0}
        y={0}
        dy={16}
        textAnchor="end"
        fill="#666"
        transform="rotate(-45)"
      >
        {payload.value}
      </text>
    </g>
  );
};
