# Phobos

![Phobos](./docs/phobos.png)

Phobos is a work-in-progress fitness tracker app meant to replace a long-running spreadsheet. The idea is to track the following on a daily basis:

- Workout Type
- Workout Duration
- Miles Traveled

and to calculate a series of derived fields:

- Relative Effort
- Pace
- Calories Burned

The end-goal is allow for automated data import from [Strava](http://strava.com).

## Getting Started :rocket:

The goal is to string everything together using [Docker Compose](https://docs.docker.com/compose/). You should just need to:

```sh
# --build rebuilds images in the event that the source has changed
docker-compose up --build
```