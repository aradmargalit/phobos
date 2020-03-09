import './DurationPicker.scss';

import { InputNumber } from 'antd';
import React from 'react';

const hmsToMinutes = ({ hours, minutes, seconds }) =>
  hours * 60 + minutes + seconds / 60;

// Value is an object of {hours, minutes, seconds, total}
export default function DurationPicker({ value, onChange }) {
  const { hours, minutes, seconds } = value;

  const onHoursChange = newHours => {
    const newValue = { hours: newHours, minutes, seconds };
    onChange({ ...newValue, total: hmsToMinutes(newValue) });
  };

  const onMinutesChange = newMinutes => {
    if (newMinutes > 59) return;
    const newValue = { hours, minutes: newMinutes, seconds };
    onChange({ ...newValue, total: hmsToMinutes(newValue) });
  };

  const onSecondsChange = newSeconds => {
    if (newSeconds > 59) return;
    const newValue = { hours, minutes, seconds: newSeconds };
    onChange({ ...newValue, total: hmsToMinutes(newValue) });
  };

  return (
    <div className="duration-input-container">
      <InputNumber
        placeholder="hours"
        value={hours}
        onChange={onHoursChange}
        min={0}
      />
      :
      <InputNumber
        placeholder="minutes"
        value={minutes}
        onChange={onMinutesChange}
        min={0}
        max={59}
      />
      :
      <InputNumber
        placeholder="seconds "
        value={seconds}
        onChange={onSecondsChange}
        min={0}
        max={59}
      />
    </div>
  );
}
