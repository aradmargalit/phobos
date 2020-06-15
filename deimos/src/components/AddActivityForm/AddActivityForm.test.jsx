import { shallow } from 'enzyme';
import React from 'react';

import AddActivityForm from './AddActivityForm';

const generateComponent = props => {
  const defaultProps = {
    user: { given_name: 'Test' },
    activityTypes: [
      { id: 1, name: 'Running' },
      { id: 2, name: 'Biking' },
    ],
  };
  const mergedProps = { ...defaultProps, ...props };

  return shallow(<AddActivityForm {...mergedProps} />);
};

describe('<AddActivityForm />', () => {
  it('renders correctly', () => {
    expect(generateComponent()).toMatchSnapshot();
  });
});
