import { Modal, Button } from 'antd';
import React, { useState } from 'react';
import './AddActivityModal.scss';
import { PlusCircleOutlined } from '@ant-design/icons';
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
      <Button icon={<PlusCircleOutlined />} type="primary" onClick={showModal}>
        Add Activity
      </Button>
      <Modal
        title="Add New Activity"
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
