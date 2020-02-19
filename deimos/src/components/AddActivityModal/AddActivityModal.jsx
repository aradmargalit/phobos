import { Modal, Button } from 'antd';
import React, { useState } from 'react';
import './AddActivityModal.scss';
import AddActivityForm from '../AddActivityForm';

export default function AddActivityModal() {
  const [visible, setVisible] = useState(false);

  const showModal = () => {
    setVisible(true);
  };

  const handleOk = (e) => {
    console.log(e);
    setVisible(false);
  };

  const handleCancel = (e) => {
    console.log(e);
    setVisible(false);
  };


  return (
    <div>
      <Button icon="plus-circle" type="primary" onClick={showModal}>
        Add Activity
      </Button>
      <Modal
        title="Add New Activity"
        visible={visible}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <AddActivityForm />
      </Modal>
    </div>
  );
}
