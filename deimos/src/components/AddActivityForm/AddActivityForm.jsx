import './AddActivityForm.scss';

import { EditOutlined, RocketOutlined } from '@ant-design/icons';
import {
  Button,
  DatePicker,
  Form,
  Input,
  InputNumber,
  message,
  notification,
  Select,
  Spin,
} from 'antd';
import moment from 'moment';
import React, { useContext, useState } from 'react';

import { postActivity, putActivity } from '../../apis/phobos-api';
import UserContext from '../../contexts';
import CalculatedActivityFields from '../CalculatedActivityFields';
import EmojiOption from '../EmojiOption';


const dateFormat = 'MMMM Do, YYYY';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm({
  activityTypes, loading, refetch, initialActivity, modalClose,
}) {
  const { user } = useContext(UserContext);
  const [activity, setActivity] = useState({ ...initialActivity, unit: 'miles' });
  const [form] = Form.useForm();

  const editing = !!initialActivity;

  const upsert = async (values, apiCall) => {
    try {
      await apiCall(values);
      message.success(`Successfully ${editing ? 'updated' : 'created'} activity!`);
      refetch();
    } catch (err) {
      notification.error({
        message: 'Unexpected Error',
        description: `Error: ${err}`,
      });
    }
  };

  const onFinish = async (values) => {
    if (editing) {
      const activityToPut = { ...values };
      activityToPut.id = initialActivity.id;
      await upsert(activityToPut, putActivity);
      modalClose();
    } else {
      await upsert(values, postActivity);
    }
  };

  const onReset = () => {
    form.resetFields();
    setActivity(form.getFieldsValue());
  };

  const onSubmit = () => {
    form.submit();
  };

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
    <Spin spinning={loading}>
      <div className="form-flex">
        <Form
          {...layout}
          form={form}
          name="add-activity"
          onFinish={onFinish}
          onValuesChange={onChange}
          initialValues={{ activity_date: moment(new Date()), unit: 'miles', ...initialActivity }}
        >
          {/* ============= NAME ============= */}
          <Item
            hasFeedback
            label="Activity Name"
            name="name"
          >
            <Input className="fullWidth" placeholder={`${user.given_name}'s Favorite Run`} />
          </Item>
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
            <DatePicker format={dateFormat} className="fullWidth" placeholder="2020-01-01" />
          </Item>

          {/* ============= ACTIVITY SELECT ============= */}
          <Item
            hasFeedback
            label="Activity Type"
            name="activity_type_id"
            rules={[
              {
                required: true,
                message: 'Activity Type is required',
              },
            ]}
          >
            <Select allowClear showSearch optionFilterProp="children">
              {activityTypes.map(
                ({ id, name }) => EmojiOption({ value: id, title: name }),
              )}
            </Select>
          </Item>

          {/* ============= DURATION ============= */}
          <Item
            hasFeedback
            label="Duration (min)"
            name="duration"
            rules={[{ required: true, message: 'Duration is required' }]}
          >
            <InputNumber
              className="fullWidth"
              precision={2}
              placeholder={30}
              min={0.5}
              step={0.5}
            />
          </Item>

          {/* ============= DISTANCE ============= */}
          <Item label="Distance" style={{ marginBottom: 0 }}>
            <Item name="distance" className="inline-item">
              <InputNumber
                precision={2}
                min={0.1}
                step={0.1}
                placeholder={5}
              />
            </Item>
            <Item name="unit" className="inline-item">
              <Select>
                <Option value="miles">miles</Option>
                <Option value="yards">yards</Option>
              </Select>
            </Item>
          </Item>
        </Form>
        <CalculatedActivityFields activity={activity} />
      </div>
      <div className="button-row">
        <Button className="button-row-item" onClick={onSubmit} icon={editing ? <EditOutlined /> : <RocketOutlined rotate={45} />} type="primary">
          {editing ? 'Edit' : 'Submit'}
        </Button>
        <Button className="button-row-item" ghost onClick={onReset} type="primary">
          Reset
        </Button>
      </div>
    </Spin>
  );
}
