import Ember from 'ember';

export default Ember.Controller.extend({
  actions: {
    onUpdate(filter) {
      var model = this.get('model');
      this.set('model.players', this.store.query('players', { name: filter.get('text') }));
    },
    selected(player) {
      this.set('selectedPlayer', player);
    },
    create() {
      this.set('selectedPlayer', this.store.createRecord('players'));
    },
    submit() {
      this.get('selectedPlayer').save();
    }
  }
});
