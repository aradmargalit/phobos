/* eslint-disable react/jsx-one-expression-per-line */
import './Statistics.scss';

import {
  CheckOutlined,
  ClockCircleOutlined,
  LineChartOutlined,
} from '@ant-design/icons';
import { Empty, Spin, Statistic } from 'antd';
import React, { useContext, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Cell, Line, LineChart, Pie, PieChart } from 'recharts';

import { fetchStatistics } from '../../apis/phobos-api';
import { StatsContext } from '../../contexts';

const makeGraphData = arr => arr.map((datum, idx) => ({ idx, datum }));

const COLORS = ['#0088FE', '#00C49F', '#FFBB28', '#FF8042'];

const renderLabel = d => d.name;

export default function Statistics({ loading, setLoading }) {
  const { stats, setStats } = useContext(StatsContext);
  const {
    workouts,
    hours,
    miles,
    last_ten: lastTen,
    type_breakdown: typeBreakdown,
  } = stats;

  useEffect(() => {
    fetchStatistics(setStats, setLoading);
  }, [setStats, setLoading]);

  if (loading) return <Spin />;

  return (
    <Spin spinning={loading}>
      <div className="statistics">
        <div className="stats">
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
          <Empty />
          <div className="statistics--trendline">
            <h3>Last 10 Days</h3>
            <LineChart width={250} height={65} data={makeGraphData(lastTen)}>
              <Line
                type="monotone"
                dataKey="datum"
                stroke="#0e5a6d"
                strokeWidth={2}
                dot={false}
              />
            </LineChart>
            <Link to="/graph">More Graph </Link>.
          </div>

          <div className="statistics--pie">
            <PieChart width={400} height={200}>
              <Pie
                data={typeBreakdown}
                innerRadius={40}
                outerRadius={60}
                fill="#8884d8"
                paddingAngle={1}
                dataKey="portion"
                nameKey="name"
                label={renderLabel}
              >
                {typeBreakdown.map((entry, index) => (
                  <Cell key={entry.name} fill={COLORS[index % COLORS.length]} />
                ))}
              </Pie>
            </PieChart>
          </div>
        </div>
      </div>
    </Spin>
  );
}
