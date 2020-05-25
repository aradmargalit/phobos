import { shallow } from 'enzyme';
import React from 'react';

import Trendline from './Trendline';

const generateComponent = props => {
  const defaultProps = {};

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<Trendline {...mergedProps} />);
};

describe('<Trendline />', () => {
  it('renders correctly', () => {
    // TODO: Mock the API call to do something useful with this
    generateComponent();
  });
});
