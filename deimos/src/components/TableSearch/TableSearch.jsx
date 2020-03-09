import './TableSearch.scss';

import { Input } from 'antd';
import { uniq as _uniq } from 'lodash';
import React from 'react';

import { formatDate } from '../../utils/dataFormatUtils';

const { Search } = Input;

export default function TableSearch({
  tableActivities,
  setTableActivities,
  setFiltered,
}) {
  const onChangeHandler = e => {
    if (!e.target.value || !e.target.value.length) {
      setFiltered(false);
      setTableActivities(null);
      return;
    }

    const searchTerm = e.target.value.toLowerCase();

    const filteredByName = tableActivities.filter(ta =>
      ta.name.toLowerCase().includes(searchTerm)
    );
    const filteredByType = tableActivities.filter(ta =>
      ta.activity_type.name.toLowerCase().includes(searchTerm)
    );

    const filteredByDate = tableActivities.filter(ta =>
      formatDate(ta.activity_date)
        .toLowerCase()
        .includes(searchTerm)
    );
    setFiltered(true);
    setTableActivities(
      _uniq([...filteredByName, ...filteredByType, ...filteredByDate])
    );
  };
  return (
    <Search
      className="search-bar"
      allowClear
      placeholder="Search by name, type, or date..."
      onSearch={value => console.log(value)}
      onChange={onChangeHandler}
      width="50%"
    />
  );
}
