/* eslint-disable import/prefer-default-export */
export const totalToHMS = duration => ({
  hours: Math.floor(duration / 60),
  minutes: Math.floor(duration % 60),
  seconds: Math.floor(duration - Math.floor(duration) * 60),
  total: duration,
});
