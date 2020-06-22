import './EditableName.scss';

import { Button, Input, Tooltip } from 'antd';
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

  const ElipsesButton = ({ text, maxChar, onClickHandler }) => {
    const shouldTruncate = text.length > maxChar;
    // Determine if we need to truncate or not, based on maxChar
    const displayText = shouldTruncate ? `${text.substring(0, maxChar)}...` : text;

    const displayButton = (
      <Button type="link" className="editable-name-button" onClick={onClickHandler}>
        {displayText}
      </Button>
    );

    return shouldTruncate ? (
      <Tooltip placement="topLeft" title={text} mouseEnterDelay={0} mouseLeaveDelay={0}>
        {displayButton}
      </Tooltip>
    ) : (
      displayButton
    );
  };

  return (
    <div>
      {!editing ? (
        <ElipsesButton text={name} maxChar={20} onClickHandler={() => setEditing(true)} />
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
