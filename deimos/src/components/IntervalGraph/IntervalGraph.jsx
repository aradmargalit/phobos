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

// Determine X Axis interval
const getInterval = dataLength => {
  // If there are few points, we can show them all
  if (dataLength <= 10) {
    return 0;
  }

  // If there are between 10 and 75 items, tick every 5.
  if (dataLength > 10 && dataLength <= 75) {
    return 10;
  }

  // If there are more than 75, default to soem sane number, like 10
  return 15;
};

// Takes the highest point in the graph, adds a cushion, and rounding it off
const maxToCeiling = (max, fixedTop) => fixedTop || Math.max(Math.ceil(max * 1.1), 1);

const getAxisYDomain = (data, dataKey, xAxisKey, leftBound, rightBound, fixedTop) => {
  // Data is already sorted, so push everything from left to right into an array
  // Once we find the data point at the left bound, set hitLeft to true
  // Once we reach the data point at the right bound, set hitRight to true
  const newSlice = [];
  let hitLeft = false;
  let hitRight = false;

  // For each data point, either skip or push it to newSlice, depending on if it's "in bounds"
  data.forEach(point => {
    // If I'm done, because I already found the right bound, return early
    if (hitRight) return;

    const xAxisValue = point[xAxisKey];

    // If we aren't yet at the left bound, continue until we do.
    if (!hitLeft) {
      if (xAxisValue !== leftBound) return;
      // If we haven't returned, we've hitleft!
      hitLeft = true;
    }

    // If we didn't return earlier, we're neither before left bound nor ahead of right bound
    // which means it's safe to push in
    newSlice.push(point);

    // Lastly, if this is the last point, cancel the for loop by setting hitRight to true
    if (xAxisValue === rightBound) {
      hitRight = true;
    }
  });

  return [0, maxToCeiling(Math.max(...newSlice.map(d => d[dataKey])), fixedTop), newSlice];
};

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
  fixedTop,
}) {
  // Find the highest point in the graph, and set defaultTop to the MAX(highestPoint, projection)
  const highestPoint = Math.max(...data.map(d => d[dataKey]));
  const defaultTop = projection
    ? Math.max(0, maxToCeiling(Math.max(projection.y, highestPoint), fixedTop))
    : maxToCeiling(highestPoint, fixedTop);

  // 'dataMin' and 'dataMax' let recharts default to the left and right bounds of the data
  const initialState = {
    top: defaultTop,
    bottom: 0,
    left: 'dataMin',
    right: 'dataMax',
    refLeft: '',
    refRight: '',
    dataSlice: data,
  };

  const [state, setState] = useState(initialState);

  // For React to know if "data" has changed, it needs to either always be the same length
  // which is impossible, so join the data to a string.
  const dataString = data.map(d => d[xAxisKey]).join(',');

  useEffect(() => {
    // Whenever something changes, zoom out just to be safe
    setState(initialState);
    // eslint-disable-next-line
  }, [dataString]);

  if (loading) return <Spin />;

  // Helper to determine if we are zoomedIn
  const isZoomed = state.left !== 'dataMin';

  const zoomIn = () => {
    const { refLeft, refRight } = state;
    let newLeft = refLeft;
    let newRight = refRight;

    // If the bounds are the same, or there's no right bound yet, return and clear refs
    if (refLeft === refRight || refRight === '' || !refRight) {
      setState({ ...state, refLeft: '', refRight: '' });
      return;
    }

    // If they drag right-to-left, swap them so it's as if they dragged left to right
    const leftIndex = data.findIndex(x => x[xAxisKey] === refLeft);
    const rightIndex = data.findIndex(x => x[xAxisKey] === refRight);
    if (rightIndex < leftIndex) {
      [newLeft, newRight] = [refRight, refLeft];
    }

    // yAxis domain
    const [newBottom, newTop, newSlice] = getAxisYDomain(
      data,
      dataKey,
      xAxisKey,
      newLeft,
      newRight,
      fixedTop
    );

    setState({
      dataSlice: newSlice,
      left: newLeft,
      right: newRight,
      top: newTop,
      bottom: newBottom,
      refLeft: '',
      refRight: '',
    });
  };

  const { left, right, top, bottom, refLeft, refRight, dataSlice } = state;

  return (
    <div className="interval-graph-wrapper">
      <div className="graph-header">
        <h2>{title}</h2>
      </div>
      <p>Drag to select a range to focus on.</p>
      <Button disabled={!isZoomed} onClick={() => setState(initialState)}>
        Zoom Out
      </Button>
      <ResponsiveContainer width="100%" height={450}>
        <AreaChart
          className="interval-graph"
          data={dataSlice}
          margin={{ top: 30, right: 30, left: 30, bottom: 0 }}
          padding={{ top: 30, right: 30, left: 30, bottom: 10 }}
          onMouseDown={e => e && setState({ ...state, refLeft: e.activeLabel })}
          onMouseMove={e => e && refLeft && setState({ ...state, refRight: e.activeLabel })}
          onMouseLeave={() => setState({ ...state, refRight: '', refLeft: '' })}
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
            tickCount={getInterval(dataSlice.length)}
            dataKey={xAxisKey}
            height={120}
            tick={<AngledGraphTick />}
          />
          <YAxis domain={[bottom, top]} padding={{ top: fixedTop ? 10 : 0 }} />
          <CartesianGrid strokeDasharray="3 3" />
          {!isZoomed && (
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
          )}
          {!isZoomed && projection && (
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
          )}
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
