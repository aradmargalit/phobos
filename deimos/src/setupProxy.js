const proxy = require('http-proxy-middleware');

const BACKEND_URL = process.env.REACT_APP_API_URL;

module.exports = app => {
  app.use(proxy('/auth/google', { target: BACKEND_URL }));
  app.use(proxy('/private/users/current', { target: BACKEND_URL }));
};
