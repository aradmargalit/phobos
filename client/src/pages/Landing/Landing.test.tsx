import React from 'react';
import { render } from '@testing-library/react';
import Landing from '.';

describe('Landing', () => {
  it('renders without crashing', () => {
    render(<Landing />);
  });
});
