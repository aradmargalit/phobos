/* eslint-disable @typescript-eslint/no-empty-function */
import { createContext } from 'react';
import { FetchedData, User } from '../types';

interface UserState {
  user: FetchedData<User>;
  setUser: (user: FetchedData<User>) => void;
}

export const initialUserState: FetchedData<User> = {
  payload: null,
  loading: true,
};

const UserContext = createContext<UserState>({
  user: initialUserState,
  setUser: () => {},
});

export default UserContext;
