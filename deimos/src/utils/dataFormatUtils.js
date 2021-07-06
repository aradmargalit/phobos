import moment from 'moment';

const dateFormat = 'MMMM Do, YYYY';
const minutesInDay = 24 * 60;

export const minutesToHMS = (minutes) => {
  if (!minutes || minutes === 0 || !Number.isFinite(minutes)) {
    return '-';
  }

  const hms = moment().startOf('day').add(minutes, 'minutes').format('HH:mm:ss');

  return `${minutes >= minutesInDay ? `${Math.floor(minutes / minutesInDay)}d, ` : ''}${hms}`;
};

export const formatDate = (date) => moment(date).format(dateFormat);
