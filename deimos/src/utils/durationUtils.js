/* eslint-disable import/prefer-default-export */
export const makeDurationBreakdown = min => ({
  hours: Math.floor(min / 60),
  minutes: Math.floor(min % 60),
  seconds: Math.floor((min - Math.floor(min)) * 60),
  total: min,
});
