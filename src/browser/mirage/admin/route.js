import Mirage from 'ember-cli-mirage';
import LoginsFixture from './fixture';

export default {
  route: function(router) {
    router.post('admin/logins', (db, request) => {
      return new Mirage.Response(20);
    });
  }
};
