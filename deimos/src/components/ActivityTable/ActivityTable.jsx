import { Empty } from 'antd';
import React from 'react';

export default function ActivityTable({ activities }) {
  return activities ? <h1>Here..</h1> : (
    <Empty description="No activities...yet!" />
  );
}
