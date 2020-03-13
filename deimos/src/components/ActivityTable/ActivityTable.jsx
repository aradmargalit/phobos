import './ActivityTable.scss';

import { DeleteOutlined, EditOutlined } from '@ant-design/icons';
import { Button, Empty, Input, Modal, Popconfirm, Spin, Table } from 'antd';
import { debounce as _debounce } from 'lodash';
import moment from 'moment';
import React, { useState } from 'react';

import { deleteActivity } from '../../apis/phobos-api';
import { formatDate, minutesToHMS } from '../../utils/dataFormatUtils';
import AddActivityForm from '../AddActivityForm';
import { filterActivities, makeDurationBreakdown, toCol } from './tableUtils';

const { Search } = Input;

export default function ActivityTable({ loading, activities, refetch }) {
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingActivity, setEditingActivity] = useState(null);
  const [searchTerm, setSearchTerm] = useState(null);

  if (loading) return <Spin />;
  if (!activities) return <Empty description="No activities...yet!" />;

  const bouncedSetSearchTerm = _debounce(setSearchTerm, 500);

  const onChangeHandler = e => {
    if (!e.target.value || !e.target.value.length) {
      setSearchTerm(null);
      return;
    }
    bouncedSetSearchTerm(e.target.value);
  };

  const renderEditButtons = activity => (
    <Button
      onClick={() => {
        const toEdit = { ...activity };
        toEdit.activity_date = moment(activity.activity_date);
        // This is dumb, but we need to set the full duration object until I get smarter
        toEdit.duration = makeDurationBreakdown(activity.duration);
        setEditingActivity(toEdit);
        setEditModalVisible(true);
      }}
    >
      Edit
    </Button>
  );

  const confirmDelete = ({ id }) => (
    <Popconfirm
      title="Are you sure?"
      okText="Delete"
      icon={<DeleteOutlined style={{ color: 'red' }} />}
      onConfirm={async () => {
        await deleteActivity(id);
        refetch();
      }}
    >
      <Button ghost type="danger">
        Delete
      </Button>
    </Popconfirm>
  );

  const closeModal = () => setEditModalVisible(false);
  const columns = [
    {
      title: 'No.',
      dataIndex: 'idx',
      render: (text, record, idx) => activities.length - idx,
    },
    { ...toCol('Name'), width: 250 },
    // We want to format this one as the time it was entered, since it's time is 00:00:00
    // and we don't want to cross date boundaries by converting timezones
    toCol('Activity Date', formatDate),
    {
      title: 'Activity Type',
      dataIndex: ['activity_type', 'name'],
    },
    toCol('Duration', minutesToHMS),
    toCol('Distance', (distance, record) =>
      distance > 0 ? `${distance} ${record.unit}` : '-'
    ),
    {
      title: <EditOutlined />,
      key: 'edit',
      align: 'center',
      render: renderEditButtons,
    },
    {
      title: <DeleteOutlined />,
      key: 'delete',
      align: 'center',
      render: confirmDelete,
    },
  ];

  return (
    <div>
      <Search
        className="search-bar"
        allowClear
        placeholder="Search by name, type, or date..."
        onSearch={onChangeHandler}
        onChange={onChangeHandler}
        width="50%"
      />
      <Table
        rowKey="id"
        dataSource={filterActivities(searchTerm, activities)}
        columns={columns}
      />
      <Modal
        title="Edit Activity"
        visible={editModalVisible}
        onOk={closeModal}
        onCancel={closeModal}
        width={750}
        footer={null}
        destroyOnClose
      >
        <AddActivityForm
          refetch={refetch}
          initialActivity={editingActivity}
          modalClose={closeModal}
        />
      </Modal>
    </div>
  );
}
