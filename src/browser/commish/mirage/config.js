import PlayersFixture from './players/fixture';
import Mirage from 'ember-cli-mirage';

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

  this.get('players', (/*db*/) => {
    return {players: PlayersFixture};
  });

  this.post('players');
  this.patch('players/:id', (db, request) => { 
    var player = JSON.parse(request.requestBody);
    return {players: [player]}; 
  });

  this.post('players/query', (db, request) => {
    var name = JSON.parse(request.requestBody).name;
    return {
      players: PlayersFixture.filter(function(elt) { 
        return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
      })
      };
    }
  );
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
