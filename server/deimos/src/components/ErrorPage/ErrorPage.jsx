import { RedoOutlined } from '@ant-design/icons';
import { Alert, Button, Spin } from 'antd';
import React, { useContext, useState } from 'react';
import { withRouter } from 'react-router-dom';

import { fetchUser } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';

const phoberror = require('./phoberror.png');

const messageMap = {
  500: 'Our system is having trouble. You can retry, or come back later',
  502: 'Our system is having trouble. You can retry, or come back later',
  504: 'Our system is having trouble. You can retry, or come back later',
  404: "We weren't able to find that page.",
  strava:
    'Please make sure you allow Phobos to read your activities and try again.',
};

function ErrorPage({ history }) {
  const { setUser } = useContext(UserContext);

  const [loading, setLoading] = useState(false);

  if (loading) return <Spin />;

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
        onClick={async () => {
          setLoading(true);
          setTimeout(() => {
            setLoading(false);
          }, 500);
          await fetchUser(setUser);
          history.push('/home');
        }}
      >
        Retry
      </Button>
    </div>
  );
}

export default withRouter(ErrorPage);
