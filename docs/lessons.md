# Lessons Learned

I wanted to keep a list of lessons I learned while building this projects, partially as a helpful guide to other developers at roughly my point in their careers, and partially as a reminder to myself for future work.

## Frontend :computer:

### Create React App

- When building the frontend container, there need to be 2 phases, a "build" phase (which runs `yarn build` or similar, and a "run" phase, which serves up the static bundle. In this project, I chose Nginx for that portion).

  Becuase `process.env.REACT_APP_X` is interpolated/replaced at _build_ time, you need to expose environment variables from `docker-compose.yml` during the "build phase".

- Now for the real plot twist - I wouldn't recommend serving the front end on its own. It's possible, but for small-scale applications, I find it easier to have the API serve up the client bundle.

- I hadn't used React Hooks before, and they're awesome! I'm a huge fan of writing components functionally and using hooks to manage state. :thumbsup:

## Backend :rocket:

- I learned to be careful with HTTP Status codes like `Permanent Redirect` or `Moved Permanently`. These are cached by the browser indefinitely and make recovering from a lazy typo in development very difficult. (Hint: clearing the browser cache (images and files) solves the issue.)

- [Gin](https://github.com/gin-gonic/gin) is awesome, but leaning on the community contributions was a lifesaver. That's how I was able to easily get session management and other tedious tasks.

- Getting a single database connection (or any other singleton) available to controller functions across files is easy if you define the database as part of the environment, and then make each controller function a method where the `Env` is the receiver. Nice little trick.

  :warning: I've since adjusted my view here a bit, and wish that I'd composed this differently. In retrospect, I'd have initialized a service which contains the data access objects, which connect to the DB as needed.

## Database :inbox_tray:

- MySQL uses `?` for parameterized queries. SQLite uses `$1`, `$2`, etc. This took me a little while to figure out and fix during database setup.

- Any scripts linked under `/docker-entrypoint-initdb.d/` will only run if the Docker volume has not yet been created. To re-run the scripts, you need to remove the volume with:

  ```sh
  docker-compose stop -v

  # OR if the containers are already stopped:

  docker-compose rm -v
  ```

- The database takes a minute to get ready to accept connections. While it's possible to build a healthcheck dependency into the docker compose file, it seemed way easier to build a static retry into the API.

## Infrastructure :whale:

- Docker Compose is an awesome way to keep everything linked together, and being able to always have a stable build with a single command is worth the headache.

- Any values in `.env` are automatically exported to the `docker-compose.yml` file, making it dead-simple to pass secrets.

- The database doesn't require its own `Dockerfile` since the official image provides a way to create root and API users, run init scripts, and define your own volumes. :books:

## Deployment

- I wanted to go AWS with EKS, but it was just way too expensive. Google Cloud provided a completely free, fully managed solution, and works well since I lean entirely on Google Auth for this project.
