import { snakeCase as _snakeCase, uniq as _uniq } from 'lodash';

import { formatDate } from '../../utils/dataFormatUtils';

export const toCol = (name, render) => {
  const snakeName = _snakeCase(name);
  const col = { title: name, dataIndex: snakeName, ellipsis: true };
  if (render) {
    col.render = render;
  }
  return col;
};

export const filterActivities = (term, activities) => {
  if (!term || term.length === 0) return activities;

  const filteredByName = activities.filter(ta =>
    ta.name.toLowerCase().includes(term)
  );

  const filteredByType = activities.filter(ta =>
    ta.activity_type.name.toLowerCase().includes(term)
  );

  const filteredByDate = activities.filter(ta =>
    formatDate(ta.activity_date)
      .toLowerCase()
      .includes(term)
  );

  return _uniq([...filteredByName, ...filteredByType, ...filteredByDate]);
};
