/* eslint-disable import/prefer-default-export */
export const makeDurationBreakdown = duration => ({
  hours: Math.floor(duration / 60),
  minutes: Math.floor(duration % 60),
  seconds: Math.floor(duration - Math.floor(duration) * 60),
  total: duration,
});
