'use strict';

/* Create module for navbar directive */
angular.module('directives.loadingScreen1', [])
    .directive('loadingScreen1', ['$http',
        '$location', 'CustomTimout',
        function ($http, $location, CustomTimout) {
            function preFn(scope, element, attr) {
            }
            /* Do the directive's logic here */
            function postFn(scope, element, attr) {
                scope.isShow=true;
                scope.toggle=function(){
                    scope.isShow=!scope.isShow;
                }
            }
            return {
                restrict: 'E',
                transclude:true,
                scope: {
                    toggle:'=',
                },
                replace: true,
                templateUrl: 'Assets/angularjs/Components/Widgets/loadingScreen/loadingScreen.widget.html',
                compile: function(scope, element, attr) {
                    return {
                        pre: preFn,
                        post: postFn
                    }
                }
            };
        }
    ]);