import { shallow } from 'enzyme';
import React from 'react';

import ActivityGraph from './ActivityGraph';

const generateComponent = props => {
  const defaultProps = {};

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<ActivityGraph {...mergedProps} />);
};

describe('<ActivityGraph />', () => {
  it('renders without crashing', () => {
    generateComponent({
      monthlyData: [{ month: 'January 2020', duration: 123 }],
    });
  });
});
