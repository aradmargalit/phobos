import './ActivityTable.scss';

import { DeleteOutlined, EditOutlined } from '@ant-design/icons';
import { Button, Empty, Input, Modal, Popconfirm, Spin, Table } from 'antd';
import { debounce as _debounce } from 'lodash';
import moment from 'moment';
import React, { useState } from 'react';

import { deleteActivity } from '../../apis/phobos-api';
import { formatDate, minutesToHMS } from '../../utils/dataFormatUtils';
import { makeDurationBreakdown } from '../../utils/durationUtils';
import EditActivity from '../EditActivity';
import {
  activityTypeSorter,
  dateSorter,
  distanceSorter,
  durationSorter,
  heartRateSorter,
  nameSorter,
  numberSorter,
} from './sortUtils';
import { filterActivities, toCol } from './tableUtils';

const stravaIcon = require('./strava.png');

const { Search } = Input;

export default function ActivityTable({ loading, activities, refetch }) {
  const [editModalVisible, setEditModalVisible] = useState(false);
  const [editingActivity, setEditingActivity] = useState(null);
  const [searchTerm, setSearchTerm] = useState(null);

  if (loading) return <Spin />;
  if (!activities) return <Empty description="No activities...yet!" />;

  // Debounce the search term to avoid pointless, expensive re-renders
  const bouncedSetSearchTerm = _debounce(setSearchTerm, 350);

  const onSearchHandler = term => {
    // If there's no search term, we want to set null in order to do the right thing later
    if (!term || !term.length) {
      bouncedSetSearchTerm(null);
      return;
    }
    bouncedSetSearchTerm(term);
  };

  const onChangeHandler = e => {
    onSearchHandler(e.target.value);
  };

  const renderEditButtons = activity => (
    <Button
      onClick={() => {
        const toEdit = { ...activity };
        toEdit.activity_date = moment(activity.activity_date);
        // We need to re-create the duration breakdown from the total
        // in order to properly display in the form
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

  const toStravaIcon = stravaID => {
    if (stravaID < 1) return null;

    return (
      <a
        target="_blank"
        rel="noopener noreferrer"
        href={`https://www.strava.com/activities/${stravaID}`}
      >
        <img width={40} alt="strava icon" src={stravaIcon} />
      </a>
    );
  };

  const columns = [
    {
      title: '',
      dataIndex: 'strava_id',
      width: 50,
      render: token => toStravaIcon(token),
    },
    {
      title: 'No.',
      dataIndex: 'logical_index',
      width: 70,
      sorter: numberSorter,
    },
    { ...toCol('Name'), width: 250, sorter: nameSorter },
    // We want to format this one as the time it was entered, since it's time is 00:00:00
    // and we don't want to cross date boundaries by converting timezones
    { ...toCol('Activity Date', formatDate), sorter: dateSorter },
    {
      title: 'Activity Type',
      dataIndex: ['activity_type', 'name'],
      sorter: activityTypeSorter,
    },
    { ...toCol('Duration', minutesToHMS), sorter: durationSorter },
    {
      ...toCol('Distance', (distance, record) =>
        distance > 0 ? `${distance} ${record.unit}` : '-'
      ),
      sorter: distanceSorter,
    },
    {
      title: 'Heart Rate',
      dataIndex: 'heart_rate',
      sorter: heartRateSorter,
    },
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

  const closeModal = () => setEditModalVisible(false);

  return (
    <div>
      <Search
        className="search-bar"
        allowClear
        placeholder="Search by name, type, or date..."
        onSearch={onSearchHandler}
        onChange={onChangeHandler}
        width="50%"
      />
      <Table rowKey="id" dataSource={filterActivities(searchTerm, activities)} columns={columns} />
      <Modal
        title="Edit Activity"
        visible={editModalVisible}
        onOk={closeModal}
        onCancel={closeModal}
        width={750}
        footer={null}
        destroyOnClose
      >
        <EditActivity refetch={refetch} initialActivity={editingActivity} modalClose={closeModal} />
      </Modal>
    </div>
  );
}
