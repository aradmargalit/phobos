import React from 'react';
import { Statistic } from 'antd';
import moment from 'moment';

const formatPace = (fraction) => moment().startOf('day').add(fraction, 'minutes').format('m:ss');
const singularize = (input) => input.slice(0, -1);

export default function CalculatedActivityFields({ activity }) {
  const { duration, distance, unit } = activity;

  const u = unit || 'miles';

  return (

    <Statistic title={`min / ${singularize(u)}`} value={formatPace(duration / distance)} />
  );
}
