import Mirage from 'ember-cli-mirage';
import MatchesFixture from './fixture';

export default {
  route: function(router) {
    router.get('matches', (/*db*/) => {
      return {matches: MatchesFixture};
    });

    router.get('matches/:id', (db, request) => {
      return {
        matches: MatchesFixture.filter(function(elt) {
          return elt.publicId === request.params.id;
        })
      };
    });

    router.post('matches', (/*db, request*/) => {
      return new Mirage.Response(201);
    });

    router.patch('matches/:id', (db, request) => { 
      var match = JSON.parse(request.requestBody);
      match.publicId = request.params.id;
      match.id = request.params.id;
      return {matches: [match]}; 
    });

    router.post('matches/queries', (db, request) => {
      var name = JSON.parse(request.requestBody).name;
      return {
        matches: MatchesFixture.filter(function(elt) { 
          return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
        })
      };
    });
  }
};
