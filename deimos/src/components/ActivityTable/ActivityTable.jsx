import './ActivityTable.scss';

import {
  DeleteOutlined,
  EditOutlined,
} from '@ant-design/icons';
import {
  Button, Empty, Modal,
  Popconfirm,
  Spin, Table,
} from 'antd';
import { snakeCase as _snakeCase } from 'lodash';
import moment from 'moment';
import React, { useState } from 'react';

import { deleteActivity } from '../../apis/phobos-api';
import { formatDate, minutesToHMS, toLocalDate } from '../../utils/dataFormatUtils';
import AddActivityForm from '../AddActivityForm';
import TableSearch from '../TableSearch';

const toCol = (name, render) => {
  const snakeName = _snakeCase(name);
  const col = { title: name, dataIndex: snakeName };
  if (render) {
    col.render = render;
  }
  return col;
};


export default function ActivityTable({
  loading, activityTypes, activities, refetch,
}) {
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingActivity, setEditingActivity] = useState(null);
  const [tableActivities, setTableActivities] = useState(activities);
  const [filtered, setFiltered] = useState(false);

  if (loading || !activityTypes.length) return <Spin />;
  if (!activities) return <Empty description="No activities...yet!" />;
  if (!tableActivities || (!filtered && tableActivities.length !== activities.length)) {
    setTableActivities(activities);
  }


  const renderEditButtons = (activity) => (
    <Button onClick={() => {
      const toEdit = { ...activity };
      toEdit.activity_date = moment(activity.activity_date);
      setEditingActivity(toEdit);
      setEditModalVisible(true);
    }}
    >
      Edit
    </Button>
  );

  const columns = [
    {
      title: 'No.',
      dataIndex: 'id',
    },
    toCol('Name'),
    // We want to format this one as the time it was entered, since it's time is 00:00:00
    // and we don't want to cross date boundaries by converting timezones
    toCol('Activity Date', formatDate),
    {
      title: 'Activity Type',
      dataIndex: 'activity_type_id',
      render: (id) => {
        const aT = activityTypes.find((at) => at.id === id);
        return aT ? aT.name : '';
      },
    },
    toCol('Duration', minutesToHMS),
    toCol('Distance', (distance, record) => ((distance > 0) ? `${distance} ${record.unit}` : '-')),
    toCol('Created At', toLocalDate),
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
      render: ({ id }) => (
        <Popconfirm
          title="Are you sure?"
          okText="Delete"
          icon={<DeleteOutlined style={{ color: 'red' }} />}
          onConfirm={async () => {
            await deleteActivity(id);
            refetch();
          }}
        >
          <Button ghost type="danger">Delete</Button>
        </Popconfirm>
      ),
    },
  ];

  return (
    <div>
      <TableSearch
        activityTypes={activityTypes}
        setFiltered={setFiltered}
        tableActivities={tableActivities}
        setTableActivities={setTableActivities}
      />
      <Table
        rowKey="id"
        dataSource={tableActivities}
        columns={columns}
      />
      <Modal
        title="Edit Activity"
        visible={editModalVisible}
        onOk={() => setEditModalVisible(false)}
        onCancel={() => setEditModalVisible(false)}
        width={600}
        footer={null}
        destroyOnClose
      >
        <AddActivityForm
          activityTypes={activityTypes}
          loading={loading}
          refetch={refetch}
          initialActivity={editingActivity}
          modalClose={() => setEditModalVisible(false)}
        />
      </Modal>
    </div>
  );
}
