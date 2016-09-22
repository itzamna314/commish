import Ember from 'ember';
import ResourceList from 'commish/components/resource-list/component';

export default ResourceList.extend({
  filter: null,
  init() {
    this._super(...arguments);
  },
  rows: Ember.computed('players.[]', 'players.@each.isNew', 'filter', function() {
    let filterFn = this.get('filter');
    return this.get('players').filter( (p) => {
      return !p.get('isNew') && (!filterFn || filterFn(p));
    });
  }),
  actions: {
    selected (player) {
      this.get('players').forEach( (p) => {
        p.set('isSelected', false);
      });
      player.set('isSelected', true);
      this.get('selected')(player);
    },
    filter(filter) {
      this.set('filter', filter);
    }
  }
});
