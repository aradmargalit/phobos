import { shallow } from 'enzyme';
import React from 'react';

import ActivityGraph from './ActivityGraph';

const generateComponent = props => {
  const defaultProps = {
    monthlyData: [{ month: 'January 2020', duration: 123 }],
  };

  const mergedProps = { ...defaultProps, ...props };

  return shallow(<ActivityGraph {...mergedProps} />);
};

describe('<ActivityGraph />', () => {
  it('renders without crashing', () => {
    generateComponent();
  });

  describe('monthly average', () => {
    it('calculates the correct average with one month', () => {
      const component = generateComponent();
      expect(component.find('ReferenceLine').prop('y')).toEqual(123 / 60);
    });

    it('calculates the correct average with multiple months', () => {
      const component = generateComponent({
        monthlyData: [
          { month: 'January 2020', duration: 60 },
          { month: 'February 2020', duration: 180 },
        ],
      });
      expect(component.find('ReferenceLine').prop('y')).toEqual(2);
    });
  });
});
