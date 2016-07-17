import Ember from 'ember';

export default Ember.Route.extend({
  model() {
    return Ember.RSVP.hash({
      players: this.store.findAll('player'), 
      teams: this.store.findAll('team')
    });
  }
});
