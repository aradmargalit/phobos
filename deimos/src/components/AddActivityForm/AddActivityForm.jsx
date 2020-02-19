import {
  Row, Col, DatePicker, Select, InputNumber,
} from 'antd';
import React from 'react';
import { useForm } from 'react-hook-form';
import './AddActivityForm.scss';

export default function AddActivityForm() {
  const { register, handleSubmit, errors } = useForm();
  const onSubmit = (data) => console.log(data);
  console.log(errors);

  const { Option } = Select;

  return (
    <form
      className="ant-form ant-form-horizontal login-form"
      onSubmit={handleSubmit(onSubmit)}
    >
      <Row gutter={4}>
        <Col span={6}>
          <strong>Activity Date</strong>
        </Col>
        <Col span={18}>
          <DatePicker
            placeholder="2020-01-01"
            name="Activity Date"
            ref={register({ required: true })}
          />
        </Col>
      </Row>
      <Row gutter={4}>
        <Col span={6}>
          <strong>Activity Type</strong>
        </Col>
        <Col span={18}>
          <Select
            showSearch
            style={{ width: 200 }}
            placeholder="Activity Type"
            optionFilterProp="children"
            // onChange={onChange}
            // onFocus={onFocus}
            // onBlur={onBlur}
            // onSearch={onSearch}
            // filterOption={(input, option) =>
            //   option.props.children.toLowerCase().indexOf(input.toLowerCase()) >=
            //   0
            // }
            ref={register}
          >
            <Option value="run">
              Run
              <span role="img" aria-label="running emoji">
                🏃
              </span>
            </Option>
            <Option value="bike">
              Bike
              <span role="img" aria-label="biking emoji">
                🚴
              </span>
            </Option>
            <Option value="swim">
              Swim
              <span role="img" aria-label="swimming emoji">
                🏊
              </span>
            </Option>
          </Select>
        </Col>
      </Row>
      <Row gutter={4}>
        <Col span={6}>
          <strong>Duration</strong>
        </Col>
        <Col span={18}>
          <div>
            <InputNumber min={1} max={10} defaultValue={3} ref={register} />
            <i>  minutes</i>
          </div>
        </Col>
      </Row>
    </form>
  );
}
