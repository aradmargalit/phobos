import { filterActivities, toCol } from './tableUtils';

describe('tableUtils', () => {
  describe('toCol', () => {
    describe('dataIndex', () => {
      const scenarios = [
        { input: 'Arad Margalit', out: 'arad_margalit' },
        { input: '', out: '' },
        { input: null, out: '' },
      ];
      scenarios.forEach(({ input, out }) => {
        it('snake cases the display name for the data index', () => {
          expect(toCol(input).dataIndex).toEqual(out);
        });
      });
    });

    describe('renderMethod', () => {
      const dummyFunc = () => 'Hello World...';
      it('preserves injected render method', () => {
        expect(toCol('', dummyFunc).render).toEqual(dummyFunc);
      });

      it('has an undefined render method if nothing is passed in', () => {
        expect(toCol('').render).toEqual(undefined);
      });
    });
  });

  describe('filterActivities', () => {
    // A simplified representation for the sake of the tests
    const activities = [
      {
        name: 'Run 1',
        activity_type: { name: 'Run' },
        activity_date: '2020-01-02',
      },
      {
        name: 'Running with Jan',
        activity_type: { name: 'Run' },
        activity_date: '2006-01-02',
      },
      {
        name: 'Splishy Splashy',
        activity_type: { name: 'Swim' },
        activity_date: '2020-02-03',
      },
    ];

    describe('when search term is null or empty', () => {
      it('returns activities as-is', () => {
        expect(filterActivities('', activities)).toEqual(activities);
        expect(filterActivities(null, activities)).toEqual(activities);
        expect(filterActivities(undefined, activities)).toEqual(activities);
      });
    });

    describe('when search term matches the activity names', () => {
      it('filters activities by name', () => {
        expect(filterActivities('run', activities)).toHaveLength(2);
        expect(filterActivities('1', activities)).toHaveLength(1);
        expect(filterActivities('jan', activities)).toHaveLength(2);
      });
    });

    describe('when search term matches the activity type', () => {
      it('filters activities by activity type name', () => {
        expect(filterActivities('swim', activities)).toHaveLength(1);
        expect(filterActivities('SwiM', activities)).toHaveLength(1);
      });
    });

    describe('when search term matches some portion of the date', () => {
      it('filters activities by activity date', () => {
        expect(filterActivities('Jan', activities)).toHaveLength(2);
        expect(filterActivities('March', activities)).toHaveLength(0);
        expect(filterActivities('2006', activities)).toHaveLength(1);
      });
    });

    describe('when search term matches multiple filter criteria', () => {
      it('sorts by name => type => date', () => {
        expect(filterActivities('Jan', activities)[0].name).toEqual('Running with Jan');
        expect(filterActivities('Jan', activities)[1].name).toEqual('Run 1');
      });
    });
  });
});
