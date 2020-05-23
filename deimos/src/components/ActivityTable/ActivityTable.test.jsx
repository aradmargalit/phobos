import { shallow } from 'enzyme';
import React from 'react';

import ActivityTable from './ActivityTable';

const generateComponent = props => {
  const defaultProps = {};
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<ActivityTable {...mergedProps} />);
};

describe('<ActivityTable />', () => {
  it('renders correctly', () => {
    expect(generateComponent());
  });

  describe('while loading', () => {
    const component = generateComponent({ loading: true });
    it('has a spinner and no data', () => {
      expect(component.find('Spin')).toHaveLength(1);
      expect(component.find('Empty')).toHaveLength(0);
    });
  });

  describe('with no activities...yet', () => {
    it('renders an Empty component', () => {
      expect(generateComponent().find('Empty')).toHaveLength(1);
    });
  });

  describe('with activities', () => {
    it('renders a table with no pagination', () => {
      const component = generateComponent({
        activities: [
          {
            strava_id: 1,
            logical_index: 1,
            name: 'Fun Run',
            activity_date: new Date('2020-01-01'),
            activity_type: { name: 'Run' },
            duration: 12,
            distance: 40,
          },
        ],
      });
      const table = component.find('Table');
      expect(table).toHaveLength(1);
      expect(table.prop('dataSource')).toHaveLength(1);
      expect(table.prop('columns')).toHaveLength(10);
    });
  });
});
