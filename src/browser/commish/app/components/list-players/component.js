import Ember from 'ember';

export default Ember.Component.extend({
  filteredPlayers: null,
  init() {
    this._super(...arguments);
    this.set('filteredPlayers', this.get('players'));
  },
  actions: {
    selected (player) {
      this.get('players').forEach( (p) => {
        p.set('isSelected', false);
      });
      player.set('isSelected', true);
      this.get('selected')(player);
    },
    onUpdate(filter) {
      let filtered = this.get('players').filter( (p) => {
        return p.get('name').toLowerCase().indexOf(filter.get('text').toLowerCase()) > -1;
      });
      this.set('filteredPlayers', filtered); 
    }
  }
});
