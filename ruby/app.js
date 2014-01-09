angular.module('shortlyApp', [])
.config(function($routeProvider) {
  $routeProvider
  .when('/', {
    controller: 'HomeCtrl',
    template: "<h1>I'm in home controller</h1>"
    //templateUrl: "templates/home.html"
  })
  .when('/create', {
    controller: 'CreateCtrl',
    template: "<h1>I'm in home controller</h1>"
  })
  .otherwise({
    redirectTo: '/'
  });
})
.controller('CreateCtrl',
  function($scope) {

})
.controller('HomeCtrl',
  function($scope) {

    $scope.name = 'Ari';
});