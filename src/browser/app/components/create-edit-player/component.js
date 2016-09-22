import Ember from 'ember';

export default Ember.Component.extend({
  classNames: ["create-edit-player"],
  formMode: 'player',
  filteredTeams: Ember.computed('teams.@each', 'selectedPlayer.teams.@each', 'teamNameFilter', function() {
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
    createPlayer() {
      this.get('onCreatePlayer')();
    },
    createTeam() {
      this.get('onCreateTeam')(this.get('teamNameFilter'));
    },
    cancelPlayer() {
      this.set('teamNameFilter', null);
      this.get('onCancelPlayer')();
    },
    submitPlayer() { 
      this.set('teamNameFilter', null);
      this.get('onSubmitPlayer')();
    },
    showTeams() {
      this.set('formMode', 'teams');
    },
    showPlayer() {
      this.set('formMode', 'player');
    },
    addToTeam(team) {
      this.get('selectedPlayer.teams').pushObject(team);
    },
    removeFromTeam(team) {
      this.get('selectedPlayer.teams').removeObject(team);
    }
  }
});
