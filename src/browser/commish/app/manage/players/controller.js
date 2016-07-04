import Ember from 'ember';

export default Ember.Controller.extend({
  actions: {
    selected(player) {
      this.set('selectedPlayer', player);
      this.get('selectedPlayer.teams');
    },
    create() {
      this.set('selectedPlayer', this.store.createRecord('player'));
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
