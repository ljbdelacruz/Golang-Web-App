angular.module('ngStarterKit')
.controller('MainAppCtrl',
            ['$scope','$http',
             function ($scope, $http) {
             $scope.controller={};
             $scope.controller.list=[{label:'Hello'}, {label:'World'}];
             $scope.controller.model={username:'', password:''};
             
             $scope.Login=function(){
                $http.post("/CheckAuthentication/", {username:""+$scope.controller.model.username, password:""+$scope.controller.model.password}).success(function (data, status, headers, config) {
                        console.log(data);
                }).error(function (data, status, headers, config) {
                        console.log(data);
                });
                // $http.post("/Home/Insert", {ID:"1ee", FirstName:"Lainel John", MiddleName:"B", LastName:"Dela Cruz", BirthDate:"12/05/1995"}).success(function (data, status, headers, config) {
                //         console.log(data);
                // }).error(function (data, status, headers, config) {
                //         console.log(data);
                // });
             }
             $scope.imageController={};
             $scope.imageController.model={file:""};
             $scope.UploadImage=function(){
                var fd = new FormData();
                fd.append('file', $scope.imageController.model.file);
                $http.post("/Home/UploadImage", fd, {
                        transformRequest: angular.identity,
                        headers: {'Content-Type': 'multipart/form-data'}
                    }).success(function(){
                    })
                    .error(function(){
                    });
                // $http.post("/Home/UploadImage", {Source:""+$scope.imageController.model.file}).success(function (data, status, headers, config) {
                //         console.log(data);
                // }).error(function (data, status, headers, config) {
                //         console.log(data);
                // });
             }

}]);