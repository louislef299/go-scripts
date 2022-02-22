# The ***Dock*** Project

This script makes it easy to work with docker in a repetitive manner. Docker has a lot of different functionalities, but the goal is to make it easy enough to work as a simple file system that is isolated from the root file system. We will be utilizing the docker API and golang to create this function.

## Functions

- `dock -init`: This creates a new container with an ubuntu image
  - Relevant docker commands:
    1. docker run --name <optional-name> -dt ubuntu
    2. docker exec -it <optional-name> "/bin/bash"
- `dock -new=<image>`: Create a new container with a given image
    1. docker run --name <optional-name> -dt <image>
    2. docker exec -it <optional-name> "/bin/bash"
- `dock -restore=<imageID>`: Restores a saved container session
- `dock -list`: Lists all active containers
    1. docker ps -a
- `dock -save`: Saves the current container environment
- `dock -saveAndClose`: Saves the environment and shuts it down
- `dock -listSaved`: Lists all of your saved sessions and when they were committed
- `dock -kill=<containerID>`: Kills a running container. `***WARNING***` *If* you kill a container and it was not saved, you will lose any progress made on the container. 
    1. docker kill <containerID>
    2. docker rm <containerID>

## Resources

- [Learn Docker in 100 seconds](https://www.youtube.com/watch?v=Gjnup-PuquQ)
- [Advanced Docker](https://www.youtube.com/watch?v=gAkwW2tuIqE)
- [Learn Go in 100 seconds](https://www.youtube.com/watch?v=446E-r0rXHI)
- [Docker Cheat Sheet](https://www.docker.com/sites/default/files/d8/2019-09/docker-cheat-sheet.pdf)


[1/29]
In order to build a docker image, use the `docker build` command. This will build an image in the read-only space, but will be available for when wanting to perform a `docker run` on a specific image. In order to persist data among images, we will need to create a persistent volume to attach to the container. This can be achieved using the `docker volume create` function. The volume will be saved under the /var/lib/docker/volumes directory. You can then mount the volume upon creation of the container using `docker run -v <data_volume>:/var/lib/mysql sql` to mount an sql volume for instance. This will mount the volume into the read/write layer. 

You can also mount an existing folder to a docker container using a process called *bind mounting*. In order to do this, replace the <data_volume> with the location of the folder that you want to mount to the container. So, there are two types of mounting: 
    1. Volume mounting - creates a volume from the volumes directory
    2. Bind mounting - mounts a directory directly from the docker host

Also, the `-v` option to mount a volume is an old way of mounting a volume to a docker container. Using the `--mount` command is much more verbose way of mounting a volume and is preferred. So, the better way of writing the above command would be:
```
docker run --mount type=bind,source=/data/mysql,target=/var/lib/mysql mysql
```
Where source is the location on the docker host and target is the location on the container.