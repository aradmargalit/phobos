import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { Form, InputNumber, Modal, Spin } from 'antd';
import React, { useState } from 'react';

const unitSizeMap = {
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
const getMaxBy = (metricName, unit) => {
  // If it's a percentage, it's always capped at 100%, easy!
  if (metricName.includes('%')) {
    return 100;
  }

  if (metricName.toLowerCase() === 'miles') {
    return unitSizeMap[unit] * 200; // Assume nobody is running or biking >= 200 miles per day
  }

  // Hours
  return unitSizeMap[unit] * 24; // 24 hours per day
};

const { Item } = Form;

// TODO convert to form, I want validation + reset + onSubmit func
export default function GoalsModal({ visible, onCancel, unit, metricName }) {
  const [form] = Form.useForm();
  const [loading, setLoading] = useState(false);

  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  const onFinish = values => {
    setLoading(true);
    const payload = {
      unit,
      metricName,
      goal: values.goal,
    };

    // eslint-disable-next-line no-console
    console.log(payload);

    // Simulate network call
    setTimeout(() => {
      setLoading(false);
      handleCancel();
    }, 5000);
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
          `Set a ${unit}ly ${metricName} goal...`,
          <AimOutlined style={{ marginRight: '10px' }} />
        )}
      </h3>
      <h5>{`This value will persist across ${unit.toLowerCase()}s until you clear it.`}</h5>
      <Spin spinning={loading}>
        <Form form={form} onFinish={onFinish}>
          <Item
            label={`New Goal (${metricName})`}
            name="goal"
            rules={[{ required: true, message: 'You must save a goal or press "Cancel"' }]}
          >
            <InputNumber placeholder={2} min={1} max={getMaxBy(metricName, unit)} />
          </Item>
        </Form>
      </Spin>
    </Modal>
  );
}
