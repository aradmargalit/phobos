import React, { useEffect, useState } from 'react';
import { Button, Spin } from 'antd';
import './IdentityButton.scss';

const googleIcon =
  'https://pluspng.com/img-png/google-logo-png-google-logo-icon-png-transparent-background-1000.png';

let BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

export default function IdentityButton() {
  const [loading, setLoading] = useState(true);
  const [errors, setErrors] = useState(false);
  const [user, setUser] = useState(null);

  async function fetchData() {
    setLoading(true);
    // Make sure to include the cookie with the request!
    const res = await fetch(`${BACKEND_URL}/private/users/current`, {
      credentials: 'include'
    });

    res
      .json()
      .then(res => setUser(res.user))
      .catch(err => setErrors(err))
      .finally(() => setLoading(false));
  }

  useEffect(() => {
    fetchData();
  }, []);

  if (errors) return <h1>{errors}</h1>;
  if (loading) return <Spin />;

  return user ? (
    <div>
      <h1 className="ant-page-header-heading-title welcome">Welcome back,</h1>
      <h1 className="ant-page-header-heading-title primary">
        {user.given_name}
      </h1>
      <Button href={`${BACKEND_URL}/users/logout`} type="danger" ghost>
        Logout
      </Button>
    </div>
  ) : (
    <Button href={`${BACKEND_URL}/auth/google`}>
      <img
        alt="google-icon"
        src={googleIcon}
        width="20"
        style={{ padding: '0 5px 0 0' }}
      />
      Login with Google
    </Button>
  );
}
