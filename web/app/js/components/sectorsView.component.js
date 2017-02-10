module.exports = function(angularApp) {
  var HeroDetailController;
  HeroDetailController = function() {
    var ctrl;
    ctrl = this;
    ctrl["delete"] = function() {
      return ctrl.onDelete({
        hero: ctrl.hero
      });
    };
    return ctrl.update = function(prop, value) {
      return ctrl.onUpdate({
        hero: ctrl.hero,
        prop: prop,
        value: value
      });
    };
  };
  angularApp.module('heroApp'.component('heroDetail', {
    templateUrl: 'heroDetail.html',
    controller: HeroDetailController,
    bindings: {
      hero: '<',
      onDelete: '&',
      onUpdate: '&'
    }
  }));
  return angularApp;
};
