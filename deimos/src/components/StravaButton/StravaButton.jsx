import './StravaButton.scss';

import {
  ApiOutlined,
  CheckCircleOutlined,
  DeleteOutlined,
  EditOutlined,
  PlusCircleOutlined,
  ScissorOutlined,
} from '@ant-design/icons';
import { Button, Modal } from 'antd';
import React, { useContext, useState } from 'react';

import { deauthStrava, fetchUser } from '../../apis/phobos-api';
import { UserContext } from '../../contexts';

const registeredModalContent = (setVisible, setUser) => (
  <div>
    <h2>Stop Strava Updates?</h2>
    <h4>Any new activities you add to Strava will no longer be synced to Phobos.</h4>

    <Button
      type="danger"
      icon={<ScissorOutlined />}
      onClick={async () => {
        await deauthStrava();
        await fetchUser(setUser);
        setVisible(false);
      }}
    >
      Disable Updates
    </Button>
  </div>
);

const unregisteredModalContent = () => (
  <div className="strava-unregistered-content">
    <h2>Strava Sync</h2>
    <p>
      <PlusCircleOutlined />
      Any new activities you add to Strava will be synced to Phobos.
    </p>
    <p>
      <EditOutlined />
      Updating activities in Strava will automatically update them here.
    </p>
    <p>
      <DeleteOutlined />
      Deleting an activity in Strava deletes it in Phobos.
    </p>
    <br />
    <Button href="/strava/auth" type="primary" icon={<ApiOutlined />}>
      Enable Strava Sync
    </Button>
  </div>
);

export default function StravaButton({ registered, loading }) {
  const { setUser } = useContext(UserContext);

  const [visible, setVisible] = useState(false);

  return (
    <>
      <Button
        disabled={loading}
        className={`strava-button${registered ? '--registered' : ''}`}
        onClick={() => setVisible(true)}
      >
        {registered ? 'Connected with Strava' : 'Connect with Strava'}
        {registered && <CheckCircleOutlined />}
      </Button>
      <Modal visible={visible} onCancel={() => setVisible(false)} footer={null}>
        {registered ? registeredModalContent(setVisible, setUser) : unregisteredModalContent()}
      </Modal>
    </>
  );
}
