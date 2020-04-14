import { shallow } from 'enzyme';
import React from 'react';

import SkippedGraph from './SkippedGraph';

const generateComponent = props => {
  const defaultProps = {};

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<SkippedGraph {...mergedProps} />);
};

describe('<SkippedGraph />', () => {
  it('matches snapshots', () => {
    expect(
      generateComponent({
        monthlyData: [{ month: 'January 2020', days_skipped: 123 }],
      })
    ).toMatchSnapshot();
  });
});
