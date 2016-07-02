export default function() {

  // These comments are here to help you get started. Feel free to delete them.

  /*
    Config (with defaults).

    Note: these only affect routes defined *after* them!
  */

  // this.urlPrefix = '';    // make this `http://localhost:8080`, for example, if your API is on a different server
  this.namespace = 'api';    // make this `api`, for example, if your API is namespaced
  // this.timing = 400;      // delay for each request, automatically set to 0 during testing

  /*
    Shorthand cheatsheet:

    this.get('/posts');
    this.post('/posts');
    this.get('/posts/:id');
    this.put('/posts/:id'); // or this.patch
    this.del('/posts/:id');

    http://www.ember-cli-mirage.com/docs/v0.2.x/shorthands/
  */

  this.playersFixture = [
    {
      "publicId":"E438746A260F11E6875C81C96BBD740C",
      "name":"Geraldine",
      "age":37,
      "gender":"female"
    },
    {
      "publicId":"02735D8C261011E6875C81C96BBD740C",
      "name":"Gerald",
      "age":42,
      "gender":"male"
    },
    {
      "publicId":"331E1DFA261011E6875C81C96BBD740C",
      "name":"Gerald",
      "age":42,
      "gender":"male"
    },
    {
      "publicId":"3E26CB0C261011E6875C81C96BBD740C",
      "name":"Gerald",
      "age":42,
      "gender":"male"
    },
    {
      "publicId":"62D8F744300C11E6A09D409B4CAB7549",
      "name":"Geraldine",
      "age":37,
      "gender":"female"
    }
  ];

  this.get('players', (/*db*/) => {
    return {players: this.playersFixture};
  });

  this.post('players');
  this.patch('players/:id', (db, request) => { 
    var player = JSON.parse(request.requestBody);
    return {players: [player]}; 
  });

  this.post('players/query', (db, request) => {
    var name = JSON.parse(request.requestBody).name;
    return {
      players: this.playersFixture.filter(function(elt) { 
        return !name || elt.name.toLowerCase().indexOf(name.toLowerCase()) > -1;
      })
      };
    }
  );
}
