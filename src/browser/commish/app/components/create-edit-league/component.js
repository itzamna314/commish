import Ember from 'ember';

export default Ember.Component.extend({
  classNames: ["create-edit-league"],
  formMode: 'league',
  filteredTeams: Ember.computed('teams.@each', 'selectedLeague.teams.@each', 'teamNameFilter', function() {
    let selectedLeague = this.get('selectedLeague');
    return this.get('teams').filter( (item) => {
        if ( !selectedLeague ) { return true; }

        return !selectedLeague.get('teams').any( (t) => {
          return t.get('id') === item.get('id');
        } );
    }).filter( (item) => {
      let nameFilter = this.get('teamNameFilter');
      if ( !nameFilter ) { return true; }
      return item.get('name').toLowerCase().indexOf(nameFilter.toLowerCase()) > -1;
    });
  }),
  actions: {
    createLeague() {
      this.get('onCreateLeague')();
    },
    createTeam() {
      this.get('onCreateTeam')(this.get('teamNameFilter'));
    },
    cancelLeague() {
      this.set('teamNameFilter', null);
      this.get('onCancelLeague')();
    },
    submitLeague() { 
      this.set('teamNameFilter', null);
      this.get('onSubmitLeague')();
    },
    showTeams() {
      this.set('formMode', 'teams');
    },
    showLeague() {
      this.set('formMode', 'league');
    },
    addToTeam(team) {
      this.get('selectedLeague.teams').pushObject(team);
    },
    removeFromTeam(team) {
      this.get('selectedLeague.teams').removeObject(team);
    }
  }
});
