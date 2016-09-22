import JSONAPIAdapter from 'ember-data/adapters/json-api';
import Ember from 'ember'; 
import ENV from 'commish/config/environment';

export default JSONAPIAdapter.extend({
  identity: Ember.inject.service(),
  namespace: 'api',
  headers: Ember.computed('identity.token', function() {
    return {
      "X-COMMISH-CONNECTION": ENV.connectionId,
      "Authorization": `Bearer ${this.get('identity.token')}`
    };
  }),
  query (store, type, query) {
    var url = this.buildURL(type.modelName, null, null, 'query', query);

    if (this.sortQueryParams) {
      query = this.sortQueryParams(query);
    }

    return this.ajax(url + '/queries', 'POST', { data: query });
  }
});
