/* eslint-disable camelcase */
import { BACKEND_URL } from '../constants';

const options = {
  method: 'POST',
  mode: 'cors',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
};

const protectedGet = async (setValue, setLoading, endpoint, dataKey) => {
  // Make sure to include the cookie with the request!
  const res = await fetch(`${BACKEND_URL}${endpoint}`, {
    credentials: 'include',
  });

  const response = await res.json();
  setValue(response[dataKey]);
  setLoading(false);
};

export const fetchActivityTypes = async (setActivityTypes, setLoading) => {
  await protectedGet(setActivityTypes, setLoading, '/metadata/activity_types', 'activity_types');
};

export const fetchActivities = async (setActivities, setLoading) => {
  await protectedGet(setActivities, setLoading, '/private/activities', 'activities');
};

export const postActivity = async (activity) => {
  const res = await fetch(`${BACKEND_URL}/private/activities`, { ...options, method: 'POST', body: JSON.stringify(activity) });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};

export const putActivity = async (activity) => {
  const res = await fetch(`${BACKEND_URL}/private/activities/${activity.id}`, { ...options, method: 'PUT', body: JSON.stringify(activity) });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};

export const deleteActivity = async (id) => {
  const res = await fetch(`${BACKEND_URL}/private/activities/${id}`, { ...options, method: 'DELETE' });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};
