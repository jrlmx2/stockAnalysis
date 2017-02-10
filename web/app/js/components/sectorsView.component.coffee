module.exports = (angularApp) ->
  HeroDetailController = () ->
    ctrl = this;

    ctrl.delete = () ->
      ctrl.onDelete hero: ctrl.hero

    ctrl.update = (prop, value) ->
      ctrl.onUpdate
        hero: ctrl.hero
        prop: prop
        value: value


  angularApp.module 'heroApp' .component 'heroDetail',
    templateUrl: 'heroDetail.html'
    controller: HeroDetailController
    bindings:
      hero: '<'
      onDelete: '&'
      onUpdate: '&'

  angularApp
