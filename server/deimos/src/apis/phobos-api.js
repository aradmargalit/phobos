/* eslint-disable camelcase */

import { message, notification } from 'antd';

// Common Options
const options = {
  headers: {
    'Content-Type': 'application/json',
  },
};

const protectedGet = async (setValue, endpoint, dataKey = null) => {
  const res = await fetch(endpoint);
  if (!res || !res.ok) {
    setValue(curr => ({ ...curr, loading: false, errors: res.statusText }));
    return;
  }

  const response = await res.json();
  setValue({
    payload: dataKey ? response[dataKey] : response,
    loading: false,
    errors: null,
  });
};

const protectedUpsert = async (endpoint, method, body) => {
  const res = await fetch(endpoint, {
    ...options,
    method,
    body: JSON.stringify(body),
  });

  if (!res || !res.ok) {
    notification.error({
      message: `Could not ${method} to ${endpoint}...`,
      description: res.statusText,
      duration: 2,
    });
    return;
  }

  message.success(`Successfully ${method === 'POST' ? 'created' : 'updated'}!`);
};

export const protectedDelete = async endpoint => {
  const res = await fetch(endpoint, {
    ...options,
    method: 'DELETE',
  });

  if (!res || !res.ok) {
    notification.error({
      message: `Could not DELETE ${endpoint}...`,
      description: res.statusText,
      duration: 2,
    });
    return;
  }

  message.success(`Successfully deleted!`);
};

// Activity Types
export const fetchActivityTypes = async setActivityTypes => {
  await protectedGet(
    setActivityTypes,
    '/metadata/activity_types',
    'activity_types'
  );
};

// Activities
export const fetchActivities = async setActivities => {
  await protectedGet(setActivities, '/private/activities', 'activities');
};

export const postActivity = async activity => {
  await protectedUpsert('/private/activities', 'POST', activity);
};

export const putActivity = async activity => {
  await protectedUpsert(`/private/activities/${activity.id}`, 'PUT', activity);
};

export const deleteActivity = async id => {
  await protectedDelete(`/private/activities/${id}`);
};

// Statistics
export const fetchMonthlySums = async setMonthlySums => {
  await protectedGet(setMonthlySums, '/private/activities/monthly');
};

export const fetchStatistics = async setStats => {
  const utcOffset = new Date().getTimezoneOffset() / 60;
  await protectedGet(setStats, `/private/statistics?utc_offset=${utcOffset}`);
};

// Quick Adds
export const postQuickAdd = async values => {
  await protectedUpsert('/private/quick_adds', 'POST', values);
};

export const fetchQuickAdds = async setQuickAdds => {
  await protectedGet(setQuickAdds, '/private/quick_adds');
};

export const deleteQuickAdd = async id => {
  await protectedDelete(`/private/quick_adds/${id}`);
};

// Strava Calls
export const fetchStravaStats = async setStravaStatistics => {
  await protectedGet(setStravaStatistics, '/strava/statistics');
};

// User
export const fetchUser = async setUser => {
  await protectedGet(setUser, '/private/users/current', 'user');
};

export const deauthStrava = async () => {
  const res = await fetch('/strava/deauth');
  if (!res || !res.ok) {
    notification.error({
      message: 'Could not deauthorize Strava.',
      description: 'Please try again!',
      duration: 2,
    });
  }
};
