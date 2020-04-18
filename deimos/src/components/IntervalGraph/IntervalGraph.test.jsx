import { shallow } from 'enzyme';
import React from 'react';

import IntervalGraph from './IntervalGraph';

const generateComponent = props => {
  const defaultProps = {
    data: [{ month: 'January 2020', duration: 123 }],
    projection: { x: 1, y: 2 },
    average: 6,
  };

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<IntervalGraph {...mergedProps} />);
};

describe('<IntervalGraph />', () => {
  it('renders', () => {
    generateComponent();
  });
});
