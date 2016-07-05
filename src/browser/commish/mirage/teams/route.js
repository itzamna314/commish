import Mirage from 'ember-cli-mirage';
import TeamsFixture from './fixture';

export default function(router) {
  router.get('teams', (/*db*/) => {
    return {teams: TeamsFixture};
  });

  router.post('teams', (/*db, request*/) => {
   return new Mirage.Response(201);
  });
}
