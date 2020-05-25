/* eslint-disable react/no-array-index-key */
/* eslint-disable react/jsx-one-expression-per-line */
import './Statistics.scss';

import { CheckOutlined, ClockCircleOutlined, LineChartOutlined } from '@ant-design/icons';
import { Spin, Statistic } from 'antd';
import React, { useContext, useEffect } from 'react';
import { Link } from 'react-router-dom';

import { fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';
import Trendline from '../Trendline';

const iconTitle = (text, icon) => (
  <span>
    {text}
    {icon}
  </span>
);

export default function Statistics() {
  const { stats, setStats } = useContext(StatsContext);
  const { workouts, hours, miles } = stats.payload;

  useEffect(() => {
    fetchStatistics(setStats);
  }, [setStats]);

  if (stats.loading) return <Spin />;

  return (
    <Spin spinning={stats.loading}>
      <div className="stats">
        <div className="statistics-grid">
          <Statistic title={iconTitle('Total Workouts', <CheckOutlined />)} value={workouts} />
          <Statistic
            title={iconTitle('Total Hours Active', <ClockCircleOutlined />)}
            value={hours.toFixed(2)}
          />
          <Statistic
            title={iconTitle('Total Mileage', <LineChartOutlined />)}
            value={miles.toFixed(2)}
          />
          <Trendline />
          <Link className="ant-btn" to="/graph">
            More Graphs <LineChartOutlined style={{ marginLeft: '10px' }} />
          </Link>
        </div>
      </div>
    </Spin>
  );
}
