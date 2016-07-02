import JSONAPISerializer from 'ember-data/serializers/json-api';
import _ from 'lodash/lodash';

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
  }
});
