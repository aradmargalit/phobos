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
  const [loading, setLoading] = useState(true);
  const [activityTypes, setActivityTypes] = useState([]);

  const wrappedFinish = values =>
    onFinish(values, setLoading, refetch, form, setActivity);

  useEffect(() => {
    fetchActivityTypes(setActivityTypes, setLoading);
  }, [setLoading]);

  return (
    <div className="outer-form-wrapper">
      <AddActivityForm
        refetch={refetch}
        form={form}
        user={user}
        setActivity={setActivity}
        activityTypes={activityTypes}
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

/*
  const defaultActivity = {
    unit: 'miles',
    duration: {
      hours: null,
      minutes: null,
      seconds: null,
      total: 0,
    },
  };
*/
