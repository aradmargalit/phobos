import { FetchedData } from '../types';

// eslint-disable-next-line import/prefer-default-export
export function defaultState<T>(): FetchedData<T> {
  return {
    payload: null,
    loading: true,
  };
}
