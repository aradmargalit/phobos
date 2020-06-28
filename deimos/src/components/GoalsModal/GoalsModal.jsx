import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { InputNumber, Modal, notification } from 'antd';
import React, { useState } from 'react';

const iconTitle = (text, icon) => (
  <span>
    {icon}
    {text}
  </span>
);

export default function GoalsModal({ visible, onCancel, unit, metricName }) {
  const [goal, setGoal] = useState(null);

  const handleSubmit = async () => {
    if (!goal || goal < 1) {
      notification.error({
        message: 'Goal must be > 0. Please cancel if you do not wish to set a goal.',
      });
      return;
    }

    // Close the modal
    onCancel();
  };

  return (
    <Modal className="goals-modal" visible={visible} onCancel={onCancel} onOk={handleSubmit}>
      <h3>
        {iconTitle(
          `Set a ${unit}ly ${metricName} goal...`,
          <AimOutlined style={{ marginRight: '10px' }} />
        )}
      </h3>
      <div className="goal-entry">
        <InputNumber value={goal} onChange={v => setGoal(v)} placeholder={2} min={0} />
        <h4>{metricName}</h4>
      </div>
    </Modal>
  );
}
