import './ActivityGraph.scss';

import { Spin } from 'antd';
import React, { useEffect, useState } from 'react';
import {
  CartesianGrid,
  Line, LineChart, Tooltip, XAxis, YAxis,
} from 'recharts';

import { fetchMonthlySums } from '../../apis/phobos-api';

export default function ActivityGraph() {
  const [monthlyData, setMonthlyData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchMonthlySums(setMonthlyData, setLoading);
  }, [setLoading]);

  const transform = (data) => data.map(({ month, duration }) => (
    { month, duration: parseFloat((duration / 60).toFixed(2)) }
  ));

  return (
    <Spin spinning={loading}>
      <LineChart
        className="activity-graph"
        width={1000}
        height={500}
        data={transform(monthlyData)}
        margin={{
          top: 5, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="10 10" />
        <XAxis tick={false} dataKey="activity_date" />
        <YAxis />
        <Tooltip />
        <Line type="monotone" dataKey="duration" stroke="#de541e" dot={false} strokeWidth={3} />
      </LineChart>
    </Spin>

  );
}


/*
import React, { PureComponent } from 'react';
import {
  LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend,
} from 'recharts';

const data = [
  {
    name: 'Page A', uv: 4000, pv: 2400, amt: 2400,
  },
  {
    name: 'Page B', uv: 3000, pv: 1398, amt: 2210,
  },
  {
    name: 'Page C', uv: 2000, pv: 9800, amt: 2290,
  },
  {
    name: 'Page D', uv: 2780, pv: 3908, amt: 2000,
  },
  {
    name: 'Page E', uv: 1890, pv: 4800, amt: 2181,
  },
  {
    name: 'Page F', uv: 2390, pv: 3800, amt: 2500,
  },
  {
    name: 'Page G', uv: 3490, pv: 4300, amt: 2100,
  },
];

export default class Example extends PureComponent {
  static jsfiddleUrl = 'https://jsfiddle.net/alidingling/xqjtetw0/';

  render() {
    return (
      <LineChart
        width={500}
        height={300}
        data={data}
        margin={{
          top: 5, right: 30, left: 20, bottom: 5,
        }}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis />
        <Tooltip />
        <Legend />
        <Line type="monotone" dataKey="pv" stroke="#8884d8" activeDot={{ r: 8 }} />
        <Line type="monotone" dataKey="uv" stroke="#82ca9d" />
      </LineChart>
    );
  }
}
*/
