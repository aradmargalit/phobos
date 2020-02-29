import React from 'react';
import { Statistic } from 'antd';
import moment from 'moment';

const formatPace = (minutes) => {
  if (!minutes || minutes === 0) {
    return '-';
  }
  return moment().startOf('day').add(minutes, 'minutes').format('m:ss');
};
const singularize = (input) => input.slice(0, -1);

export default function CalculatedActivityFields({ activity }) {
  const { duration, distance, unit } = activity;

  return (
    <Statistic title={`min / ${singularize(unit)}`} value={(duration / distance)} formatter={formatPace} />
  );
}
