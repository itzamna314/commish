import Ember from 'ember';

export default Ember.Component.extend({
  init() {
    this._super(...arguments);
    this.set('filteredLeagues', this.get('leagues'));
  }
});
