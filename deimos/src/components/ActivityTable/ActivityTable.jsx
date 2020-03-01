import './ActivityTable.scss';

import { Empty, Spin, Table } from 'antd';
import { snakeCase as _snakeCase } from 'lodash';
import moment from 'moment';
import React from 'react';


const toCol = (name, render) => {
  const snakeName = _snakeCase(name);
  const col = { title: name, dataIndex: snakeName };
  if (render) {
    col.render = render;
  }
  return col;
};

const dateFormat = 'MMMM Do, YYYY';

const formatDate = (date) => {
  const localDate = new Date(`${date} UTC`);
  return moment(localDate).format(dateFormat);
};

export default function ActivityTable({ loading, activityTypes, activities }) {
  if (loading || !activityTypes.length) return <Spin />;
  if (!activities) return <Empty description="No activities...yet!" />;

  const columns = [
    {
      title: 'No.',
      dataIndex: 'id',
    },
    toCol('Name'),
    toCol('Activity Date', formatDate),
    {
      title: 'Activity Type',
      dataIndex: 'activity_type_id',
      render: (id) => activityTypes.find((at) => at.id === id).name,
    },
    toCol('Duration', (duration) => `${duration} min`),
    toCol('Distance', (distance, record) => `${distance} ${record.unit}`),
    toCol('Created At', formatDate),
  ];
  return <Table pagination={false} rowKey="id" dataSource={activities} columns={columns} />;
}
