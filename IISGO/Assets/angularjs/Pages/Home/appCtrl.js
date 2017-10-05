angular.module('ngStarterKit')
.controller('MainAppCtrl',
            ['$scope','$http','$location',
             function ($scope, $http, $location) {
                 $scope.navigate=function(path){
                    $location.path(path);
                 }

            }]);