import './AddActivityForm.scss';

import { DatePicker, Form, Input, InputNumber, Select } from 'antd';
import moment from 'moment';
import React from 'react';

import DurationPicker from '../DurationPicker';
import EmojiOption from '../EmojiOption';

const dateFormat = 'MMMM Do, YYYY';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm({
  form,
  user,
  onFinish,
  activityTypes,
  setActivity,
  initialActivity,
}) {
  const onChange = (_, all) => {
    setActivity(all);
  };

  const layout = {
    labelCol: {
      span: 6,
    },
    wrapperCol: {
      span: 12,
    },
  };

  return (
    <div className="form-flex">
      <Form
        {...layout}
        autoComplete="off"
        form={form}
        name="add-activity"
        onFinish={onFinish}
        onValuesChange={onChange}
        initialValues={{
          activity_date: moment(new Date()),
          unit: 'miles',
          ...initialActivity,
        }}
      >
        {/* ============= NAME ============= */}
        <Item label="Activity Name" name="name">
          {/* Data LP ignore is to stop LastPass from trying to fill the form */}
          <Input
            data-lpignore="true"
            className="fullWidth"
            placeholder={user && `${user.given_name}'s Favorite Run`}
          />
        </Item>
        {/* ============= DATEPICKER ============= */}
        <Item
          label="Activity Date"
          name="activity_date"
          rules={[
            {
              required: true,
              message: 'Activity Date is required',
            },
          ]}
        >
          <DatePicker
            format={dateFormat}
            className="fullWidth"
            placeholder="2020-01-01"
          />
        </Item>
        {/* ============= ACTIVITY SELECT ============= */}
        <Item
          label="Activity Type"
          name="activity_type_id"
          rules={[
            {
              required: true,
              message: 'Activity Type is required',
            },
          ]}
        >
          <Select
            allowClear
            showSearch
            placeholder="Run"
            optionFilterProp="children"
          >
            {activityTypes.map(({ id, name }) =>
              EmojiOption({ value: id, title: name })
            )}
          </Select>
        </Item>
        {/* ============= DURATION ============= */}
        <Item
          label="Duration"
          name="duration"
          rules={[{ required: true, message: 'Duration is required' }]}
        >
          <DurationPicker />
        </Item>
        {/* ============= DISTANCE ============= */}
        <Item label="Distance" className="distance-wrapper">
          <Item name="distance" className="inline-item">
            <InputNumber precision={2} min={0.1} step={0.1} placeholder={5} />
          </Item>
          <Item name="unit" className="unit inline-item">
            <Select>
              <Option value="miles">miles</Option>
              <Option value="yards">yards</Option>
            </Select>
          </Item>
        </Item>
      </Form>
    </div>
  );
}
