import { formatDate, minutesToHMS } from './dataFormatUtils';

describe('minutesToHMS', () => {
  describe('with no minutes', () => {
    it('returns a "-"', () => {
      expect(minutesToHMS(0)).toEqual('-');
      expect(minutesToHMS(null)).toEqual('-');
      expect(minutesToHMS(undefined)).toEqual('-');
      expect(minutesToHMS('2005/12/12')).toEqual('-');
    });
  });

  describe("with fewer than a day's worth of minutes", () => {
    it('returns a correctly formatted string', () => {
      expect(minutesToHMS(5)).toEqual('00:05:00');
      expect(minutesToHMS(50)).toEqual('00:50:00');
      expect(minutesToHMS(60)).toEqual('01:00:00');
      expect(minutesToHMS(1234)).toEqual('20:34:00');
    });
  });

  describe('with fractional minutes', () => {
    it('returns a correctly formatted string', () => {
      expect(minutesToHMS(5.5)).toEqual('00:05:30');
      expect(minutesToHMS(120.75)).toEqual('02:00:45');
    });
  });

  describe('with a day or more worth of minutes', () => {
    it('returns a correctly formatted string', () => {
      expect(minutesToHMS(60 * 24)).toEqual('1d, 00:00:00');
      expect(minutesToHMS(10000)).toEqual('6d, 22:40:00');
    });
  });
});

describe('formatDate', () => {
  it('formats dates as MMMM, Do, YYYY', () => {
    expect(formatDate('2020-01-02')).toEqual('January 2nd, 2020');
    expect(formatDate('3000-07-12 12:12:11')).toEqual('July 12th, 3000');
  });
});
