### Docker Learning


#### Pull image

> docker pull <image_name>

#### Run a container in detached mode with specific name

> docker run -d -name <container_name> <image_name>


#### Run a docker-compose.yml file

> docker-compose -f <docker_compose_file_name.yml> up

#### remove conatiners and images started by a docker-compose

> docker-compose -f <docker_compose_file_name.yml> down

#### To stop and restart 

- > docker-compose -f <docker_compose_file_name.yml> start

- > docker-compose -f <docker_compose_file_name.yml> stop
