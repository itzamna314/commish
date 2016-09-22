import Mirage from 'ember-cli-mirage';
import GamesFixture from './fixture';

export default {
  route: function(router) {
    router.get('games', (/*db*/) => {
      return {games: GamesFixture};
    });

    router.get('games/:id', (db, request) => {
      return {
        games: GamesFixture.filter(function(elt) {
          return elt.publicId === request.params.id;
        })
      };
    });

    router.post('games', (/*db, request*/) => {
      return new Mirage.Response(201);
    });

    router.patch('games/:id', (db, request) => { 
      var game = JSON.parse(request.requestBody);
      game.publicId = request.params.id;
      game.id = request.params.id;
      return {games: [game]}; 
    });

    router.post('games/queries', (db, request) => {
      var name = JSON.parse(request.requestBody).name;
      return {
        games: GamesFixture.filter(function(elt) { 
          return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
        })
      };
    });
  }
};
