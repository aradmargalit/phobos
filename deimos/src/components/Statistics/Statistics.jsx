import './Statistics.scss';

import { CheckOutlined, ClockCircleOutlined, LineChartOutlined } from '@ant-design/icons';
import { Button, Spin, Statistic } from 'antd';
import React, { useEffect, useState } from 'react';
import {
  Line, LineChart,
} from 'recharts';

import { fetchStatistics } from '../../apis/phobos-api';

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

export default function Statistics() {
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({ workouts: 0, hours: 0, miles: 0 });
  const { workouts, hours, miles } = stats;

  useEffect(() => {
    fetchStatistics(setStats, setLoading);
  }, [setLoading]);

  return (
    <Spin spinning={loading}>
      <div className="statistics">
        <div className="stats">
          <Statistic
            title={(
              <span>
                Total Workouts
                {' '}
                <CheckOutlined />
              </span>
          )}
            value={workouts}
          />
          <Statistic
            title={(
              <span>
                Total Hours Active
                {' '}
                <ClockCircleOutlined />
              </span>
          )}
            value={hours.toFixed(2)}
          />
          <Statistic
            title={(
              <span>
                Total Miles
                {' '}
                <LineChartOutlined />
              </span>
          )}
            value={miles.toFixed(2)}
          />
        </div>
        <LineChart width={300} height={100} data={data}>
          <Line type="monotone" dataKey="pv" stroke="#0e5a6d" strokeWidth={2} dot={false} />
        </LineChart>
        <Button href="http://localhost:3000/graph">More Graph?</Button>
      </div>
    </Spin>
  );
}
