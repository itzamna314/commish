import Ember from 'ember';

export default Ember.Component.extend({
  state: 'pick-teams',
  resultTeamsChanged: Ember.observer('homeTeam', 'awayTeam', function() {
    let home = this.get('homeTeam');
    let away = this.get('awayTeam');

    if (!home || !away) {
      return;
    }

    if (home.get('id') === away.get('id')) {
      home.set('isError', true);
      away.set('isError', true);
      return;
    }

    home.set('isError', false);
    away.set('isError', false);
    this.set('state', 'enter-scores');
  }),
  actions: {
    setEnterResultState (state) {
      this.set('state', state);

      // Clear state
      if (state === 'pick-teams') {
        this.get('league.teams').forEach( (t) => {
          t.set('isHomeSelected', false);
          t.set('isAwaySelected', false);
        });
        this.set('homeTeam', null);
        this.set('awayTeam', null);
      }
    },
    selectHomeTeam (team) {
      this.get('league.teams').forEach ( (t) => t.set('isHomeSelected', false));
      team.set('isHomeSelected', true);
      this.set('homeTeam', team);
    },
    selectAwayTeam (team) {
      this.get('league.teams').forEach ( (t) => t.set('isAwaySelected', false));
      team.set('isAwaySelected', true);
      this.set('awayTeam', team);
    }
  }
});
