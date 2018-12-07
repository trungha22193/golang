# Feature Overview
Provide CRUD Apps with SQLite3 

# GuideLine
- Step 1: get Source Code
- Step 2: setup GOPATH
- Step 3: execute below command below
	- $ go get github.com/labstack/echo
	- $ go get github.com/mattn/go-sqlite3
	- $ go build server.go
	- $ ./server


If you open another terminal window and list the contents of our app directory you should now see a file called "storage.db". Run the following to make sure the file is a valid SQLite file.
$ sqlite3 storage.db
(*) This command should bring up a prompt. At the prompt type ".tables" and then hit "Enter". You should see the "tasks" table listed. To exit the prompt type ".quit".



