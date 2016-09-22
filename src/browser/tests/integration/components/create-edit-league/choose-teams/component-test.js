import { moduleForComponent, test } from 'ember-qunit';
import hbs from 'htmlbars-inline-precompile';

moduleForComponent('create-edit-league/choose-teams', 'Integration | Component | create edit league/choose teams', {
  integration: true
});

test('it renders', function(assert) {
  // Set any properties with this.set('myProperty', 'value');
  // Handle any actions with this.on('myAction', function(val) { ... });

  this.render(hbs`{{create-edit-league/choose-teams}}`);

  assert.equal(this.$().text().trim(), '');

  // Template block usage:
  this.render(hbs`
    {{#create-edit-league/choose-teams}}
      template block text
    {{/create-edit-league/choose-teams}}
  `);

  assert.equal(this.$().text().trim(), 'template block text');
});
