# essentially a pushbullet clone, I wanted to build an opensource self hosted version
# essentially the simplest version, with a database, encryption at rest, and file upload



# https://lets-go.alexedwards.net/sample/02.08-html-templating-and-inheritance.html

everything lives in the docker container, using it similarly to WSL because, I am lazy. 

docker container logs <container_id> -f to continously see the logs, air is always running in the background so it will build code automatically. can wathc the logs. 

any go mods etc just bash into the container 

docker exec -it <container_id> bash

## helpers
### https://lets-go.alexedwards.net/sample/02.08-html-templating-and-inheritance.html