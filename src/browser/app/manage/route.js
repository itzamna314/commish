import Ember from 'ember';

export default Ember.Route.extend({
  identity: Ember.inject.service(),
  beforeModel () {
    if ( !this.get('identity.token') ) {
      this.transitionTo('/login');
    }
  },
  
});
