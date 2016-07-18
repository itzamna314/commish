import Model from 'ember-data/model';
import attr from 'ember-data/attr';
import { belongsTo, hasMany } from 'ember-data/relationships';

export default Model.extend({
  homeTeam: belongsTo('team', {inverse: 'homeMatches'}),
  awayTeam: belongsTo('team', {inverse: 'awayMatches'}),
  league: belongsTo('league'),
  games: hasMany('games'),
  state: attr()
});
