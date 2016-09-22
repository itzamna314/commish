import Ember from 'ember';

export default Ember.Component.extend({
  identity: Ember.inject.service(),
  userName: Ember.computed('identity.username', function(){
    return this.get('identity.username');
  }),
  actions: {
    logIn () {
      this.get('router').transitionTo('login');
    },
    logOut() {
      this.get('identity').logout();
      this.get('router').transitionTo('login');
    }
  }
});
