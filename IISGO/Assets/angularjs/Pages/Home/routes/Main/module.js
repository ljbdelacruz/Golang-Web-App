angular.module('modules.Main', [])
.config(['$routeProvider',
         function ($routeProvider) {
             /* Handle route on location path is '/' */
             $routeProvider.when('/', {
                 templateUrl: '/Assets/angularjs/Pages/Home/routes/Main/Main.page.html',
                 controller: 'MainCtrl'
             });
         }
]);