import ResourceList from 'commish/components/resource-list/component';
import Ember from 'ember';

export default ResourceList.extend({
  filter: null,
  init() {
    this._super(...arguments);
  },
  rows: Ember.computed('leagues.[]', 'leagues.@each.isNew', 'filter', function() {
    let filterFn = this.get('filter');
    return this.get('leagues').filter( (p) => {
      return !p.get('isNew') && (!filterFn || filterFn(p));
    });
  }),
  actions: {
    rowSelected (league) {
      this.get('leagues').forEach( (l) => {
        l.set('isSelected', false);
      });
      league.set('isSelected', true);
      this.get('selected')(league);
    },
    filter(filter) {
      this.set('filter', filter);
    }
  }
});
