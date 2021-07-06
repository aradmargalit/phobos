import './DurationPicker.scss';

import { InputNumber } from 'antd';
import React from 'react';

const hmsToMinutes = ({ hours, minutes, seconds }) => hours * 60 + minutes + seconds / 60;

// Value is an object of {hours, minutes, seconds, total}
export default function DurationPicker({ value, onChange }) {
  // Default to nulls in the event that value is null or empty
  const hms = { hours: null, minutes: null, seconds: null, ...value };
  const { hours, minutes, seconds } = hms;

  // Common utility to calculate the new state and call the form's onChange event
  const calcAndUpdate = (newValue) => onChange({ ...newValue, total: hmsToMinutes(newValue) });

  // When one of the three inputs change, exit if they exceed the max, or update with the new value
  const onInputChange = (key, _value, max) => {
    if (max && value > max) return;
    calcAndUpdate({ hours, minutes, seconds, [key]: _value });
  };

  return (
    <div className="duration-input-container">
      <InputNumber
        placeholder="hours"
        value={hours}
        onChange={(_value) => onInputChange('hours', _value)}
        min={0}
      />
      :
      <InputNumber
        placeholder="minutes"
        value={minutes}
        onChange={(_value) => onInputChange('minutes', _value, 59)}
        min={0}
        max={59}
      />
      :
      <InputNumber
        placeholder="seconds"
        value={seconds}
        onChange={(_value) => onInputChange('seconds', _value, 59)}
        min={0}
        max={59}
      />
    </div>
  );
}
