import Ember from 'ember';

export default Ember.Component.extend({
  classNames: ["create-edit-player"],
  formMode: 'player',
  actions: {
    createPlayer() {
      this.get('onCreatePlayer')();
    },
    cancelPlayer() {
      this.get('onCancelPlayer')();
    },
    submitPlayer() { 
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
