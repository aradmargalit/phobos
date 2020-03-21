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

export const onReset = (form, setActivity) => {
  form.resetFields();
  setActivity(form.getFieldsValue());
};

const commonPost = async (setLoading, apiCall, values, refetch, entity) => {
  setLoading(true);
  try {
    await apiCall(values);
    onSuccess(entity);
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
  await commonPost(setLoading, postActivity, postValues, refetch, 'activity');
  form.resetFields();
  setActivity(form.getFieldsValue());
};

// TODO Round 2 fefactor
export const onSaveQuickAdd = async (form, refetch, setLoading) => {
  const values = form.getFieldsValue();
  // Try to leverage inbuilt validation
  try {
    await form.validateFields();
  } catch (e) {
    return;
  }
  const postValues = {
    ...values,
    duration: values.duration.total,
  };

  // Custom validation
  if (!values.name) {
    message.error('Name is required for saving a quick add.');
    return;
  }

  await commonPost(setLoading, postQuickAdd, postValues, refetch, 'quick add');
};

/*
    
  */

/*

  */

/*
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
  */

// export const upsert = async (
//   setLoading,
//   onSuccess,
//   onError,
//   refetch,
//   values,
//   apiCall,
//   needsDate = true
// ) => {
//   setLoading(true);
//   // First, we need to grab the total for duration
//   const postValues = {
//     ...values,
//     duration: values.duration.total,
//     activity_date: new Date(`${values.activity_date} UTC`),
//   };
//   // TODO:: Clean this up
//   if (!needsDate) delete postValues.activity_date;

//   try {
//     await apiCall(postValues);
//     onSuccess();
//     refetch();
//   } catch (err) {
//     onError();
//   } finally {
//     setLoading(false);
//   }
// };
