import Ember from 'ember';

export default Ember.Component.extend({
  filteredTeams: Ember.computed('teams.[]', 'league.teams.[]', 'teamNameFilter', function() {
    let selectedPlayer = this.get('selectedPlayer');
    return this.get('teams').filter( (item) => {
        if ( !selectedPlayer ) { return true; }

        return !selectedPlayer.get('teams').any( (t) => {
          return t.get('id') === item.get('id');
        } );
    }).filter( (item) => {
      let nameFilter = this.get('teamNameFilter');
      if ( !nameFilter ) { return true; }
      return item.get('name').toLowerCase().indexOf(nameFilter.toLowerCase()) > -1;
    });
  }),
  actions: {
    createTeam() {
      this.get('onCreateTeam')(this.get('teamNameFilter'));
    },
    addToLeague(team) {
      this.get('league.teams').pushObject(team);
    },
    removeFromLeague(team) {
      this.get('league.teams').removeObject(team);
    }
  }
});
