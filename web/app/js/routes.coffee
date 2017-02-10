module.exports = (angularApp) ->
  angularApp.config ($stateProvider) ->

    sectors =
      name: 'sectors'
      url: '/'
      component: 'sectorsView'

    secotrsDetail =
      name: 'sectorsDetail'
      url: "/{sector}"
      component: 'sectorView'

    $stateProvider.state [sectors,sectorsDetail]

  angularApp
