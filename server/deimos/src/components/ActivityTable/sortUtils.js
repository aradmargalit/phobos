import moment from 'moment';

export const numberSorter = (a, b) => {
  return a.logical_index - b.logical_index;
};

export const nameSorter = (a, b) => {
  return a.name.localeCompare(b.name);
};

export const dateSorter = (a, b) => {
  return moment(a.activity_date).diff(moment(b.activity_date));
};

export const activityTypeSorter = (a, b) => {
  return a.activity_type.name.localeCompare(b.activity_type.name);
};

export const durationSorter = (a, b) => {
  return a.duration - b.duration;
};

export const distanceSorter = (a, b) => {
  return a.distance - b.distance;
};
