import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('manage', function() {
    this.route('players');
  });
  this.route('login');
});

export default Router;
