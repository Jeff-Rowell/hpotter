'use strict';

module.exports = function (environment) {
  const ENV = {
    modulePrefix: 'hpotter-ui',
    environment,
    rootURL: '/',
    locationType: 'history',
    EmberENV: {
      EXTEND_PROTOTYPES: false,
      FEATURES: {},
    },

    APP: {},
  };

  if (environment === 'production') {
    ENV.locationType = 'hash';
  }

  return ENV;
};
