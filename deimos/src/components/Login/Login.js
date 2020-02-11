import React from 'react';
import { Button } from 'antd';
import './Login.scss';

export default function Login() {
  let BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

  return (
    <div className="button-container">
      <Button href={`${BACKEND_URL}/auth/google`}>Login with Google</Button>
      <Button href={`${BACKEND_URL}/private/users/current`}>
        Current User
      </Button>
      <Button href={`${BACKEND_URL}/users/logout`}>Logout</Button>
    </div>
  );
}
