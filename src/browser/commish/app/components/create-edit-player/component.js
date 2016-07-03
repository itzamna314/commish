import Ember from 'ember';

export default Ember.Component.extend({
  classNames: ["create-edit-player"],
  actions: {
    createPlayer() {
      this.get('onCreatePlayer')();
    },
    cancelPlayer() {
      this.get('onCancelPlayer')();
    },
    submitPlayer() { 
      this.get('onSubmitPlayer')();
    }
  }
});
