import Ember from 'ember';

export default Ember.Controller.extend({
  actions: {
    selected(league) {
      this.set('selectedLeague', league);
    },
    createLeague() {
      this.set('selectedLeague', this.store.createRecord('league'));
    },
    createTeam(teamName) {
      this.store.createRecord('team', {name: teamName}).save().then(
        (team) => {
          this.get('selectedLeague.teams').pushObject(team);
        });
    },
    cancel() {
      if (this.get('selectedLeague.isNew')) {
        this.get('selectedLeague').deleteRecord();
      }
      this.set('selectedLeague', null);
      this.get('model.leagues').forEach( (p) => {
        p.set('isSelected', false);
      });
    },
    submit() {
      this.get('selectedLeague').save().then(
        () => {
          this.get('model.leagues').update();
        }
      );
      this.get('model.leagues').forEach( (p) => {
        p.set('isSelected', false);
      });
      this.set('selectedLeague', null);
    }
  }
});
