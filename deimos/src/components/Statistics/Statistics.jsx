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
        </div>
        <div className="statistics--trendline">
          <h3>Last 10 Days</h3>
          <ResponsiveContainer id="trendline-wrapper" width="75%">
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
          <Link to="/graph">More Graph </Link>.
        </div>
      </div>
    </Spin>
  );
}

// TODO: Find a home for these
// const COLORS = ['#bd4946', '#d4524e', '#ec5b57', '#f07c79', '#f49d9a'];

// const renderLabel = d => d.name;

/* <div className="statistics--dow">
            <h3>Daily Breakdown</h3>
            <BarChart
              width={250}
              height={125}
              data={dayBreakdown}
              margin={{
                top: 5,
                right: 30,
                left: 20,
                bottom: 5,
              }}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="day_of_week" />
              <YAxis />
              <Tooltip />
              <Bar dataKey="count" fill="#0e5a6d">
                {dayBreakdown.map((entry, index) => (
                  <Cell
                    key={`cell-${index}`}
                    fill={COLORS[index % COLORS.length]}
                  />
                ))}
              </Bar>
            </BarChart>
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
                </div> */
