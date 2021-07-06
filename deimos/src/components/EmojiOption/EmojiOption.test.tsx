import { Select } from 'antd';
import { shallow } from 'enzyme';
import React from 'react';

import EmojiOption from './EmojiOption';

const { Option } = Select;

const generateComponent = (props?) => {
  const defaultProps = {
    value: 'run',
    title: 'run',
  };

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<EmojiOption {...mergedProps} />);
};

describe('<EmojiOption />', () => {
  it('renders without crashing', () => {
    const component = generateComponent();
    expect(component.find(Option)).toHaveLength(1);
    expect(component.find('span')).toHaveLength(1);
  });

  it('renders when no emoji matches', () => {
    generateComponent({
      title: 'handball',
      value: 'handball',
    });
  });
});
