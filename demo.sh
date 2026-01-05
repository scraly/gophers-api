#!/bin/bash

########################
# include the magic
########################
. bin/demo-magic.sh

# hide the evidence
clear

# Run the API (in background to not have to open anew term?)
task run &

p "List existing Gophers"
pe "curl localhost:8080/gophers | jq ."

p "Get a Gopher"
pe "curl -s "localhost:8080/gopher?name=5th-element" | jq ."

p "Add a new Gopher"
pe "curl -X POST localhost:8080/gopher \
   -H \"Content-Type: application/json\" \
   -d '{\"name\":\"yoda-gopher\",\"displayname\":\"Yoda Gopher\",\"url\":\"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png\"}' "

p "Update a Gopher"
pe "curl -s -X PUT localhost:8080/gopher \
   -H \"Content-Type: application/json\" \
   -d '{\"name\":\"yoda-gopher\",\"displayname\":\"El mejor Yoda Gopher\",\"url\":\"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png\"}' | jq . "

p "Delete a Gopher"
pe "curl -X DELETE \"localhost:8080/gopher?name=5th-element\""

p "✅"

# Kill the Gophers API app running
lsof -n -i4TCP:8080 | awk '{print $2}' | tail -1  | xargs kill -9