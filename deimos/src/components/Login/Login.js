import React from 'react';
import { Button } from 'antd';

export default function Login() {
  return (
    <div>
      <Button href="/auth/google">Login with Google</Button>
      <Button href="/currentUser">Current UserS</Button>
      <h1>BEURL is: {process.env.API_URL}</h1>
    </div>
  );
}
