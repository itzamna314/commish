// import Mirage from 'ember-cli-mirage';
import LeaguesFixture from './fixture';

export default {
  route: function(router) {
    router.get('leagues', (/*db*/) => {
      return {leagues: LeaguesFixture};
    });

    router.post('leagues', (db, request) => {
      return {leagues: LeaguesFixture.push(request) };
    });

    router.patch('leagues/:id', (db, request) => { 
      var league = JSON.parse(request.requestBody);
      league.publicId = request.params.id;
      league.id = request.params.id;
      return {leagues: [league]}; 
    });

    router.post('leagues/queries', (db, request) => {
      var name = JSON.parse(request.requestBody).name;
      return {
        leagues: LeaguesFixture.filter(function(elt) { 
          return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
        })
      };
    });
  }
};
