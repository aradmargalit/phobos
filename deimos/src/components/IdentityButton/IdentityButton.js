import React from 'react';
import { Button } from 'antd';

const googleIcon =
  'https://pluspng.com/img-png/google-logo-png-google-logo-icon-png-transparent-background-1000.png';

export default function IdentityButton() {
  let BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

  return (
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
