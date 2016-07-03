import JSONAPISerializer from 'ember-data/serializers/json-api';
import _ from 'lodash/lodash';
import Ember from 'ember';

export default JSONAPISerializer.extend({
  normalizeResponse (store, type, payload) {
    let data = _.reduce(payload, function(accum, resources, resourceName) {
      return accum.concat(_.map(resources, function(res) {
        return {
          id: res.publicId,
          type: resourceName,
          attributes: res
        };
      }));
    }, []);
 
    return {
      data: data     
    };
  },
  serialize (snapshot/*, options */) {
    let keys = Ember.get(snapshot, 'type.fields._keys').toArray();
    let values = _.map(keys, (k) => {
      return snapshot.attr(k);
    });

    let data = _.zipObject(keys, values);
    return data;
  }
});
