import Ember from 'ember';

export default Ember.Component.extend({
  init() {
    this._super(...arguments);
  },
  actions: {
    rowSelected (row) {
      this.get('rows').forEach( (r) => {
        r.set('isSelected', false);
      });
      row.set('isSelected', true);
      if ( typeof(this.get('selected')) === 'function' ) {
        this.get('selected')(row);
      }
    }
  }
});
