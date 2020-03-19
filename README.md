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
docker-compose up --build
```

Alternatively, you can set up a `.env` file to store these secrets at the project root. Docker compose will automatically pick up the `.env` file's variables.

Check out [the sample file](./.env.sample) to see the layout and format for this file.

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
API_DB_STRING= \
COOKIE_SECRET_TOKEN = \
FRONTEND_URL = \
SERVER_URL = \
./server
```

#### MySQL DB :inbox_tray:

You should just use Docker for this. It's way easier.

```sh
docker-compose up -d mysql
```

## Deployment :rocket:

This project is deployed on [Google Cloud Platform](https://cloud.google.com/), using [Cloud SQL](https://cloud.google.com/sql/) as a managed relational database solution and [Cloud Run](https://cloud.google.com/run) as a very simple managed container service.

### Setting Up the GCloud CLI

Follow [the instructions for your OS](https://cloud.google.com/sdk/docs/quickstart-macos) to get the `gcloud` CLI tool up and running.

This assumes you've already created a "project" in the cloud console that you can hook into. On running `gcloud init`, you'll login with your Google Account.

### Build and Deploy

To build the application, `cd` into the `server` directory and run:

```sh
# Format for the tag: gcr.io/[PROJECT NAME]/[IMAGE NAME]
docker build . --tag gcr.io/phobos-prod/phobos-server
```

Then, confirm you can push to gcr like so:

```sh
gcloud auth configure-docker gcr.io
```

Finally, push your built image to GCR.

```sh
docker push gcr.io/phobos-prod/phobos-server
```

To deploy, create a new revision in the Google Cloud Console and deploy. You can also do this from the CLI, if you'd prefer.

## More Documentation

While the above is all you need to get started, I'll keep more documents browseable [in the docs folder.](./docs)

### Table of Contents

- [Lessons Learned :mortar_board:](./docs/lessons.md)

## Etymology

[Phobos](<https://en.wikipedia.org/wiki/Phobos_(moon)>) is the larger of the two moons orbiting Mars :rocket:. Deimos is the smaller of the two moons.

I had to Google all that, so I'll admit that in reality, the app is named after my cat, Phoebe, who I affectionately call "Phobo". :cat2:
