const proxy = require('http-proxy-middleware');

const BACKEND_URL = process.env.REACT_APP_API_URL;

module.exports = app => {
  app.use(proxy('/auth', { target: BACKEND_URL }));
  app.use(proxy('/private/', { target: BACKEND_URL }));
  app.use(proxy('/metadata/', { target: BACKEND_URL }));
  app.use(proxy('/users/', { target: BACKEND_URL }));
};
