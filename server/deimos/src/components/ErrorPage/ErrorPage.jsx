import { RedoOutlined } from '@ant-design/icons';
import { Alert, Button, notification } from 'antd';
import React from 'react';
import { withRouter } from 'react-router-dom';

const phoberror = require('./phoberror.png');

const messageMap = {
  500: 'Our system is having trouble. You can retry, or come back later',
  502: 'Our system is having trouble. You can retry, or come back later',
  504: 'Our system is having trouble. You can retry, or come back later',
  404: "we weren't able to find that page",
  strava:
    'Please make sure you allow Phobos to read your activities and try again.',
};

function ErrorPage({ history }) {
  // Every time we land on the error page, we need to check if the specific error is included in the URL
  const errorType = window.location.href.split('/').slice(-1);
  const errorMessage = `It looks like something went wrong. Sorry about that! ${messageMap[
    errorType
  ] || ''}`;
  return (
    <div className="error-container">
      <Alert
        style={{ width: '50%', margin: '50px auto' }}
        message={errorMessage}
        type="error"
      />
      <div>
        <img
          width={350}
          style={{ marginBottom: '30px' }}
          src={phoberror}
          alt="sad phoebe"
        />
      </div>
      <Button
        icon={<RedoOutlined />}
        type="primary"
        onClick={() => {
          notification.info({
            message: 'Checking if things are fixed!',
            duration: 2,
          });
          history.push('/home');
        }}
      >
        Retry
      </Button>
    </div>
  );
}

export default withRouter(ErrorPage);
