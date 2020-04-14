import { shallow } from 'enzyme';
import React from 'react';

import MileageGraph from './MileageGraph';

const generateComponent = props => {
  const defaultProps = {};

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<MileageGraph {...mergedProps} />);
};

describe('<MileageGraph />', () => {
  it('matches snapshots', () => {
    expect(
      generateComponent({
        monthlyData: [{ month: 'January 2020', miles: 123 }],
      })
    ).toMatchSnapshot();
  });
});
