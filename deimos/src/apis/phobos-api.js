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

const protectedUpsert = async (endpoint, method, body) => {
  const res = await fetch(`${BACKEND_URL}${endpoint}`, { ...options, method, body: JSON.stringify(body) });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};

export const fetchActivityTypes = async (setActivityTypes, setLoading) => {
  await protectedGet(setActivityTypes, setLoading, '/metadata/activity_types', 'activity_types');
};

export const fetchActivities = async (setActivities, setLoading) => {
  await protectedGet(setActivities, setLoading, '/private/activities', 'activities');
};

export const postActivity = async (activity) => {
  await protectedUpsert('/private/activities', 'POST', activity);
};

export const putActivity = async (activity) => {
  await protectedUpsert(`/private/activities/${activity.id}`, 'PUT', activity);
};

export const deleteActivity = async (id) => {
  const res = await fetch(`${BACKEND_URL}/private/activities/${id}`, { ...options, method: 'DELETE' });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};
