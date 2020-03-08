import './Statistics.scss';

import { CheckOutlined, ClockCircleOutlined, LineChartOutlined } from '@ant-design/icons';
import { Button, Spin, Statistic } from 'antd';
import React, { useEffect, useState } from 'react';
import {
  Line, LineChart,
} from 'recharts';

import { fetchStatistics } from '../../apis/phobos-api';

const makeGraphData = (arr) => arr.map((datum, idx) => ({ idx, datum }));

export default function Statistics() {
  const [loading, setLoading] = useState(true);
  const [stats, setStats] = useState({
    workouts: 0, hours: 0, miles: 0, last_ten: [],
  });
  const {
    workouts, hours, miles, last_ten: lastTen,
  } = stats;

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
        <h3>Last 10 Days</h3>
        <LineChart width={300} height={100} data={makeGraphData(lastTen)}>
          <Line type="monotone" dataKey="datum" stroke="#0e5a6d" strokeWidth={2} dot={false} />
        </LineChart>
        <Button href="http://localhost:3000/graph">More Graph?</Button>
      </div>
    </Spin>
  );
}
