import './QuickAdd.scss';

import { DeleteOutlined } from '@ant-design/icons';
import { Button, Empty, Spin } from 'antd';
import React, { useEffect, useState } from 'react';

import { fetchQuickAdds } from '../../apis/phobos-api';

export default function QuickAdd() {
  const [quickAdds, setQuickAdds] = useState(null);
  const [loading, setLoading] = useState(true);
  useEffect(() => {
    fetchQuickAdds(setQuickAdds, setLoading);
  }, [setQuickAdds, setLoading]);

  if (loading) return <Spin />;
  if (!quickAdds || !quickAdds.length)
    return <Empty description="Save a workout to quickly add it later!" />;

  return (
    <div className="quick-add">
      {quickAdds.map(qa => (
        <div className="quick-add--button" key={qa.id}>
          <Button>{qa.name}</Button>
          <Button ghost type="danger">
            <DeleteOutlined />
          </Button>
        </div>
      ))}
    </div>
  );
}
