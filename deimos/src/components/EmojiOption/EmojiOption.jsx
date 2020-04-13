import { Select } from 'antd';
import PropTypes from 'prop-types';
import React from 'react';

const { Option } = Select;

const nameEmojiMap = {
  ride: '🚴',
  run: '🏃',
  bike: '🚴',
  swim: '🏊',
  walk: '🚶',
  hike: '🥾',
  'alpine ski': '⛷️',
  'backcountry ski': '⛷️',
  canoe: '🛶',
  crossfit: '🏋️',
  'e-bike ride': '🚴',
  elliptical: '🚶',
  handcycle: '✋',
  'ice skate': '⛸️',
  'inline skate': '🎢',
  kayak: '🚣',
  'kitesurf session': '🪁',
  'nordic ski': '⛷️',
  'rock climb': '🧗',
  'roller ski': '🙄',
  row: '🚣',
  snowboard: '🏂',
  snowshoe: '🎿',
  'stair stepper': '📶',
  'stand up paddle': '🌊',
  surf: '🏄',
  'virtual ride': '🚲',
  'virtual run': '👟',
  'weight training': '🏋️',
  'windsurf session': '⛵',
  wheelchair: '🦽',
  workout: '🏅',
  yoga: '🧘',
  basketball: '⛹️‍♀️',
  soccer: '⚽',
  ultimate: '🥏',
  tennis: '🎾',
  volleyball: '🏐',
};

// This _should_ be a HOC, but Antd doesn't allow it, so here we are
export default function EmojiOption({ value, title }) {
  const emoji = nameEmojiMap[title.toLowerCase()];

  return (
    <Option key={value} value={value}>
      {title}
      {emoji && (
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
