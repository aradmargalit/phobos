import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { Form, InputNumber, Modal, notification, Spin } from 'antd';
import React, { useState } from 'react';

import { postGoal } from '../../apis/phobos-api';

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

// TODO convert to form, I want validation + reset + onSubmit func
export default function GoalsModal({ visible, onCancel, period, metricName }) {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  const onFinish = async values => {
    setLoading(true);
    const payload = {
      period: period.toLowerCase(),
      metric: metricName.toLowerCase(),
      goal: values.goal,
    };

    try {
      await postGoal(payload);
      onCancel();
    } catch (e) {
      notification.error({ message: e });
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal
      className="goals-modal"
      width={700}
      visible={visible}
      onCancel={handleCancel}
      okText="Submit Goal"
      onOk={() => form.submit()}
      destroyOnClose
    >
      <h3>
        {iconTitle(
          `Set a ${period}ly ${metricName} goal...`,
          <AimOutlined style={{ marginRight: '10px' }} />
        )}
      </h3>
      <h5>{`This value will persist across ${period.toLowerCase()}s until you clear it.`}</h5>
      <Spin spinning={loading}>
        <Form form={form} onFinish={onFinish}>
          <Item
            label={`New Goal (${metricName})`}
            name="goal"
            rules={[{ required: true, message: 'You must save a goal or press "Cancel"' }]}
          >
            <InputNumber placeholder={2} min={1} max={getMaxBy(metricName, period)} />
          </Item>
        </Form>
      </Spin>
    </Modal>
  );
}
