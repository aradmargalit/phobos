import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { Button, Form, InputNumber, Modal, notification, Spin } from 'antd';
import React, { useEffect, useState } from 'react';

import { deleteGoal, fetchGoals, postGoal, putGoal } from '../../apis/phobos-api';

const { Item } = Form;

// Maps between a period and how many days are theoretically contained within it
const periodSizeMap = {
  Week: 7,
  Month: 31, // Not always true, but good enough,
  Year: 365,
};

// Helper function to make a title with an icon in the beginning of it
const iconTitle = (text, icon) => (
  <span>
    {icon}
    {text}
  </span>
);

// Is this a bad function name? Maybe. Gets the theoretical maximum for a given metric over a period
// Makes sure you don't set goals >24 hours per day
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

export default function GoalsModal({
  visible,
  onCancel,
  period,
  metricName,
  setGoals,
  currentGoal,
}) {
  // Use a form this way so we can programmatically manipulate it
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  // After every render, if the modal will show, make sure the value in field reflects the props
  useEffect(() => {
    if (visible) {
      form.setFieldsValue({ goal: currentGoal ? currentGoal.goal : null });
    }
  }, [visible, currentGoal, form]);

  // On cancel, clear the form and close the modal
  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  // Delete, refetch, and close the modal
  const handleDelete = async () => {
    await deleteGoal(currentGoal.id);
    await fetchGoals(setGoals);
    handleCancel();
  };

  const onFinish = async (values) => {
    setLoading(true);

    // In both an update and create case, set the core payload
    const payload = {
      period: period.toLowerCase(),
      metric: metricName.toLowerCase(),
      goal: values.goal,
    };

    // Start by assuming we'll POST
    let apiFunc = postGoal;

    // If there's a current Goal, this is an update
    if (currentGoal) {
      // If the goal is the same, we don't need to waste the API call
      if (currentGoal.goal === values.goal) {
        setLoading(false);
        handleCancel();
        return;
      }

      // If this is a true update, set the ID and override apiFunc to be a put
      payload.id = currentGoal.id;
      apiFunc = putGoal;
    }

    try {
      await apiFunc(payload);
      await fetchGoals(setGoals);
      onCancel();
    } catch (e) {
      notification.error({ message: 'Failed to update goal' });
    } finally {
      setLoading(false);
    }
  };

  // Always show a cancel and submit button
  let footer = [
    <Button key="cancel" onClick={handleCancel}>
      Cancel
    </Button>,
    <Button key="submit" type="primary" loading={loading} onClick={() => form.submit()}>
      Submit Goal
    </Button>,
  ];

  // Conditionally show a delete goal button by pushing to the front of the line
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
      footer={footer}
      forceRender
      destroyOnClose
    >
      <h3>
        {/* Cutesy title, uses the period and metric to form a sentence */}
        {iconTitle(
          `Set a ${period}ly ${metricName} goal...`,
          <AimOutlined style={{ marginRight: '10px' }} />
        )}
      </h3>
      <h5>
        {`This goal will stay on the graph (even across ${period.toLowerCase()}s), until deleted.`}
      </h5>
      <Spin spinning={loading}>
        <Form form={form} onFinish={onFinish}>
          <Item
            label={`New Goal (${metricName})`}
            name="goal"
            rules={[{ required: true, message: 'You must save a goal or press "Cancel"' }]}
          >
            <InputNumber
              // Don't submit on Enter in order to make sure max is respected.
              // Without this, you can submit goals before the Input can correct itself
              onKeyDown={(e) => (e.keyCode === 13 ? e.preventDefault() : '')}
              min={1}
              max={getMaxBy(metricName, period)}
            />
          </Item>
        </Form>
      </Spin>
    </Modal>
  );
}
