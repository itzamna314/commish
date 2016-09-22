import Ember from 'ember';
import ResourceList from 'commish/components/resource-list/component';

export default ResourceList.extend({
  standings: Ember.inject.service(),
  init() { 
    this._super(...arguments);
    this.get('standings').loadLeagueTable(this.get('league')).then(
      (table) => {
        this.set('leagueTable', table);
      }
    );
  },
  leagueChanged: Ember.observer('league.teams.[]', function() {
    this.get('standings').loadLeagueTable(this.get('league')).then(
      (table) => {
        this.set('leagueTable', table);
      }
    );
  }),
  actions: {
  }
});
