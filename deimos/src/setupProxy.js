const proxy = require('http-proxy-middleware');
const BACKEND_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';

module.exports = function(app) {
  console.log(BACKEND_URL);
  app.use(proxy('/auth/google', { target: BACKEND_URL }));
  app.use(proxy('/currentUser', { target: BACKEND_URL }));
};
