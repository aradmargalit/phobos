/* eslint-disable camelcase */
import {
  activityTypeSorter,
  dateSorter,
  distanceSorter,
  durationSorter,
  nameSorter,
  numberSorter,
} from './sortUtils';

describe('sortUtils', () => {
  describe('numberSorter', () => {
    it('correctly sorts based on logical_indices', () => {
      const data = [{ logical_index: 2 }, { logical_index: 1 }, { logical_index: 3 }];
      expect(data.sort(numberSorter).map(({ logical_index }) => logical_index)).toStrictEqual([
        1, 2, 3,
      ]);
    });

    it('behaves when no data is passed in', () => {
      const data = [];
      expect(data.sort(numberSorter)).toStrictEqual([]);
    });
  });

  describe('nameSorter', () => {
    it('correctly sorts based on name', () => {
      const data = [{ name: 'testman' }, { name: 'whoahtest' }, { name: 'alf' }];
      expect(data.sort(nameSorter).map(({ name }) => name)).toStrictEqual([
        'alf',
        'testman',
        'whoahtest',
      ]);
    });
  });

  describe('dateSorter', () => {
    it('correctly sorts based on epoch', () => {
      const data = [{ epoch: 123 }, { epoch: 12345 }, { epoch: 1 }];
      expect(data.sort(dateSorter).map(({ epoch }) => epoch)).toStrictEqual([1, 123, 12345]);
    });
  });

  describe('activityTypeSorter', () => {
    it('correctly sorts based on activity type name', () => {
      const data = [
        { activity_type: { name: 'Swim' } },
        { activity_type: { name: 'Bike' } },
        { activity_type: { name: 'Run' } },
      ];
      expect(data.sort(activityTypeSorter).map((x) => x.activity_type.name)).toStrictEqual([
        'Bike',
        'Run',
        'Swim',
      ]);
    });
  });

  describe('durationSorter', () => {
    it('correctly sorts based on duration', () => {
      const data = [{ duration: 12.76 }, { duration: 1 }, { duration: 0 }];
      expect(data.sort(durationSorter).map(({ duration }) => duration)).toStrictEqual([
        0, 1, 12.76,
      ]);
    });
  });

  describe('distance', () => {
    it('correctly sorts based on distance', () => {
      const data = [{ meters: 12.76 }, { meters: 1 }, { meters: 0 }];
      expect(data.sort(distanceSorter).map(({ meters }) => meters)).toStrictEqual([0, 1, 12.76]);
    });
  });
});
