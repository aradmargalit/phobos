import { shallow } from 'enzyme';
import React from 'react';

import DurationPicker from './DurationPicker';

const generateComponent = props => {
  const defaultProps = {
    value: { hours: 1, minutes: 2, seconds: 3, onChange: null },
  };
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<DurationPicker {...mergedProps} />);
};

describe('<DurationPicker />', () => {
  it('renders correctly', () => {
    expect(generateComponent()).toMatchSnapshot();
  });
});
