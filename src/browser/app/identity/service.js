import Ember from 'ember';

export default Ember.Service.extend({
  store: Ember.inject.service(),
  username: null,
  token: null,
  init () {
    let username = window.localStorage.getItem('username');
    let token = window.localStorage.getItem('token');
    if ( username && token ) {
      this.set('username', username);
      this.set('token', token);
    }
  },
  logout() {
    window.localStorage.removeItem('username');
    window.localStorage.removeItem('token');
    this.set('username', null);
    this.set('token', null);
  },
  loadIdentity (username, password, demandNew) {
    this.set('username', username);

    if (this.token && !demandNew) {
      return new Ember.RSVP.Promise((resolve /*, reject*/) => {
        resolve(this.get('token'));
      }); 
    }

    return new Ember.RSVP.Promise( (resolve, reject) => {
      let adapter = Ember.getOwner(this).lookup("adapter:application");
      let host = adapter.get('host') || window.location.host;
      let namespace = adapter.get('namespace');
      Ember.$.ajax(`http://${host}/${namespace}/admin/logins`, {
        contentType: 'application/json',
        data: JSON.stringify({
          identifier: username,
          password: password
        }),
        dataType: 'json',
        method: 'POST',
        success: (data, status) => {
          if ( data && data.user && data.user.token ) {
            this.set('token', data.user.token);
            window.localStorage.setItem('username', username);
            window.localStorage.setItem('token', this.get('token'));
            resolve(this.get('token'));
          } else {
            reject(`Did not receive a token: ${status}`);
          }
        },
        error: (data, status) => {
          reject(`Received non-success from server: ${status}`);
        }
      });
    });
  }
});
