import Mirage from 'ember-cli-mirage';
import Players from './players/route';
import Teams from './teams/route';

export default function() {

  this.namespace = 'api';    // make this `api`, for example, if your API is namespaced

  this.post('/admin/logins', (db, request) => {
    console.log(request.requestBody);
    var req = JSON.parse(request.requestBody);
    if ( req.identifier === 'mirage' && req.password === 'mirage' ) {
      return {
        user: {
          connection: 'foobar',
          identifier: 'mirage',
          token: 'Dummy-Mirage'
        }
      };
    } else {
      return new Mirage.Response(403, { "message": "Invalid username or password"});
    }
  });

  Players(this);
  Teams(this);
}

/*
  Shorthand cheatsheet:

  this.get('/posts');
  this.post('/posts');
  this.get('/posts/:id');
  this.put('/posts/:id'); // or this.patch
  this.del('/posts/:id');

  http://www.ember-cli-mirage.com/docs/v0.2.x/shorthands/
*/
