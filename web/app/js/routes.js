module.exports = function(angularApp) {
  angularApp.config(function($stateProvider) {
    var secotrsDetail, sectors;
    sectors = {
      name: 'sectors',
      url: '/',
      component: 'sectorsView'
    };
    secotrsDetail = {
      name: 'sectorsDetail',
      url: "/{sector}",
      component: 'sectorView'
    };
    return $stateProvider.state([sectors, sectorsDetail]);
  });
  return angularApp;
};
