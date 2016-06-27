import Ember from 'ember';

export default Ember.Component.extend({
  init() {
    this._super(...arguments);
    this.set('filter', Ember.Object.create());
  },
  onUpdated: Ember.observer('filterText', (sender) => {
    sender.set('filter.text', sender.get('filterText'));
    if ( sender.get('onUpdate') ) {
      sender.get('onUpdate')(sender.get('filter'));
    }
  })
});
