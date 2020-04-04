import React, { useContext, useEffect, useState } from 'react';

import { fetchActivityTypes } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';
import AddActivityForm from '../AddActivityForm';
import CalculatedActivityFields from '../CalculatedActivityFields';
import FormButtons from '../FormButtons';
import { onFinish, onReset, onSaveQuickAdd, onSubmit } from './createUtils';

export default function CreateActivity({
  form,
  refetch,
  activity,
  setActivity,
}) {
  const { user } = useContext(UserContext);
  const [loading, setLoading] = useState(false);
  const [activityTypes, setActivityTypes] = useState({
    payload: [],
    loading: false,
    errors: null,
  });

  const wrappedFinish = values =>
    onFinish(values, setLoading, refetch, form, setActivity);

  useEffect(() => {
    fetchActivityTypes(setActivityTypes);
  }, [activityTypes]);

  return (
    <div className="outer-form-wrapper">
      <AddActivityForm
        refetch={refetch}
        form={form}
        user={user.payload}
        setActivity={setActivity}
        activityTypes={activityTypes.payload}
        onFinish={wrappedFinish}
      />
      <CalculatedActivityFields activity={activity} />
      <FormButtons
        loading={loading}
        onSubmit={() => onSubmit(form)}
        onReset={() => onReset(form, setActivity)}
        onSaveQuickAdd={() => onSaveQuickAdd(form, refetch, setLoading)}
      />
    </div>
  );
}
