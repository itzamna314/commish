import JSONAPISerializer from 'ember-data/serializers/json-api';
import _ from 'lodash/lodash';
import Ember from 'ember';

export default JSONAPISerializer.extend({
  inflector: new Ember.Inflector(Ember.Inflector.defaultRules),
  singularRequests: ['updateRecord', 'findRecord', 'createRecord'],
  normalizeResponse (store, type, payload, id, requestType) {
    let data = _.reduce(payload, (accum, resources, resourceName) => {
      return accum.concat(_.map(resources, (res) => {
        let attrs = {};
        let rels = {};

        type.eachAttribute( (name/*, meta*/) => {
          attrs[name] = res[name];
        });

        type.eachRelationship( (name, meta) => {
          let rel = null;

          if ( meta.kind === 'hasMany' ) {
            if (!res[name]) { 
              rel = []; 
            } else {
              rel = {
                "data": _.map(res[name], (d) => {
                  return {
                    type: meta.type,
                    id: d
                  };
                })
              };
            }
          } else if (meta.kind === 'belongsTo') {
            if (!res[name]) {
              rel = null;
            } else {
              rel = {
                "data": {
                  type: meta.type,
                  id: res[name]
                }
              }; 
            }
          }

          rels[name] = rel;
        });

        return {
          id: res.publicId,
          type: this.inflector.singularize(resourceName),
          attributes: attrs,
          relationships: rels
        };
      }));
    }, []);

    if (this.singularRequests.some( (s) => s === requestType) ) {
      data = data[0];
    }

    return {
      data: data     
    };
  },
  serialize (snapshot/*, options */) {
    let data = Ember.get(snapshot, '_attributes');
    snapshot.eachRelationship( (name, rel) => {
      if (rel.kind === 'hasMany') {
        data[this.inflector.pluralize(rel.type)] = snapshot.hasMany(name, {ids: true});
      } else if (rel.kind === 'belongsTo') {
        data[this.inflector.singularize(rel.type)] = snapshot.belongsTo(name, {id: true});
      }
    });
    return data;
  }
});
