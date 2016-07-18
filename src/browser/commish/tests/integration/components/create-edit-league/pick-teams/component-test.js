import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';

moduleForComponent('create-edit-league/pick-teams', 'Integration | Component | create edit league/pick teams', {
  integration: true
});

test('it renders', function(assert) {
  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{create-edit-league/pick-teams}}`);

  assert.equal(this.$().text().trim(), '');

  // Template block usage:
  this.render(hbs`
    {{#create-edit-league/pick-teams}}
      template block text
    {{/create-edit-league/pick-teams}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
