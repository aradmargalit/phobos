import { Statistic } from 'antd';
import React from 'react';

import { minutesToTime } from '../../utils/dataFormatUtils';

const singularize = (input) => input.slice(0, -1);

export default function CalculatedActivityFields({ activity }) {
  const { duration, distance, unit } = activity;

  return (
    <Statistic title={`min / ${singularize(unit)}`} value={(duration / distance)} formatter={minutesToTime} />
  );
}
