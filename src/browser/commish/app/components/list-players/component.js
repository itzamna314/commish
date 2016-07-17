import Ember from 'ember';
import ResourceList from 'commish/components/resource-list/component';

export default ResourceList.extend({
  rows: null,
  init() {
    this._super(...arguments);
    this.set('rows', this.get('players'));
  },
  actions: {
    selected (player) {
      this.get('players').forEach( (p) => {
        p.set('isSelected', false);
      });
      player.set('isSelected', true);
      this.get('selected')(player);
    },
    filter(filter) {
      let filtered = this.get('players').filter( (p) => {
        return filter(p);
      });
      this.set('rows', filtered); 
    }
  }
});
