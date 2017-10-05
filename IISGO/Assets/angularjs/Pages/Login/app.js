'use strict';
//this app is for login page
angular.element(document)
    .ready(function () {
        angular.bootstrap(document, ['ngStarterKit']);
    });
angular.module('ngStarterKit', [
            'ngRoute',
            /*Widgets*/
            'directives.imageUploader1',
            /*Services*/
            'otherApp'
])
.config(['$routeProvider',
                 function ($routeProvider) {
                     $routeProvider.otherwise({ redirectTo: '/' });
                 }
]);