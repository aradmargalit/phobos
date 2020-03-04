import './DurationPicker.scss';

import { InputNumber } from 'antd';
import React from 'react';

const calculateHMS = (rawMinutes) => {
  // If we don't have anything yet, that's okay! Just fill with nulls
  if (!rawMinutes) return [null, null, null];

  // 1.Figure out the minutes, without seconds
  const roundMinutes = Math.floor(rawMinutes);

  // 2. Set seconds (this part is easy-ish)
  const seconds = Math.ceil((rawMinutes - roundMinutes) * 60);

  // 3. Pull hours out of roundMinutes
  const hours = Math.floor(roundMinutes / 60);

  // 4. Refined minutes will be the modulo from the hours
  const minutes = roundMinutes % 60;

  return [hours, minutes, seconds];
};

const hmsToMinutes = (hours, minutes, seconds) => (hours * 60 + minutes + seconds / 60);
const leadingZero = (value) => ((value && value < 10) ? `0${value}`.replace(/0*/, '0') : value);

export default function DurationPicker({ value, onChange }) {
  const [hours, minutes, seconds] = calculateHMS(value);

  const onHoursChange = (newHours) => {
    onChange(hmsToMinutes(newHours, minutes, seconds));
  };

  const onMinutesChange = (newMinutes) => {
    onChange(hmsToMinutes(hours, newMinutes, seconds));
  };

  const onSecondsChange = (newSeconds) => {
    onChange(hmsToMinutes(hours, minutes, newSeconds));
  };

  return (
    <div className="duration-input-container">
      <InputNumber
        formatter={leadingZero}
        placeholder="hours"
        value={hours}
        onChange={onHoursChange}
        min={0}
      />
      :
      <InputNumber
        formatter={leadingZero}
        placeholder="minutes"
        value={minutes}
        onChange={onMinutesChange}
        min={0}
        max={59}
      />
      :
      <InputNumber
        formatter={leadingZero}
        placeholder="seconds "
        value={seconds}
        onChange={onSecondsChange}
        min={0}
        max={59}
      />
    </div>
  );
}
