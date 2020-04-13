import { putActivity } from '../../apis/phobos-api';

export const onSubmit = form => {
  form.submit();
};

export const onReset = async (form, setActivity) => {
  await form.resetFields();
  setActivity(form.getFieldsValue());
};

const commonPut = async (setLoading, apiCall, values, refetch) => {
  setLoading(true);
  await apiCall(values);
  refetch();
  setLoading(false);
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
