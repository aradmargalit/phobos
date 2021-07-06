import { shallow } from 'enzyme';
import React from 'react';

import DOWBarChart from './DOWBarChart';

const generateComponent = (props) => {
  const defaultProps = {
    dayBreakdown: [{ day_of_week: 'Monday', count: 10 }],
  };
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<DOWBarChart {...mergedProps} />);
};

describe('<DOWBarChart />', () => {
  it('renders correctly', () => {
    expect(generateComponent()).toMatchSnapshot();
  });
});
