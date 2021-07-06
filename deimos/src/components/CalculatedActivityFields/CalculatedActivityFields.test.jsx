import { shallow } from 'enzyme';
import React from 'react';

import CalculatedActivityFields from './CalculatedActivityFields';

const generateComponent = (props) => {
  const defaultProps = {
    activity: {
      duration: { total: 100 },
      distance: 5,
      unit: 'miles',
    },
  };
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<CalculatedActivityFields {...mergedProps} />);
};

describe('<CalculatedActivityFields />', () => {
  it('renders correctly', () => {
    expect(generateComponent()).toMatchSnapshot();
  });

  it('correctly singularizes unit', () => {
    expect(generateComponent().prop('title')).toEqual('min / mile');
  });

  it('correctly calculates pace', () => {
    expect(generateComponent().prop('value')).toEqual(20);
  });
});
