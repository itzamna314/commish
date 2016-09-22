/* jshint node: true */

module.exports = function(environment) {
  var ENV = {
    modulePrefix: 'commish',
    environment: environment,
    baseURL: '/',
    locationType: 'auto',
    connectionId: '0B132F40EA6E11E58650365B26ADF0AB',
    EmberENV: {
      FEATURES: {
        // Here you can enable experimental features on an ember canary build
        // e.g. 'with-controller': true
      }
    },

    APP: {
      // Here you can pass flags/options to your application instance
      // when it is created
    },

    EXTEND_PROTOTYPES: {
      Date: false
    }
  };

  if (environment == 'local') {
    ENV['ember-cli-mirage'] = {
      enabled: true
    };
    ENV.connectionId = 'F089772C424E11E69C56973441A785A3';
  }

  if (environment === 'development') {
    // ENV.APP.LOG_RESOLVER = true;
    // ENV.APP.LOG_ACTIVE_GENERATION = true;
    // ENV.APP.LOG_TRANSITIONS = true;
    // ENV.APP.LOG_TRANSITIONS_INTERNAL = true;
    // ENV.APP.LOG_VIEW_LOOKUPS = true;
    ENV['ember-cli-mirage'] = {
        enabled: false
    };
  }

  if (environment === 'test') {
    // Testem prefers this...
    ENV.baseURL = '/';
    ENV.locationType = 'none';

    // keep test console output quieter
    ENV.APP.LOG_ACTIVE_GENERATION = false;
    ENV.APP.LOG_VIEW_LOOKUPS = false;

    ENV.APP.rootElement = '#ember-testing';
  }

  if (environment === 'production') {

  }

  return ENV;
};