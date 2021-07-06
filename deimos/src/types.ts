/**
 * Global types used throughout the application
 */

// Utility Types & Aliases
type ID = number;

export interface FetchedData<T> {
  payload: T | null;
  loading: boolean;
  errors?: string[];
}

type DayOfWeekCount = {
  dayOfWeek: 'Monday' | 'Tuesday' | 'Wednesday' | 'Thursday' | 'Friday' | 'Saturday' | 'Sunday';
  count: number;
};

type WorkoutTypePortion = {
  workoutType: string;
  portion: number;
};

export interface User {
  createdAt: string;
  email: string;
  givenName: string;
  id: ID;
  name: string;
  hasStravaToken: boolean;
  updatedAt: string;
}

export interface Stats {
  workouts: number;
  hours: number;
  miles: number;
  dayBreakdown: DayOfWeekCount[];
  typeBreakdown: WorkoutTypePortion[];
}
