import './QuickAdd.scss';

import { DeleteOutlined } from '@ant-design/icons';
import { Button, Empty, Spin } from 'antd';
import React from 'react';
import { CSSTransition, TransitionGroup } from 'react-transition-group';

import { deleteQuickAdd, fetchQuickAdds } from '../../apis/phobos-api';

export default function QuickAdd({ quickAdds, setQuickAdds, setQuickAdd }) {
  const onDelete = async (id) => {
    await deleteQuickAdd(id);
    await fetchQuickAdds(setQuickAdds);
  };

  if (quickAdds.loading) return <Spin />;
  if (!quickAdds.payload || !quickAdds.payload.length)
    return <Empty description="Save a workout to quickly add it later!" />;

  return (
    <TransitionGroup className="quick-add">
      {quickAdds.payload.map((qa) => (
        <CSSTransition key={qa.id} timeout={250} classNames="move">
          <div className="quick-add--list">
            <Button className="quick-add--button" onClick={() => setQuickAdd(qa)}>
              {qa.name}
            </Button>
            <Button ghost type="danger" onClick={() => onDelete(qa.id)}>
              <DeleteOutlined />
            </Button>
          </div>
        </CSSTransition>
      ))}
    </TransitionGroup>
  );
}
