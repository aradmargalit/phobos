import { Empty, Spin, Table } from 'antd';
import { snakeCase as _snakeCase } from 'lodash';
import moment from 'moment';
import React, { useEffect, useState } from 'react';

import { fetchActivities } from '../../apis/phobos-api';

const toCol = (name, render) => {
  const snakeName = _snakeCase(name);
  const col = { title: name, dataIndex: snakeName };
  if (render) {
    col.render = render;
  }
  return col;
};

const dateFormat = 'MMMM Do, YYYY';

const columns = [
  toCol('Name'),
  toCol('Activity Date', (date) => moment(date).format(dateFormat)),
  toCol('Activity Type'),
  toCol('Duration', (duration) => <p>{`${duration} min`}</p>),
  toCol('Distance', (distance, record) => <p>{`${distance} ${record.unit}`}</p>),
  toCol('Created At', (date) => moment(date).format(dateFormat)),
];

export default function ActivityTable() {
  const [activities, setActivities] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchActivities(setActivities, setLoading);
  }, [setLoading]);

  if (loading) return <Spin />;
  if (!activities) return <Empty description="No activities...yet!" />;

  return <Table rowKey="id" dataSource={activities} columns={columns} />;
}
