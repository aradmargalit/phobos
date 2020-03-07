import moment from 'moment';

const dateFormat = 'MMMM Do, YYYY';

export const minutesToHMS = (minutes) => {
  if (!minutes || minutes === 0) {
    return '-';
  }
  return moment().startOf('day').add(minutes, 'minutes').format('HH:mm:ss');
};

export const formatDate = (date) => moment(date).format(dateFormat);

export const toLocalDate = (date) => {
  const localDate = new Date(`${date} UTC`);
  return moment(localDate).format(dateFormat);
};
