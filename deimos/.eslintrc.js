module.exports = {
  env: {
    browser: true,
    es6: true,
    jest: true,
  },
  extends: ['plugin:react/recommended', 'airbnb', 'plugin:prettier/recommended'],
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
  },
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
    ecmaVersion: 2018,
    sourceType: 'module',
  },
  plugins: ['react', 'simple-import-sort', 'prettier'],
  rules: {
    'react/prop-types': 0,
    'react/jsx-props-no-spreading': 0,
    'simple-import-sort/imports': 'error',
    'max-len': ['error', { code: 100 }],
    'prettier/prettier': 'error',
    'react/jsx-wrap-multilines': ['error', { declaration: false, assignment: false }],
  },
};
