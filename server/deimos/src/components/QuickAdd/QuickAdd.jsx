import './QuickAdd.scss';

import { Button, Empty } from 'antd';
import React from 'react';

export default function QuickAdd() {
  return (
    <div className="quick-add">
      <Empty description="Save a workout to quickly add it later!" />
      <span>
        <Button>Sample Workout</Button>
        Trash Button TBD
      </span>
    </div>
  );
}
