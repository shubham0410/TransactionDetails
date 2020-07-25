To setUpGolang:
    sudo apt-get update
    sudo apt-get -y upgrade
    curl -O https://storage.googleapis.com/golang/go1.13.4.linux-amd64.tar.gz
    tar -xvf go1.13.4.linux-amd64.tar.gz
    sudo mv go /usr/local
    sudo nano ~/.profile
    ###### add the below lines #####
    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
    source ~/.profile
    reboot #restart the system to reflect

To Install Dependencies:
    go get github.com/astaxie/beego/orm
    go get github.com/go-sql-driver/mysql
    go get github.com/gorilla/mux
    go get github.com/astaxie/beego/orm
    go get github.com/spf13/cast

To Run the project:
    Run the queries from queries.sql into mysql

    change the mysql user name and password in main.go.
        mysqlConf := "{username}:{password}@(localhost:3306)/test?interpolateParams=true"
    
    run command: go run main.go
        server will run on 7000 port, url is:http://localhost:7000
    
    open index.html (have basic Frontend page to call all the APIs)

    Postman Collection DownloadLinK:
        https://www.getpostman.com/collections/29ac81a1d0c968749527


ScreenShot of the FE is TransationDetails.png