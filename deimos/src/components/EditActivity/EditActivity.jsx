/* eslint-disable react/jsx-curly-newline */
import { Form } from 'antd';
import React, { useEffect, useState } from 'react';

import { fetchActivityTypes } from '../../apis/phobos-api';
import AddActivityForm from '../AddActivityForm';
import CalculatedActivityFields from '../CalculatedActivityFields';
import FormButtons from '../FormButtons';
import { onFinish, onReset, onSubmit } from './editUtils';

export default function EditActivity({ refetch, initialActivity, modalClose }) {
  const [activity, setActivity] = useState(null);
  const [loading, setLoading] = useState(false);
  const [activityTypes, setActivityTypes] = useState({
    payload: [],
    loading: true,
    errors: null,
  });
  const [form] = Form.useForm();

  useEffect(() => {
    fetchActivityTypes(setActivityTypes);
  }, []);

  return (
    <div className="outer-form-wrapper">
      <AddActivityForm
        refetch={refetch}
        form={form}
        setActivity={setActivity}
        activityTypes={activityTypes.payload}
        initialActivity={initialActivity}
        onFinish={values => onFinish(values, setLoading, refetch, initialActivity.id, modalClose)}
      />
      <CalculatedActivityFields activity={activity} />
      <FormButtons
        editing
        loading={loading}
        onSubmit={() => onSubmit(form)}
        onReset={() => onReset(form, setActivity)}
      />
    </div>
  );
}
