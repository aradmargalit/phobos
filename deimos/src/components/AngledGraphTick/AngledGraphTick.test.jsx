import { shallow } from 'enzyme';
import React from 'react';

import AngledGraphTick from './AngledGraphTick';

const generateComponent = (props) => {
  const defaultProps = {
    x: 1,
    y: 2,
    payload: { value: 'December 2018' },
  };
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<AngledGraphTick {...mergedProps} />);
};

describe('<AngledGraphTick />', () => {
  it('renders correctly', () => {
    expect(generateComponent()).toMatchSnapshot();
  });
});
