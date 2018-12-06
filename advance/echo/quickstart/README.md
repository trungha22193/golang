## Feature Overview
- Build basic CRUD without DBs storing. 

## Guiline to setup
- Setup Golang by https://golang.org/dl/
- Setup GOPATH by https://github.com/golang/go/wiki/SettingGOPATH
	(*)export GOPATH=$HOME/golang
		export PATH=$GOPATH/bin:$PATH
- Run command line go get Framrework
	# cd [project_folder] && go get -u github.com/labstack/echo/...
- Execute Apps as below
	# cd ~/golang/advance/echo/quickstart && go run server.go


## API list
	- http://localhost:1234/
	- GET detail http://localhost:1234/users/123?team=zenith
	- POST http://localhost:1234/users 
	(*) use x-www-form-urllendcoded
	(*) Or execute #curl -F "name=Joe Smith" http://localhost:1234/users
	- PUT http://localhost:1234/users/123
	- DELETE http://localhost:1234/users/123
