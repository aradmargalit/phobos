import React from 'react';
import { Button } from 'antd';

export default function Login() {
  return (
    <div>
      <Button href="/auth/google">Login with Google</Button>
      <Button href="/currentUser">Current User</Button>
    </div>
  );
}
