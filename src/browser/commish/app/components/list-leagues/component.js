import ResourceList from 'commish/components/resource-list/component';

export default ResourceList.extend({
  rows: null,
  init() {
    this._super(...arguments);
    this.set('rows', this.get('leagues'));
  },
  actions: {
    rowSelected (league) {
      this.get('leagues').forEach( (l) => {
        l.set('isSelected', false);
      });
      league.set('isSelected', true);
      this.get('selected')(league);
    },
    filter(filter) {
      let filtered = this.get('leagues').filter( (p) => {
        return filter(p);
      });
      this.set('rows', filtered); 
    }
  }
});
