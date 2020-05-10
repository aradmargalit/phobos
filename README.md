# Phobos

<img src="./docs/phobos.png" alt="phoebe" width="200" />

Phobos is a fitness tracker app which functions like a fitness diary. The idea is to track the following on a daily basis:

- Workout Type
- Workout Duration
- Distance Covered

and to provide insights to the end-user:

- Pace
- Workout statistics
- Mileage over time
- Duration over time

Users can also enable automatic updates from [Strava](http://strava.com).

---

<div align=center><b>NOTE: </b>This is a personal project and is not representative of my quality of work in a professional context. This project allowed me to try many new technologies, and as a result, does not adhere to best practices.</div>

---

## Getting Started :rocket:

### Cloning the Repository

### Docker :whale:

Everything works using [Docker Compose](https://docs.docker.com/compose/), though it's not the best developer experience. However, if you want to see what the app looks like, you can follow these instructions.

You can set up a `.env` file to store these secrets at the project root. Docker compose will automatically pick up the `.env` file's variables.

Check out [the sample file](./.env.sample) to see the layout and format for this file and ask a contributor to share the Google and Strava credentials for you (or make your own!)

```sh
# --build rebuilds images in the event that the source has changed
docker-compose up --build
```

### Run Locally :computer:

I find it's best to run these in order!

#### MySQL DB :inbox_tray:

You should just use Docker for this. It's way easier.

```sh
docker-compose up -d mysql
```

#### Go Server :mailbox_with_no_mail:

Create a `.env` file under the `server` directory and fill out these variables:

```
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
API_DB_STRING=
COOKIE_SECRET_TOKEN=
FRONTEND_URL=http://localhost:3000
SERVER_URL=http://localhost:8080
STRAVA_CLIENT_ID=
STRAVA_CLIENT_SECRET=
STRAVA_WEBHOOK_SUB_ID=
```

You'll need Go Version >= 1.14

```sh
cd server
make run
```

Once the server is stood up, seed the database!

```sh
curl http://localhost:8080/admin/seed
```

#### Makefile

Check out the [Makefile](./server/makefile) for some common commands to help with development.

#### Client :moon:

To start deimos:

```sh
cd deimos
yarn install
yarn start
```

## Testing :white_check_mark:

There are unit tests for both the frontend and backend projects.

### Deimos Testing

Run `yarn test`. If you want a coverage report, run `yarn test --coverage`.

### Server Testing

Run `go test -v ./...`, or if you have [Richgo](https://github.com/kyoh86/richgo) installed, run `richgo test -v ./...` for colorized output.

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
- [Next Steps :footprints:](./docs/next.md)

## Etymology

[Phobos](<https://en.wikipedia.org/wiki/Phobos_(moon)>) is the larger of the two moons orbiting Mars :rocket:. Deimos is the smaller of the two moons.

I had to Google all that, so I'll admit that in reality, the app is named after my cat, Phoebe, who I affectionately call "Phobo". :cat2:
