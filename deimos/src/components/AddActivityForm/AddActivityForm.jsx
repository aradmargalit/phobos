import {
  Button, DatePicker, Select, InputNumber,
} from 'antd';
import React from 'react';
import { useForm, Controller } from 'react-hook-form';
import './AddActivityForm.scss';
import EmojiOption from '../EmojiOption';

export default function AddActivityForm() {
  const {
    handleSubmit, errors, control,
  } = useForm();
  const onSubmit = (data) => console.log(data);
  console.log(errors);


  return (
    <form
      className="ant-form ant-form-horizontal login-form"
      onSubmit={handleSubmit(onSubmit)}
    >

      <strong>Activity Date</strong>
      <Controller
        as={DatePicker}
        control={control}
        className="input"
        placeholder="2020-01-01"
        name="Activity Date"
        required
      />

      <strong>Activity Type</strong>
      <Controller
        as={Select}
        control={control}
        className="input"
        name="Activity Selector"
        showSearch
        style={{ width: 200 }}
        placeholder="Activity Type"
        optionFilterProp="children"
        required
      >
        { /* This should really be a HOC, but Ant gets angry if you try */ }
        {EmojiOption({ emoji: '🏃', value: 'run', title: 'Run' })}
        {EmojiOption({ emoji: '🚴', value: 'bike', title: 'Bike' })}
        {EmojiOption({ emoji: '🏊', value: 'swim', title: 'Swimming' })}
      </Controller>

      <strong>Duration</strong>
      <div className="input">
        <Controller
          as={InputNumber}
          min={1}
          placeholder={120}
          control={control}
          name="Duration"
          required
        />
        <i>  minutes</i>
      </div>
      <Button
        htmlType="submit"
        icon="rocket"
        type="primary"
      >
        Submit
      </Button>

    </form>
  );
}
