import Ember from 'ember';

export default Ember.Component.extend({
  classNames: ["create-edit-league"],
  formMode: 'league',
  actions: {
    createLeague() {
      this.get('onCreateLeague')();
    },
    createTeam(teamName) {
      this.get('onCreateTeam')(teamName);
    },
    cancelLeague() {
      this.set('teamNameFilter', null);
      this.get('onCancelLeague')();
    },
    submitLeague() { 
      this.set('teamNameFilter', null);
      this.get('onSubmitLeague')();
    },
    switchMode(mode) {
      this.set('formMode', mode);
    },
    addToTeam(team) {
      this.get('selectedLeague.teams').pushObject(team);
    },
    removeFromTeam(team) {
      this.get('selectedLeague.teams').removeObject(team);
    }
  }
});
