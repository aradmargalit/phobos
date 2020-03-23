import { message, notification } from 'antd';

import { putActivity } from '../../apis/phobos-api';

const onSuccess = () => message.success(`Successfully updated activity!`);
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

const commonPut = async (setLoading, apiCall, values, refetch) => {
  setLoading(true);
  try {
    await apiCall(values);
    onSuccess();
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
  activityId,
  modalClose
) => {
  const activityToPut = { ...values };
  activityToPut.id = activityId;
  const putValues = {
    ...activityToPut,
    duration: values.duration.total,
  };
  await commonPut(setLoading, putActivity, putValues, refetch);
  modalClose();
};
