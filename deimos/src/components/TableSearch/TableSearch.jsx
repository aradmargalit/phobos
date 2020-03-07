import { Input } from 'antd';
import React from 'react';

const { Search } = Input;


export default function TableSearch({
  tableActivities, setTableActivities, setFiltered, activityTypes,
}) {
  const onChangeHandler = (e) => {
    if (!e.target.value || !e.target.value.length) {
      setFiltered(false);
      setTableActivities(null);
      return;
    }

    const searchTerm = e.target.value.toLowerCase();
    const filteredByName = tableActivities.filter(
      (ta) => ta.name.toLowerCase().includes(searchTerm),
    );
    const filteredByType = tableActivities.filter(
      (ta) => activityTypes.find((at) => at.id === ta.activity_type_id)
        .name.toLowerCase().includes(searchTerm),
    );
    setFiltered(true);
    setTableActivities([...filteredByName, ...filteredByType]);
  };
  return (
    <Search
      allowClear
      placeholder="Search by name, type, distance, or date..."
      onSearch={(value) => console.log(value)}
      onChange={onChangeHandler}
      width="50%"
    />
  );
}
