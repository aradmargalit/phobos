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

export default function Statistics({ loading, setLoading }) {
  const { stats, setStats } = useContext(StatsContext);
  const {
    workouts,
    hours,
    miles,
    last_ten: lastTen,
    // type_breakdown: typeBreakdown,
    // day_breakdown: dayBreakdown,
  } = stats;

  useEffect(() => {
    fetchStatistics(setStats, setLoading);
  }, [setStats, setLoading]);

  if (loading) return <Spin />;

  return (
    <Spin spinning={loading}>
      <div className="stats">
        <div className="statistics-grid">
          <Statistic
            title={
              <span>
                Total Workouts <CheckOutlined />
              </span>
            }
            value={workouts}
          />
          <Statistic
            title={
              <span>
                Total Hours Active <ClockCircleOutlined />
              </span>
            }
            value={hours.toFixed(2)}
          />
          <Statistic
            title={
              <span>
                Total Miles <LineChartOutlined />
              </span>
            }
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
