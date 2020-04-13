export const numberSorter = (a, b) => a.logical_index - b.logical_index;

export const nameSorter = (a, b) => a.name.localeCompare(b.name);

export const dateSorter = (a, b) => a.epoch - b.epoch;

export const activityTypeSorter = (a, b) =>
  a.activity_type.name.localeCompare(b.activity_type.name);

export const durationSorter = (a, b) => a.duration - b.duration;

export const distanceSorter = (a, b) => a.distance - b.distance;
