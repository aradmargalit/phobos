import React from 'react';
import PropTypes from 'prop-types';
import { Select } from 'antd';

const { Option } = Select;

const nameEmojiMap = {
  run: '🏃',
  bike: '🚴',
  swim: '🏊',
};

// This _should_ be a HOC, but Antd doesn't allow it, so here we are
export default function EmojiOption({ value, title }) {
  const emoji = nameEmojiMap[title.toLowerCase()];

  return (
    <Option key={value} value={value}>
      {title}
      { emoji
      && (
      <span className="emoji" role="img" aria-label={`${value} emoji`}>
        {emoji}
      </span>
      )}
    </Option>
  );
}

EmojiOption.propTypes = {
  value: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
};
