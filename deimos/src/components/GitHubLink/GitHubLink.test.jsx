import { shallow } from 'enzyme';
import React from 'react';

import GitHubLink from './GitHubLink';

describe('<GitHubLink />', () => {
  it('renders without crashing like a good button should', () => {
    shallow(<GitHubLink />);
  });
});
