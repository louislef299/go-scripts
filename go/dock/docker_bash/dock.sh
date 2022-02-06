#!/bin/bash

'''
*--------
    Hey Pfeff
      This function just creates a docker image and connects to it with exec
      If there is another parameter, that will be the image pulled by docker
      Otherwise, just use the ubuntu image

    Here is what I need from you:
    - I want to be able to add a flags to this command
    - Instead of having $1 and $2 I want to use flags
      - For instance: -name or -image or -saveOnClose
    - After you implement the flags, I want to save the container when I close 
      the session and kill/delete the container, all with a flag -saveOnClose
      - For this, use the following docker commands:
        * docker kill
        * docker rm
      - Use --help to learn more about the command options
*--------
'''
if [[ -z "$1" ]]; then
    echo "usage: dock <container-name> <image>"
    return
fi

if [[ -z "$2" ]]; then
    docker run --name $1 -dt ubuntu
    docker exec -it $1 "/bin/bash"
    return
fi

docker run --name $1 -dt $2
docker exec -it $1 "/bin/bash"