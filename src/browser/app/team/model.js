import Model from 'ember-data/model';
import attr from 'ember-data/attr';
import { /*belongsTo, */ hasMany } from 'ember-data/relationships';

export default Model.extend({
  name: attr(),
  players: hasMany('players'),
  leagues: hasMany('leagues'),
  homeMatches: hasMany('matches', {inverse: 'homeTeam'}),
  awayMatches: hasMany('matches', {inverse: 'awayTeam'}),
});
