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

### Cloning the Repository

Given that this project uses [Go](https://golang.org/), I recommend cloning this project into your `$GOPATH/src/`. [You can learn more about the `GOPATH` if you don't understand why](https://github.com/golang/go/wiki/GOPATH).

### Docker :whale:
The goal is to string everything together using [Docker Compose](https://docs.docker.com/compose/). You should just need to:

```sh
# --build rebuilds images in the event that the source has changed
# You'll need to pass environemnt variables for the server
GOOGLE_CLIENT_ID= \
GOOGLE_CLIENT_SECRET= \
docker-compose up --build
```

Alternatively, you can set up a `server/.env` file to store these secrets. Docker compose will automatically pick up the `.env` file's variables.

```text
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
```

### Run Locally :computer:

#### Client :moon:

To start deimos:

```sh
yarn install
yarn start
```

#### Go Server :mailbox_with_no_mail:

You'll need Go Version 1.13

```sh
go build

GOOGLE_CLIENT_ID= \
GOOGLE_CLIENT_SECRET= \
./server
```

## Etymology

[Phobos](https://en.wikipedia.org/wiki/Phobos_(moon)) is the larger of the two moons orbiting Mars :rocket:. Deimos is the smaller of the two moons.

I had to Google all that, so I'll admit that in reality, the app is named after my cat, Phoebe, who I affectionately call "Phobo". :cat2:
