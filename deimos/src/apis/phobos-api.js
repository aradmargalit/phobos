import { BACKEND_URL } from '../constants';

const postOptions = {
  method: 'POST',
  mode: 'cors',
  credentials: 'include',
  headers: {
    'Content-Type': 'application/json',
  },
};

// eslint-disable-next-line import/prefer-default-export
export const fetchActivityTypes = async (setActivityTypes, setLoading) => {
  // Make sure to include the cookie with the request!
  const res = await fetch(`${BACKEND_URL}/metadata/activity_types`, {
    credentials: 'include',
  });

  res.json().then(({ activity_types: respTypes }) => {
    setActivityTypes(respTypes);
    setLoading(false);
  });
};

export const postActivity = async (activity) => {
  const res = await fetch(`${BACKEND_URL}/private/activities`, { ...postOptions, body: JSON.stringify(activity) });
  return res.json();
};
