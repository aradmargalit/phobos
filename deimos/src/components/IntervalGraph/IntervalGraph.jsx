import './IntervalGraph.scss';

import { Spin } from 'antd';
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

export default function IntervalGraph({
  loading,
  data,
  average,
  projection,
  title,
  color,
  stroke,
  dataKey,
  xAxisKey,
  unit,
  tooltipFormatter,
}) {
  if (loading) return <Spin />;

  return (
    <div className="interval-graph-wrapper">
      <div className="graph-header">
        <h2>{title}</h2>
      </div>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="interval-graph"
          data={data}
          margin={{ top: 30, right: 30, left: 30, bottom: 0 }}
          padding={{ top: 30, right: 30, left: 30, bottom: 10 }}
          syncId="trulycouldnotmatterless"
        >
          <defs>
            <linearGradient id={`g-${dataKey}`} x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor={color} stopOpacity={0.6} />
              <stop offset="95%" stopColor={color} stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis
            interval={data.length >= 50 ? 10 : Math.ceil(data.length / 5)}
            dataKey={xAxisKey}
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis />
          <CartesianGrid strokeDasharray="3 3" />
          <ReferenceLine
            y={average}
            stroke={stroke}
            strokeDasharray="3 3"
            label={{
              position: 'top',
              fontWeight: 600,
              value: `${unit}ly Average: ${average.toFixed(1)}`,
            }}
          />
          <ReferenceDot
            x={projection.x}
            y={projection.y}
            stroke={stroke}
            strokeDasharray="3 3"
            label={{
              position: 'left',
              fontWeight: 600,
              value: `This ${unit}'s Projection: ${projection.y.toFixed(1)}`,
            }}
          />
          <Tooltip separator={null} formatter={tooltipFormatter} animationDuration={300} />
          <Area
            dataKey={dataKey}
            type="monotone"
            stroke={stroke}
            fillOpacity={1}
            fill={`url(#${`g-${dataKey}`})`}
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
