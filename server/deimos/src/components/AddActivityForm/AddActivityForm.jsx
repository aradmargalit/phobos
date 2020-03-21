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
} from 'antd';
import moment from 'moment';
import React, { useContext, useEffect, useState } from 'react';

import {
  fetchActivityTypes,
  postActivity,
  postQuickAdd,
  putActivity,
} from '../../apis/phobos-api';
import { UserContext } from '../../contexts';
import CalculatedActivityFields from '../CalculatedActivityFields';
import DurationPicker from '../DurationPicker';
import EmojiOption from '../EmojiOption';

const dateFormat = 'MMMM Do, YYYY';

const { Item } = Form;
const { Option } = Select;

export default function AddActivityForm({
  form,
  refetch,
  initialActivity,
  modalClose,
}) {
  const { user } = useContext(UserContext);
  const defaultActivity = {
    unit: 'miles',
    duration: {
      hours: null,
      minutes: null,
      seconds: null,
      total: 0,
    },
  };
  const [activityTypes, setActivityTypes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [activity, setActivity] = useState(defaultActivity);

  useEffect(() => {
    fetchActivityTypes(setActivityTypes, setLoading);
  }, [setLoading]);

  const editing = !!initialActivity;

  // TODO:: This sucks
  const upsert = async (
    values,
    apiCall,
    needsDate = true,
    objName = 'activity'
  ) => {
    setLoading(true);
    // First, we need to grab the total for duration
    const postValues = {
      ...values,
      duration: values.duration.total,
      activity_date: new Date(`${values.activity_date} UTC`),
    };
    // TODO:: Clean this up
    if (!needsDate) delete postValues.activity_date;

    try {
      await apiCall(postValues);
      message.success(
        `Successfully ${editing ? 'updated' : 'created'} ${objName}!`
      );
      refetch();
    } catch (err) {
      notification.error({
        message: 'Unexpected Error',
        description: `Error: ${err}`,
      });
    } finally {
      setLoading(false);
    }
  };

  const onFinish = async values => {
    if (editing) {
      const activityToPut = { ...values };
      activityToPut.id = initialActivity.id;
      await upsert(activityToPut, putActivity);
      modalClose();
    } else {
      await upsert(values, postActivity);
    }
    form.resetFields();
    setActivity(form.getFieldsValue());
  };

  const onSaveQuickAdd = async () => {
    try {
      const values = form.getFieldsValue();
      if ((values.duration && values.duration.total <= 1) || !values.name) {
        message.error('Duration and Name are required for saving a quick add.');
        return;
      }
      await upsert(values, postQuickAdd, false, 'Quick Add');
      refetch();
    } catch (e) {
      console.log(e);
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
    <div className="outer-form-wrapper">
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
            duration: {
              hours: null,
              minutes: null,
              seconds: null,
              total: 0,
            },
            ...initialActivity,
          }}
        >
          {/* ============= NAME ============= */}
          <Item label="Activity Name" name="name">
            {/* Data LP ignore is to stop LastPass from trying to fill the form */}
            <Input
              data-lpignore="true"
              className="fullWidth"
              placeholder={`${user.given_name}'s Favorite Run`}
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
      <CalculatedActivityFields activity={activity} />
      <div className="button-row">
        <Button
          className="button-row-item"
          onClick={onSubmit}
          icon={editing ? <EditOutlined /> : <RocketOutlined rotate={45} />}
          type="primary"
          loading={loading}
          disabled={loading}
        >
          {editing ? 'Edit' : 'Submit'}
        </Button>
        <Button
          className="button-row-item"
          ghost
          onClick={onReset}
          type="primary"
        >
          Reset
        </Button>
        <Button
          className="button-row-item"
          onClick={onSaveQuickAdd}
          type="dashed"
        >
          Save for Quick-Add
        </Button>
      </div>
    </div>
  );
}
