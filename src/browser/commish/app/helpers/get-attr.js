import Ember from 'ember';

export function getAttr(params/*, hash*/) {
  if (!params || params.length !== 2 ) {
    return `Invalid get-attr usage: Expected 2 params, got ${params.length}`;
  }

  if (typeof(params[0]) !== 'object') { 
    return `Invalid get-attr usage: First param is not an object`;
  }

  if (typeof(params[1]) !== 'number' && typeof(params[1]) !== 'string') {
    return `Invalid get-attr usage: Second param is not a valid object key`;
  }

  if ( typeof(params[0].get) === 'function' ) { 
    return params[0].get(params[1]);
  }

  return params[0][params[1]];
}

export default Ember.Helper.helper(getAttr);
