# Ravelin
Tech challenge for Henrique Calixto Go+Javascript

## Building and running the server:
    - On this readMe I'm assuming you cloned this repo on this folder:
         $GOPATH/github.com/HCLXTO
      
      If you did so, please follow the rest of the instructions, if
      not, please remember to alter the path on line 8 of the 
      '/server/server.go' file to reflect your current folder structure 

     - Open the command line on this folder and run:

       ```go install monitor/monitor.go```
       ```cd server```
       ```go build server.go```
       ```./server```

       The server will be running on http://localhost:8080/

       To stop it press ```ctrl+c```