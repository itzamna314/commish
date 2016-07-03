import Ember from 'ember';

export default Ember.Controller.extend({
  actions: {
    selected(player) {
      this.set('selectedPlayer', player);
    },
    create() {
      this.set('selectedPlayer', this.store.createRecord('players'));
    },
    cancel() {
      if (this.get('selectedPlayer.isNew')) {
        this.get('selectedPlayer').deleteRecord();
      }
      this.set('selectedPlayer', null);
      this.get('model.players').forEach( (p) => {
        p.set('isSelected', false);
      });
    },
    submit() {
      this.get('selectedPlayer').save();
      this.set('selectedPlayer', null);
      this.get('model.players').forEach( (p) => {
        p.set('isSelected', false);
      });
    }
  }
});
