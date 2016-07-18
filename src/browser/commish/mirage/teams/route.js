import Mirage from 'ember-cli-mirage';
import TeamsFixture from './fixture';

export default  {
  route: function(router) {
    router.get('teams', (/*db*/) => {
      return {teams: TeamsFixture};
    });

    router.get('teams/:id', (db, request) => {
      return {
        matches: TeamsFixture.filter(function(elt) {
          return elt.publicId === request.params.id;
        })
      };
    });

    router.post('teams', (/*db, request*/) => {
      return new Mirage.Response(201);
    });
  }
};
