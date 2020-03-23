import moment from 'moment';

const dateFormat = 'MMMM Do, YYYY';

export const minutesToHMS = minutes => {
  if (!minutes || minutes === 0 || !Number.isFinite(minutes)) {
    return '-';
  }
  return moment()
    .startOf('day')
    .add(minutes, 'minutes')
    .format('HH:mm:ss');
};

export const formatDate = date => moment(date).format(dateFormat);
