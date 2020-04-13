import { makeDurationBreakdown } from './durationUtils';

describe('makeDurationBreakdown', () => {
  it('preserves the total minutes', () => {
    expect(makeDurationBreakdown(23.234).total).toEqual(23.234);
  });

  it('correctly breaks down hours', () => {
    expect(makeDurationBreakdown(60).hours).toEqual(1);
    expect(makeDurationBreakdown(120).hours).toEqual(2);
    expect(makeDurationBreakdown(123).hours).toEqual(2);
  });

  it('correctly breaks down remaining minutes', () => {
    expect(makeDurationBreakdown(60).minutes).toEqual(0);
    expect(makeDurationBreakdown(12).minutes).toEqual(12);
    expect(makeDurationBreakdown(123).minutes).toEqual(3);
  });

  it('correctly calculates down remaining seconds', () => {
    expect(makeDurationBreakdown(60).seconds).toEqual(0);
    expect(makeDurationBreakdown(12).seconds).toEqual(0);
    expect(makeDurationBreakdown(1.5).seconds).toEqual(30);
  });
});
