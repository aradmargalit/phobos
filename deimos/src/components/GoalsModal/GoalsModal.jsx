import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { Button, Form, InputNumber, Modal, notification, Spin } from 'antd';
import React, { useState } from 'react';

import { deleteGoal, fetchGoals, postGoal, putGoal } from '../../apis/phobos-api';

const periodSizeMap = {
  Week: 7,
  Month: 31, // Not always true, but good enough,
  Year: 365,
};

const iconTitle = (text, icon) => (
  <span>
    {icon}
    {text}
  </span>
);

// Is this a bad function name? Maybe.
const getMaxBy = (metricName, period) => {
  // If it's a percentage, it's always capped at 100%, easy!
  if (metricName.includes('%')) {
    return 100;
  }

  if (metricName.toLowerCase() === 'miles') {
    return periodSizeMap[period] * 200; // Assume nobody is running or biking >= 200 miles per day
  }

  // Hours
  return periodSizeMap[period] * 24; // 24 hours per day
};

const { Item } = Form;

export default function GoalsModal({
  visible,
  onCancel,
  period,
  metricName,
  setGoals,
  currentGoal,
}) {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const initialValues = currentGoal ? { goal: currentGoal.goal } : null;

  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  const handleDelete = async () => {
    await deleteGoal(currentGoal.id);
    await fetchGoals(setGoals);
    handleCancel();
  };

  const onFinish = async values => {
    setLoading(true);
    const payload = {
      period: period.toLowerCase(),
      metric: metricName.toLowerCase(),
      goal: values.goal,
    };

    let apiFunc = postGoal;
    if (currentGoal) {
      if (currentGoal.goal === values.goal) {
        setLoading(false);
        return;
      }

      payload.id = currentGoal.id;
      apiFunc = putGoal;
    }

    try {
      await apiFunc(payload);
      await fetchGoals(setGoals);
      onCancel();
    } catch (e) {
      notification.error({ message: 'Failed to update activity' });
    } finally {
      setLoading(false);
    }
  };

  let footer = [
    <Button key="cancel" onClick={handleCancel}>
      Cancel
    </Button>,
    <Button key="submit" type="primary" loading={loading} onClick={() => form.submit()}>
      Submit Goal
    </Button>,
  ];

  if (currentGoal) {
    footer = [
      <Button key="clear" type="danger" loading={loading} onClick={handleDelete}>
        Delete Goal
      </Button>,
      ...footer,
    ];
  }

  return (
    <Modal
      className="goals-modal"
      width={700}
      visible={visible}
      onCancel={handleCancel}
      okText="Submit Goal"
      destroyOnClose
      footer={footer}
    >
      <h3>
        {iconTitle(
          `Set a ${period}ly ${metricName} goal...`,
          <AimOutlined style={{ marginRight: '10px' }} />
        )}
      </h3>
      <h5>
        {`This goal will stay on the graph (even across ${period.toLowerCase()}s), until deleted.`}
      </h5>
      <Spin spinning={loading}>
        <Form form={form} initialValues={initialValues} onFinish={onFinish}>
          <Item
            label={`New Goal (${metricName})`}
            name="goal"
            rules={[{ required: true, message: 'You must save a goal or press "Cancel"' }]}
          >
            <InputNumber
              onKeyDown={e => (e.keyCode === 13 ? e.preventDefault() : '')}
              placeholder={2}
              min={1}
              max={getMaxBy(metricName, period)}
            />
          </Item>
        </Form>
      </Spin>
    </Modal>
  );
}
