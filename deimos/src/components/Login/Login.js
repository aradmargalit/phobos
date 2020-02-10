import React from 'react';
import { Button } from 'antd';

export default function Login() {
  return (
    <div>
      <Button href={`${window._env_.API_URL}/auth/google`}>Login with Google</Button>
      <Button href={`${window._env_.API_URL}/currentUser`}>Current User</Button>
      <h1>Please! {window._env_.API_URL}</h1>
    </div>
  );
}
