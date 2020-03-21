/* eslint-disable camelcase */

const options = {
  method: 'POST',
  mode: 'cors',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
};

const protectedGet = async (setValue, setLoading, endpoint, dataKey = null) => {
  // Make sure to include the cookie with the request!
  const res = await fetch(endpoint, {
    credentials: 'include',
  });

  const response = await res.json();
  setValue(dataKey ? response[dataKey] : response);
  setLoading(false);
};

const protectedUpsert = async (endpoint, method, body) => {
  const res = await fetch(endpoint, {
    ...options,
    method,
    body: JSON.stringify(body),
  });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};

export const fetchActivityTypes = async (setActivityTypes, setLoading) => {
  await protectedGet(
    setActivityTypes,
    setLoading,
    '/metadata/activity_types',
    'activity_types'
  );
};

export const fetchActivities = async (setActivities, setLoading) => {
  await protectedGet(
    setActivities,
    setLoading,
    '/private/activities',
    'activities'
  );
};

export const postActivity = async activity => {
  await protectedUpsert('/private/activities', 'POST', activity);
};

export const putActivity = async activity => {
  await protectedUpsert(`/private/activities/${activity.id}`, 'PUT', activity);
};

export const deleteActivity = async id => {
  const res = await fetch(`/private/activities/${id}`, {
    ...options,
    method: 'DELETE',
  });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};

export const fetchMonthlySums = async (setMonthlySums, setLoading) => {
  await protectedGet(setMonthlySums, setLoading, '/private/activities/monthly');
};

export const fetchStatistics = async (setStats, setLoading) => {
  await protectedGet(setStats, setLoading, '/private/statistics');
};

export const postQuickAdd = async values => {
  await protectedUpsert('/private/quick_adds', 'POST', values);
};

export const fetchQuickAdds = async (setQuickAdds, setLoading) => {
  await protectedGet(setQuickAdds, setLoading, '/private/quick_adds');
};
