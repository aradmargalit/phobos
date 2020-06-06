import './EditableName.scss';

import { Button, Input } from 'antd';
import moment from 'moment';
import React, { useState } from 'react';

import { putActivity } from '../../apis/phobos-api';

export default function EditableName({ name, record, refetch }) {
  const [editing, setEditing] = useState(false);
  const [inputValue, setInputValue] = useState(name);

  const handleInput = e => setInputValue(e.target.value);
  const handleBlur = () => {
    setEditing(false);
    setInputValue(name);
  };

  const handleSubmit = async () => {
    if (inputValue === name) {
      setEditing(false);
      return;
    }
    const activity = {
      ...record,
      name: inputValue,
      activity_date: moment(record.activity_date),
    };
    await putActivity(activity);
    await refetch();
    setEditing(false);
  };

  return (
    <div>
      {!editing ? (
        <Button type="link" className="editable-name-button" onClick={() => setEditing(true)}>
          {name}
        </Button>
      ) : (
        <>
          <Input
            autoFocus
            onPressEnter={handleSubmit}
            onFocus={e => e.target.select()}
            onBlur={handleBlur}
            size="small"
            allowClear
            onChange={handleInput}
            value={inputValue}
          />
          <p className="editable-name-help">Press Enter to save.</p>
        </>
      )}
    </div>
  );
}
