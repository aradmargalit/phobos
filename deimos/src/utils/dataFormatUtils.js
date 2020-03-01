/* eslint-disable import/prefer-default-export */
import moment from 'moment';

export const minutesToTime = (minutes) => {
  if (!minutes || minutes === 0) {
    return '-';
  }
  return moment().startOf('day').add(minutes, 'minutes').format('m:ss');
};
