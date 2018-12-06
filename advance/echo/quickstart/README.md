# Feature Overview
Provide baisc CRUD without BDs storing.

# Guiline to setup
- Setup Golang by https://golang.org/dl/
- Setup GOPATH by https://github.com/golang/go/wiki/SettingGOPATH
	Example:
	```
	export GOPATH=$HOME/golang
	export PATH=$GOPATH/bin:$PATH
	```
- Run command line to get Echo Framrework
	### cd [project_folder] && go get -u github.com/labstack/echo/...
- Execute Apps as below
	### cd ~/golang/advance/echo/quickstart && go run server.go


# API list
- GEt listhttp://localhost:1234/
- GET detail http://localhost:1234/users/123?team=zenith
- POST http://localhost:1234/users 
#### (*) Using x-www-form-urllendcoded or curl -F "name=Joe Smith" http://localhost:1234/users
- PUT http://localhost:1234/users/123
- DELETE http://localhost:1234/users/123
	
	
