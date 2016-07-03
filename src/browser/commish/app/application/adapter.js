import JSONAPIAdapter from 'ember-data/adapters/json-api';
import ENV from 'commish/config/environment';

var extension = {
  namespace: 'api',
  headers: {
    "X-COMMISH-CONNECTION": ENV.connectionId
  },
  query (store, type, query) {
    var url = this.buildURL(type.modelName, null, null, 'query', query);

    if (this.sortQueryParams) {
      query = this.sortQueryParams(query);
    }

    return this.ajax(url + '/query', 'POST', { data: query });
  }
};

export default JSONAPIAdapter.extend(extension);
