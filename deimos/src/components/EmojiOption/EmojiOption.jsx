import React from 'react';
import PropTypes from 'prop-types';
import { Select } from 'antd';

const { Option } = Select;

// This _should_ be a HOC, but Antd doesn't allow it, so here we are
export default function EmojiOption({ emoji, value, title }) {
  return (
    <Option value={value}>
      {title}
      <span className="emoji" role="img" aria-label={`${value} emoji`}>
        {emoji}
      </span>
    </Option>
  );
}

EmojiOption.propTypes = {
  emoji: PropTypes.string.isRequired,
  value: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
};
