import Ember from 'ember';
import ENV from 'commish/config/environment';

export default Ember.Service.extend({
  username: null,
  token: null,
  getToken (username, password, demandNew) {
    this.username = username;

    if (this.token && !demandNew) {
      return new Ember.Promise((resolve /*, reject*/) => {
        resolve(this.get('token'));
      }); 
    }

    return new Ember.Promise( (resolve, reject) => {
      Ember.$.ajax(ENV.BaseUrl + '/admin/logins', {
        data: {
          identifier: username,
          password: password
        },
        dataType: 'json',
        success: (data, status) => {
          if ( status === '200' ) {
            this.set('token', data);
            resolve('foobar');
          } else {
            reject(`Received non-OK success from server: ${status}`);
          }
        },
        error: (data, status) => {
          reject(`Received non-success from server: ${status}`);
        }
      });
    });
  }
});
