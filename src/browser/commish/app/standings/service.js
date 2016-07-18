import Ember from 'ember';
let Promise = Ember.RSVP.Promise;

export default Ember.Service.extend({
  loadLeagueTable (league) {
    return new Promise( (resolve, reject) => {
      league.get('teams').then( 
        (teams) => {
          Promise.all(teams.map( (t) => this.loadMatches(t, league))).then(
            (rows) => {
              let leagueTable = Ember.A();
              rows.forEach( (r) => {
                let wins = 0, losses = 0;

                r.matches.forEach( (m) => {
                  let winner = m.get('winner');
                  if ( winner ) {
                    if ( winner.get('id') === r.team.get('id')) {
                      wins++;
                    } else {
                      losses++;
                    }
                  }
                });

                leagueTable.addObject({
                  team: r.team,
                  wins: wins,
                  losses: losses,
                  gamesPlayed: r.matches.length
                });
              });
              resolve(leagueTable.sort( (l,r) =>  r.wins - l.wins));
            }
          );
        },
        (reason) => { reject(reason); }
      );
    });
  },
  loadMatches(team, league) {
    return new Promise( (resolve, reject) => {
      Promise.all([
        team.get('homeMatches'),
        team.get('awayMatches'),
        Promise.resolve(team),
        Promise.resolve(league)
      ]).then( 
        (values) => {
          let allMatches = Ember.A();
          let league = values[3];
          allMatches.addObjects(values[0].filter(m => m.get('league.id') === league.get('id')));
          allMatches.addObjects(values[1].filter(m => m.get('league.id') === league.get('id')));
          Promise.all(allMatches.map(m => this.matchWinner(m))).then( 
            (matches) => {
              resolve({
                team: values[2],
                matches: matches
              });
            },
            (reason) => reject(reason)
          );
        },
        (reason) => reject(reason)
      );
    });
  },
  matchWinner(match) {
    return new Promise( (resolve, reject) => {
      if ( match.get('winner') !== undefined ) {
        resolve(match);
        return;
      }

      match.get('games').then(
        (games) => {
          let awayWins = 0, homeWins = 0;
          games.forEach( (g) => {
            if ( g.get('homeScore') > g.get('awayScore') ) {
              homeWins++;
            } else if ( g.get('awayScore') > g.get('homeScore') ) {
              awayWins++;
            }
          });

          let winner = null;
          if ( awayWins > homeWins ) {
            winner = match.get('awayTeam');
          } else if ( homeWins > awayWins ) {
            winner = match.get('homeTeam');
          }

          match.set('winner', winner);
          resolve(match);
        },
        (reason) => reject(reason)
      );
    });
  }
});
