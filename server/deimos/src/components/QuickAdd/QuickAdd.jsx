import './QuickAdd.scss';

import { DeleteOutlined } from '@ant-design/icons';
import { Button, Empty, Spin } from 'antd';
import React, { useEffect } from 'react';
import { CSSTransition, TransitionGroup } from 'react-transition-group';

import { deleteQuickAdd, fetchQuickAdds } from '../../apis/phobos-api';

export default function QuickAdd({
  quickAdds,
  setQuickAdds,
  loading,
  setLoading,
  setQuickAdd,
}) {
  useEffect(() => {
    fetchQuickAdds(setQuickAdds, setLoading);
  }, [setQuickAdds, setLoading]);

  const onDelete = async id => {
    await deleteQuickAdd(id);
    await fetchQuickAdds(setQuickAdds, setLoading);
  };

  if (loading) return <Spin />;
  if (!quickAdds || !quickAdds.length)
    return <Empty description="Save a workout to quickly add it later!" />;

  return (
    <TransitionGroup className="quick-add">
      {quickAdds.map(qa => (
        <CSSTransition key={qa.id} timeout={250} classNames="move">
          <div className="quick-add--list">
            <Button
              className="quick-add--button"
              onClick={() => setQuickAdd(qa)}
            >
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
