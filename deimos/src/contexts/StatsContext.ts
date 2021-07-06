/* eslint-disable @typescript-eslint/no-empty-function */
import { createContext } from 'react';
import { FetchedData, Stats } from '../types';

export const initialStatsState: FetchedData<Stats> = {
  payload: null,
  loading: true,
};

interface StatsState {
  stats: FetchedData<Stats>;
  setStats: (stats: FetchedData<Stats>) => void;
}

const StatsContext = createContext<StatsState>({
  stats: initialStatsState,
  setStats: () => {},
});

export default StatsContext;
