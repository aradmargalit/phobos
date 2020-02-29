/* eslint-disable react/jsx-props-no-spreading */
import {
  Form,
  Button,
  DatePicker,
  Select,
  InputNumber,
  Slider,
  Spin,
  notification,
  message,
} from 'antd';
import PropTypes from 'prop-types';
import React, { useState, useEffect } from 'react';
import './AddActivityForm.scss';
import { RocketOutlined } from '@ant-design/icons';
import moment from 'moment';
import EmojiOption from '../EmojiOption';
import { fetchActivityTypes, postActivity } from '../../apis/phobos-api';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm({ closeModal }) {
  const [loading, setLoading] = useState(true);
  const [activityTypes, setActivityTypes] = useState([]);

  useEffect(() => {
    fetchActivityTypes(setActivityTypes, setLoading);
  }, [setLoading]);

  const onFinish = (values) => {
    setLoading(true);
    postActivity(values)
      .then((data) => {
        message.success(`Successfully created activity: ${data}`);
        closeModal();
      })
      .catch((err) => {
        notification.error({
          message: 'Unexpected Error',
          description: `Error: ${err}`,
        });
      })
      .finally(() => {
        setLoading(false);
      });
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
    <Spin spinning={loading}>
      <Form
        {...layout}
        name="add-activity"
        onFinish={onFinish}
        initialValues={{ activity_date: moment(new Date()) }}
      >
        {/* ============= DATEPICKER ============= */}
        <Item
          hasFeedback
          label="Activity Date"
          name="activity_date"
          rules={[
            {
              required: true,
              message: 'Activity Date is required',
            },
          ]}
        >
          <DatePicker className="fullWidth" placeholder="2020-01-01" />
        </Item>

        {/* ============= ACTIVITY SELECT ============= */}
        <Item
          hasFeedback
          label="Activity Type"
          name="activity_type"
          rules={[
            {
              required: true,
              message: 'Activity Type is required',
            },
          ]}
        >
          <Select allowClear showSearch optionFilterProp="children">
            {activityTypes.map(
              ({ name }) => EmojiOption({ value: name.toLowerCase(), title: name }),
            )}
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
          <Item name="distance" className="inline-item">
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
    </Spin>
  );
}

AddActivityForm.propTypes = {
  closeModal: PropTypes.func.isRequired,
};
