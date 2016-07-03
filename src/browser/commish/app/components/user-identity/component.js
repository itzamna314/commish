import Ember from 'ember';

export default Ember.Component.extend({
  identity: Ember.inject.service(),
  actions: {
    logIn () {
      this.get('router').transitionTo('login');
    }
  }
});
