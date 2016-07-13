import Mirage from 'ember-cli-mirage';
import PlayersFixture from './fixture';

export default {
  route: function(router) {
    router.get('players', (/*db*/) => {
      return {players: PlayersFixture};
    });

    router.post('players', (/*db, request*/) => {
      return new Mirage.Response(201);
    });

    router.patch('players/:id', (db, request) => { 
      var player = JSON.parse(request.requestBody);
      player.publicId = request.params.id;
      player.id = request.params.id;
      return {players: [player]}; 
    });

    router.post('players/queries', (db, request) => {
      var name = JSON.parse(request.requestBody).name;
      return {
        players: PlayersFixture.filter(function(elt) { 
          return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
        })
      };
    });
  }
};
