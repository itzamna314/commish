import Ember from 'ember';

export default Ember.Controller.extend({
  identity: Ember.inject.service(),
  username: null,
  password: null,
  actions: {
    login () {
      let id = this.get('identity');
      let user = this.get('username');
      let pass = this.get('password'); 
      id.loadIdentity(user, pass, true).then(
        () => { this.transitionToRoute('/manage'); },
        (msg) => { alert(msg); }
      );
    }
  }
});
