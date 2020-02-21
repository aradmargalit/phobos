import { Modal, Button, Icon } from 'antd';
import React, { useState } from 'react';
import './AddActivityModal.scss';
import AddActivityForm from '../AddActivityForm';

export default function AddActivityModal() {
  const [visible, setVisible] = useState(false);

  const showModal = () => {
    setVisible(true);
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
        title={(
          <span>
            Add New Activity
            <Icon className="form-icon" type="form" />
          </span>
        )}
        visible={visible}
        onCancel={handleCancel}
        destroyOnClose
        footer={null}
      >
        <AddActivityForm />
      </Modal>
    </div>
  );
}
