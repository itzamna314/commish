import Ember from 'ember';

export default Ember.Component.extend({
  init() {
    this._super(...arguments);
    this.set('filter', Ember.Object.create());
  },
  onUpdated: Ember.observer('filterText', (sender) => {
    let filterText = sender.get('filterText');
    let filterField = sender.get('filterField');
    
    let filterFunc = (o) => {
      let value = o.get(filterField);
      if (typeof(value) === 'number') {
        value = value.toString();
      }
      if (typeof(value) !== 'string') {
        return false;
      }

      return value.toLowerCase().indexOf(filterText.toLowerCase()) > -1;
    };

    if ( sender.get('onUpdate') ) {
      sender.get('onUpdate')(filterFunc);
    }
  })
});
