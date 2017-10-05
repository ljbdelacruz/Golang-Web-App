angular.module('ngStarterKit')
.controller('MainAppCtrl',
            ['$scope',
             function ($scope) {
             $scope.controller={};
             $scope.controller.list=[{label:'Hello'}, {label:'World'}];

}]);