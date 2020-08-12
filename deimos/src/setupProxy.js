const { createProxyMiddleware } = require('http-proxy-middleware');

const BACKEND_URL = process.env.REACT_APP_API_URL;

module.exports = (app) => {
  app.use(createProxyMiddleware('/auth', { target: BACKEND_URL }));
  app.use(createProxyMiddleware('/private/', { target: BACKEND_URL }));
  app.use(createProxyMiddleware('/metadata/', { target: BACKEND_URL }));
  app.use(createProxyMiddleware('/users/', { target: BACKEND_URL }));
  app.use(createProxyMiddleware('/strava', { target: BACKEND_URL }));
};
