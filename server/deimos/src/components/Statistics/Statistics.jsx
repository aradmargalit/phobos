/* eslint-disable react/no-array-index-key */
/* eslint-disable react/jsx-one-expression-per-line */
import './Statistics.scss';

import {
  CheckOutlined,
  ClockCircleOutlined,
  LineChartOutlined,
} from '@ant-design/icons';
import { Spin, Statistic } from 'antd';
import React, { useContext, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Line, LineChart, ResponsiveContainer } from 'recharts';

import { fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';

const makeGraphData = arr => arr.map((datum, idx) => ({ idx, datum }));
const iconTitle = (text, icon) => (
  <span>
    {text}
    {icon}
  </span>
);

export default function Statistics() {
  const { stats, setStats } = useContext(StatsContext);
  const { workouts, hours, miles, last_ten: lastTen } = stats.payload;

  useEffect(() => {
    fetchStatistics(stats, setStats);
  }, []);

  if (stats.loading) return <Spin />;

  return (
    <Spin spinning={stats.loading}>
      <div className="stats">
        <div className="statistics-grid">
          <Statistic
            title={iconTitle('Total Workouts', <CheckOutlined />)}
            value={workouts}
          />
          <Statistic
            title={iconTitle('Total Hours Active', <ClockCircleOutlined />)}
            value={hours.toFixed(2)}
          />
          <Statistic
            title={iconTitle('Total Mileage', <LineChartOutlined />)}
            value={miles.toFixed(2)}
          />
          <div className="statistics--trendline">
            <h3>Last 10 Days</h3>
            <ResponsiveContainer id="trendline-wrapper" width="100%">
              <LineChart data={makeGraphData(lastTen)}>
                <Line
                  type="monotone"
                  dataKey="datum"
                  stroke="#0e5a6d"
                  strokeWidth={2}
                  dot={false}
                />
              </LineChart>
            </ResponsiveContainer>
            <Link style={{ marginTop: '20px' }} className="ant-btn" to="/graph">
              More Graph <LineChartOutlined style={{ marginLeft: '10px' }} />
            </Link>
          </div>
        </div>
      </div>
    </Spin>
  );
}
