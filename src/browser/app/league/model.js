import Model from 'ember-data/model';
import attr from 'ember-data/attr';
import { /*belongsTo, */ hasMany } from 'ember-data/relationships';

export default Model.extend({
  name: attr(),
  location: attr(),
  description: attr(),
  division: attr(),
  gender: attr(),
  startDate: attr(),
  endDate: attr(),
  teams: hasMany('teams')
});
