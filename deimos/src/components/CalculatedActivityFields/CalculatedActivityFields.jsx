import { Statistic } from 'antd';
import React from 'react';

import { minutesToHMS } from '../../utils/dataFormatUtils';

const singularize = input => input.slice(0, -1);

export default function CalculatedActivityFields({ activity }) {
  if (!activity || !activity.duration || !activity.unit) {
    return (
      <Statistic
        className="form-statistic"
        title="min / mile"
        value={null}
        formatter={minutesToHMS}
      />
    );
  }

  const {
    duration: { total },
    distance,
    unit,
  } = activity;
  const totalMinutes = total;

  return (
    <Statistic
      className="form-statistic"
      title={`min / ${singularize(unit)}`}
      value={totalMinutes / distance}
      formatter={minutesToHMS}
    />
  );
}
