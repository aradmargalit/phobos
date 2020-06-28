import './GoalsModal.scss';

import { AimOutlined } from '@ant-design/icons';
import { Form, InputNumber, Modal } from 'antd';
import React from 'react';

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
const getMaxBy = (metric, unit) => {
  // If it's a percentage, it's always capped at 100%, easy!
  if (metric.includes('%')) {
    return 100;
  }

  if (metric.toLowerCase() === 'mile') {
    return unitSizeMap[unit] * 200; // Assume nobody is running or biking >= 200 miles per day
  }

  // Hours
  return unitSizeMap[unit] * 24; // 24 hours per day
};

const { Item } = Form;

// TODO convert to form, I want validation + reset + onSubmit func
export default function GoalsModal({ visible, onCancel, unit, metricName }) {
  const [form] = Form.useForm();

  const handleCancel = () => {
    form.resetFields();
    onCancel();
  };

  // eslint-disable-next-line no-unused-vars
  const onFinish = values => {
    // Do something with values
    handleCancel();
  };

  return (
    <Modal
      className="goals-modal"
      width={500}
      visible={visible}
      onCancel={handleCancel}
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
      <Form form={form} onFinish={onFinish}>
        <Item
          label={`New Goal (${metricName})`}
          name="goal"
          nostyle
          rules={[{ required: true, message: 'You must save a goal or press "Cancel"' }]}
        >
          <InputNumber placeholder={2} min={1} max={getMaxBy(metricName, unit)} />
        </Item>
      </Form>
      <div className="goal-entry" />
    </Modal>
  );
}
