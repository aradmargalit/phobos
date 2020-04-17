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

  xdescribe('monthly average', () => {
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
