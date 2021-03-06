## Feature Overview
- Provide CRUD Apps with SQLite3 

### Technical advices
- The <b>`_ "github.com/mattn/go-sqlite3"`</b> has to be top of <b>"server.go"</b>. Vice versa, show <b>ERROR</b>
- Must use "&" ~ Pointer, to fetch data like  `rows.Scan(&task.ID, &task.Name)` in `models > tasks.go`
- Must use "&" ~ Pointer, to map data like `c.Bind(&task)` in `handlers > tasks.go`

## GuideLine
- Step 1: get Source Code
- Step 2: setup GOPATH
	Example:
	```
	export GOPATH=$HOME/golang
	export PATH=$GOPATH/bin:$PATH
	```
- Step 3: execute below command below
	- $ go get github.com/labstack/echo
	- $ go get github.com/mattn/go-sqlite3
	- $ go build server.go
	- $ ./server


If you open another terminal window and list the contents of our app directory you should now see a file called "storage.db". Run the following to make sure the file is a valid SQLite file.
```
$ sqlite3 storage.db
```
(*) <i>This command should bring up a prompt. At the prompt type ".tables" and then hit "Enter". You should see the "tasks" table listed. To exit the prompt type ".quit".</i>



## API list
- GET list http://localhost:1234/tasks
- PUT http://localhost:1234/users 
```
{
    "name": "Zenith"
}
```
- DELETE http://localhost:1234/tasks/1

## User Interface
- $ ./server
- Click on http://localhost:1234/ to see App as below

![Screenshot](api_demo.png)