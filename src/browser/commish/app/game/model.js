import Model from 'ember-data/model';
import attr from 'ember-data/attr';
import { belongsTo /*, hasMany*/ } from 'ember-data/relationships';

export default Model.extend({
  match: belongsTo('match'),
  homeScore: attr(),
  awayScore: attr(),
  state: attr()
});
