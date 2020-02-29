import React, { useEffect, useState, useContext } from 'react';
import { Button, Spin, notification } from 'antd';
import { GoogleOutlined, LogoutOutlined } from '@ant-design/icons';
import { BACKEND_URL } from '../../constants';
import UserContext from '../../contexts/UserContext';
import './IdentityButton.scss';

export default function IdentityButton() {
  const {
    user, setUser, loading, setLoading,
  } = useContext(UserContext);

  const [errors, setErrors] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      // Make sure to include the cookie with the request!
      try {
        const res = await fetch(`${BACKEND_URL}/private/users/current`, {
          credentials: 'include',
        });

        res.json().then(({ user: respUser }) => {
          setUser(respUser);
          setLoading(false);
        });
      } catch (err) {
        notification.error({
          message: 'Unexpected Error',
          description: `Error: ${err}`,
        });
        setErrors('API Dead?');
        setLoading(false);
      }
    };

    fetchData();
  }, [setUser, setLoading]);

  if (errors) return <p>{errors}</p>;
  if (loading) return <Spin />;

  return user ? (
    <div>
      <h1 className="ant-page-header-heading-title welcome">Welcome back,</h1>
      <h1 className="ant-page-header-heading-title primary">
        {user.given_name}
      </h1>
      <Button
        icon={<LogoutOutlined />}
        href={`${BACKEND_URL}/users/logout`}
        type="danger"
        ghost
      >
        Logout
      </Button>
    </div>
  ) : (
    <Button icon={<GoogleOutlined />} href={`${BACKEND_URL}/auth/google`}>
      Login with Google
    </Button>
  );
}
