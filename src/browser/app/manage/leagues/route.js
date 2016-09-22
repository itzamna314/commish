import Ember from 'ember';

export default Ember.Route.extend({
  model() {
    return Ember.RSVP.hash({
      headers: ['name', 'location', 'division', 'gender', 'startDate', 'endDate'],
      leagues: this.store.findAll('league'),
      teams: this.store.findAll('team')
    });
  }
});
