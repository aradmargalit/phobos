import { Statistic } from 'antd';
import React from 'react';

import { minutesToHMS } from '../../utils/dataFormatUtils';

const singularize = input => input.slice(0, -1);

export default function CalculatedActivityFields({ activity }) {
  const {
    duration: { total },
    distance,
    unit,
  } = activity;
  const totalMinutes = total;

  return (
    <Statistic
      title={`min / ${singularize(unit)}`}
      value={totalMinutes / distance}
      formatter={minutesToHMS}
    />
  );
}
