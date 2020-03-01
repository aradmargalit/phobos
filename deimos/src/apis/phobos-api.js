/* eslint-disable camelcase */
import { BACKEND_URL } from '../constants';

const postOptions = {
  method: 'POST',
  mode: 'cors',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
};

const protectedGet = async (setValue, setLoading, endpoint, dataKey) => {
  // Make sure to include the cookie with the request!
  setLoading(true);
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
  const res = await fetch(`${BACKEND_URL}/private/activities`, { ...postOptions, body: JSON.stringify(activity) });
  const { error } = await res.json();
  if (error) {
    throw error;
  }
};
