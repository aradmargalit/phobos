import {
  Form,
  Button,
  DatePicker,
  Select,
  Input,
  InputNumber,
  Spin,
  notification,
  message,
} from 'antd';
import React, { useState, useEffect } from 'react';
import './AddActivityForm.scss';
import { RocketOutlined } from '@ant-design/icons';
import moment from 'moment';
import EmojiOption from '../EmojiOption';
import CalculatedActivityFields from '../CalculatedActivityFields';
import { fetchActivityTypes, postActivity } from '../../apis/phobos-api';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm() {
  const [loading, setLoading] = useState(true);
  const [activityTypes, setActivityTypes] = useState([]);
  const [activity, setActivity] = useState({ unit: 'miles' });
  const [form] = Form.useForm();

  useEffect(() => {
    fetchActivityTypes(setActivityTypes, setLoading);
  }, [setLoading]);

  const onFinish = (values) => {
    setLoading(true);
    postActivity(values)
      .then((data) => {
        message.success(`Successfully created activity: ${data}`);
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
          initialValues={{ activity_date: moment(new Date()), unit: 'miles' }}
        >
          {/* ============= NAME ============= */}
          <Item
            hasFeedback
            label="Name"
            name="name"
          >
            <Input className="fullWidth" placeholder="Activity Name" />
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
          <Item
            hasFeedback
            label="Duration (min)"
            name="duration"
            rules={[{ required: true, message: 'Duration is required' }]}
          >
            <InputNumber className="fullWidth" min={0.1} />
          </Item>

          {/* ============= DISTANCE ============= */}
          <Item label="Distance" style={{ marginBottom: 0 }}>
            <Item name="distance" className="inline-item">
              <InputNumber min={0.1} placeholder={5} />
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
        <Button onClick={onSubmit} icon={<RocketOutlined rotate={45} />} type="primary">
          Submit
        </Button>
        <Button ghost onClick={onReset} type="primary">
          Reset
        </Button>
      </div>
    </Spin>
  );
}
