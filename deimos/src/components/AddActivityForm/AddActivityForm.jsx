/* eslint-disable react/jsx-props-no-spreading */

import {
  Form,
  Button,
  DatePicker,
  Select,
  InputNumber,
  Slider,
  Spin,
} from 'antd';
import PropTypes from 'prop-types';
import React, { useState, useEffect } from 'react';
import './AddActivityForm.scss';
import { RocketOutlined } from '@ant-design/icons';
import moment from 'moment';
import EmojiOption from '../EmojiOption';
import { BACKEND_URL } from '../../constants';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm({ closeModal }) {
  const [loading, setLoading] = useState(true);
  const [activityTypes, setActivityTypes] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      // Make sure to include the cookie with the request!
      const res = await fetch(`${BACKEND_URL}/metadata/activity_types`, {
        credentials: 'include',
      });

      res.json().then(({ activity_types: respTypes }) => {
        setActivityTypes(respTypes);
        setLoading(false);
      });
    };

    fetchData();
  }, [setLoading]);

  const onFinish = (values) => {
    console.log('Success:', values);
    closeModal();
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

  return loading ? (
    <Spin />
  ) : (
    <Form
      {...layout}
      name="add-activity"
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      {/* ============= DATEPICKER ============= */}
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
        <DatePicker defaultValue={moment(new Date())} className="fullWidth" placeholder="2020-01-01" />
      </Item>

      {/* ============= ACTIVITY SELECT ============= */}
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
          {activityTypes.map(({ name }) => EmojiOption({ value: name.toLowerCase(), title: name }))}
        </Select>
      </Item>

      {/* ============= DURATION ============= */}
      <Item label="Duration">
        <Item
          name="duration"
          rules={[{ required: true, message: 'Duration is required' }]}
          noStyle
        >
          <InputNumber min={1} />
        </Item>
        <span className="ant-form-text"> minutes</span>
      </Item>

      {/* ============= DISTANCE ============= */}
      <Item label="Distance" style={{ marginBottom: 0 }}>
        <Item
          name="distance"
          className="inline-item"
        >
          <InputNumber min={0} placeholder={5} />
        </Item>
        <Item name="distance-units" className="inline-item">
          <Select defaultValue="miles">
            <Option value="miles">miles</Option>
            <Option value="yards">yards</Option>
          </Select>
        </Item>
      </Item>


      {/* ============= DIFFICULTY ============= */}
      <Item name="difficulty" label="Difficulty">
        <Slider
          marks={{
            0: 'Easy',
            50: 'Moderate',
            100: 'Exhausting',
          }}
          step={10}
        />
      </Item>

      {/* ============= SUBMIT ============= */}
      <Item {...tailLayout}>
        <Button htmlType="submit" icon={<RocketOutlined />} type="primary">
          Submit
        </Button>
      </Item>
    </Form>
  );
}

AddActivityForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};
