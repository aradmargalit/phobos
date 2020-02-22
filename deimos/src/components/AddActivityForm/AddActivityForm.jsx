/* eslint-disable react/jsx-props-no-spreading */

import {
  Form, Button, DatePicker, Select, InputNumber, Slider,
} from 'antd';
import React from 'react';
import './AddActivityForm.scss';
import { RocketOutlined } from '@ant-design/icons';
import EmojiOption from '../EmojiOption';

const { Item } = Form;

export default function AddActivityForm() {
  const onFinish = (values) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };

  const layout = {
    labelCol: {
      span: 6,
    },
    wrapperCol: {
      span: 12,
    },
  };

  const tailLayout = {
    wrapperCol: {
      offset: 18,
      span: 6,
    },
  };

  return (
    <Form
      {...layout}
      name="add-activity"
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <Item
        hasFeedback
        label="Activity Date"
        name="activity-date"
        rules={[
          {
            required: true,
            message: 'Activity Date is required',
          },
        ]}
      >
        <DatePicker className="fullWidth" placeholder="2020-01-01" />
      </Item>

      <Item
        hasFeedback
        label="Activity Type"
        name="activity-type"
        rules={[
          {
            required: true,
            message: 'Activity Type is required',
          },
        ]}
      >
        <Select allowClear showSearch optionFilterProp="children">
          {EmojiOption({ emoji: '🏃', value: 'run', title: 'Run' })}
          {EmojiOption({ emoji: '🚴', value: 'bike', title: 'Bike' })}
          {EmojiOption({ emoji: '🏊', value: 'swim', title: 'Swimming' })}
        </Select>
      </Item>

      <Item
        hasFeedback
        label="Duration (min)"
        name="duration"
        rules={[
          {
            required: true,
            message: 'Duration is required',
            type: 'number',
            min: 1,
          },
        ]}
      >
        <InputNumber className="fullWidth" min={1} placeholder={30} />
      </Item>

      <Item name="difficulty" label="Difficulty">
        <Slider
          marks={{
            0: 'Easy',
            50: 'Moderate',
            100: 'Exhausting',
          }}
        />
      </Item>

      <Item {...tailLayout}>
        <Button htmlType="submit" icon={<RocketOutlined />} type="primary">
          Submit
        </Button>
      </Item>
    </Form>
  );
}
