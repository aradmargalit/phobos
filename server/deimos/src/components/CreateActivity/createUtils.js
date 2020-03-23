import { message, notification } from 'antd';

import { postActivity, postQuickAdd } from '../../apis/phobos-api';

const onSuccess = entity => message.success(`Successfully created ${entity}!`);
const onError = err =>
  notification.error({
    message: 'Unexpected Error',
    description: `Error: ${err}`,
  });

export const onSubmit = form => {
  form.submit();
};

export const onReset = async (form, setActivity) => {
  await form.resetFields();
  setActivity(form.getFieldsValue());
};

const commonPost = async (setLoading, apiCall, values, refetch) => {
  setLoading(true);
  try {
    await apiCall(values);
    refetch();
  } catch (err) {
    onError(err);
  } finally {
    setLoading(false);
  }
};

export const onFinish = async (
  values,
  setLoading,
  refetch,
  form,
  setActivity
) => {
  const postValues = {
    ...values,
    duration: values.duration.total,
    activity_date: new Date(`${values.activity_date} UTC`),
  };
  await commonPost(setLoading, postActivity, postValues, refetch);
  onSuccess('activity');
  form.resetFields();
  setActivity(form.getFieldsValue());
};

export const onSaveQuickAdd = async (form, refetch, setLoading) => {
  const values = form.getFieldsValue();
  // Try to leverage inbuilt validation
  try {
    await form.validateFields();
  } catch (e) {
    return;
  }

  // We can use the values as-is, expect for the duration
  const postValues = {
    ...values,
    duration: values.duration.total,
  };

  // Custom validation
  if (!values.name) {
    message.error('Name is required for saving a quick add.');
    return;
  }

  await commonPost(setLoading, postQuickAdd, postValues, refetch);
};
