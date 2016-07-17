import _ from 'lodash/lodash';
import Ember from 'ember';

export function startCase(params/*, hash*/) {
  if (!params || params.length !== 1 ) {
    return `Illeagl start-case usage: Expected 1 param but got ${params.length}`;
  }

  return _.startCase(params[0]);
}

export default Ember.Helper.helper(startCase);
