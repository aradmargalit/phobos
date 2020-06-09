import './IntervalGraph.scss';

import { Button, Spin } from 'antd';
import React, { useEffect, useState } from 'react';
import {
  Area,
  AreaChart,
  CartesianGrid,
  ReferenceArea,
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
  const highestPoint = Math.max(...data.map(d => d[dataKey]));
  const [top, setTop] = useState(
    Math.max(0, Math.ceil(Math.max(projection.y, highestPoint) * 1.1))
  );
  const [bottom, setBottom] = useState(0);
  const [left, setLeft] = useState('dataMin');
  const [right, setRight] = useState('dataMax');
  const [refLeft, setRefLeft] = useState('');
  const [refRight, setRefRight] = useState('');
  const [dataSlice, setDataSlice] = useState(data);

  const zoomOut = () => {
    setDataSlice(data);
    setLeft('dataMin');
    setRight('dataMax');
    setTop(Math.max(0, Math.ceil(Math.max(projection.y, highestPoint) * 1.1)));
    setBottom(0);
    setRefLeft('');
    setRefRight('');
  };

  const dataString = data.map(d => d[xAxisKey]).join(',');

  useEffect(() => {
    // Whenever something changes, zoom out just to be safe
    zoomOut();
  }, [dataString]); // Need to map array to a string for dep array to work

  const getAxisYDomain = (leftBound, rightBound) => {
    // Data is already sorted, so push everything from left to right into an array
    const newSlice = [];
    let hitLeft = false;
    let hitRight = false;
    data.forEach(point => {
      // If I'm done, return early
      if (hitRight) return;

      const xAxisValue = point[xAxisKey];

      // If we aren't yet at the left bound, continue until we do.
      if (!hitLeft) {
        if (xAxisValue !== leftBound) return;
        // If we haven't returned, we've hitleft!
        hitLeft = true;
      }

      newSlice.push(point);

      // Lastly, if this is the last point, cancel the for loop.
      if (xAxisValue === rightBound) {
        hitRight = true;
      }
    });

    setDataSlice(newSlice);
    return [0, Math.max(...newSlice.map(d => d[dataKey]))];
  };

  const zoomIn = () => {
    let newLeft = refLeft;
    let newRight = refRight;
    if (refLeft === refRight || refRight === '') {
      setRefLeft('');
      setRefRight('');
      return;
    }

    // If they drag right-to-left, swap them
    const leftIndex = data.findIndex(x => x[xAxisKey] === refLeft);
    const rightIndex = data.findIndex(x => x[xAxisKey] === refRight);
    if (rightIndex < leftIndex) {
      [newLeft, newRight] = [refRight, refLeft];
    }

    // yAxis domain
    const [newBottom, newTop] = getAxisYDomain(newLeft, newRight);

    setLeft(newLeft);
    setRight(newRight);
    setTop(newTop);
    setBottom(newBottom);
    setRefLeft('');
    setRefRight('');
  };

  if (loading) return <Spin />;

  return (
    <div className="interval-graph-wrapper">
      <div className="graph-header">
        <h2>{title}</h2>
      </div>
      <p>Drag to select a range to focus on.</p>
      <Button disabled={left === 'dataMin'} onClick={zoomOut}>
        Zoom Out
      </Button>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="interval-graph"
          data={dataSlice}
          margin={{ top: 30, right: 30, left: 30, bottom: 0 }}
          padding={{ top: 30, right: 30, left: 30, bottom: 10 }}
          onMouseDown={e => e && setRefLeft(e.activeLabel)}
          onMouseMove={e => e && refLeft && setRefRight(e.activeLabel)}
          onMouseUp={zoomIn}
        >
          <defs>
            <linearGradient id={`g-${dataKey}`} x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor={color} stopOpacity={0.6} />
              <stop offset="95%" stopColor={color} stopOpacity={0} />
            </linearGradient>
          </defs>
          <XAxis
            allowDataOverflow
            domain={[left, right]}
            interval={dataSlice.length >= 50 ? 10 : Math.ceil(dataSlice.length / 5)}
            dataKey={xAxisKey}
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis domain={[bottom, top]} />
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
            animationDuration={800}
          />
          {refLeft && refRight ? (
            <ReferenceArea x1={refLeft} x2={refRight} strokeOpacity={0.3} />
          ) : null}
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
