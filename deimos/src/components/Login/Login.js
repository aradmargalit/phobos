import React from 'react';
import { Button } from 'antd';

export default function Login() {
  let BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

  return (
    <div>
      <Button href={`${BACKEND_URL}/auth/google`}>Login with Google</Button>
      <Button href={`${BACKEND_URL}/currentUser`}>Current User</Button>
    </div>
  );
}
