import Ember from 'ember';
import ResourceList from 'commish/components/resource-list/component';

export default ResourceList.extend({
  standings: Ember.inject.service(),
  enterResultState: 'closed',
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
  resultTeamsChanged: Ember.observer('enterResultHomeTeam', 'enterResultAwayTeam', function() {
    let home = this.get('enterResultHomeTeam');
    let away = this.get('enterResultAwayTeam');

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
    this.set('enterResultState', 'enter-scores');
  }),
  actions: {
    setEnterResultState (state) {
      this.set('enterResultState', state);

      if (state === 'closed') {
        // Clear state
      }
    },
    selectHomeTeam (team) {
      this.get('league.teams').forEach ( (t) => t.set('isHomeSelected', false));
      team.set('isHomeSelected', true);
      this.set('enterResultHomeTeam', team);
    },
    selectAwayTeam (team) {
      this.get('league.teams').forEach ( (t) => t.set('isAwaySelected', false));
      team.set('isAwaySelected', true);
      this.set('enterResultAwayTeam', team);
    }
  }
});
