import { defaultState } from './stateUtils';

describe('defaultState', () => {
  it('should always return the correct state', () => {
    expect(defaultState()).toEqual({
      payload: null,
      loading: true,
      errors: false,
    });
  });
});
