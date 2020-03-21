import { EditOutlined, RocketOutlined } from '@ant-design/icons';
import { Button } from 'antd';
import React from 'react';

export default function FormButtons({
  editing,
  loading,
  onSubmit,
  onReset,
  onSaveQuickAdd,
}) {
  return (
    <div className="button-row">
      <Button
        className="button-row-item"
        onClick={onSubmit}
        icon={editing ? <EditOutlined /> : <RocketOutlined rotate={45} />}
        type="primary"
        loading={loading}
        disabled={loading}
      >
        {editing ? 'Edit' : 'Submit'}
      </Button>
      <Button
        className="button-row-item"
        ghost
        onClick={onReset}
        type="primary"
      >
        Reset
      </Button>
      {/* Only show QuickAdd if the function was passed in as a prop */}
      {onSaveQuickAdd && (
        <Button
          className="button-row-item"
          onClick={onSaveQuickAdd}
          type="dashed"
        >
          Save for Quick-Add
        </Button>
      )}
    </div>
  );
}
