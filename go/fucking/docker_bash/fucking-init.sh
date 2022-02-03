#!/bin/bash

'''
*--------
    Hey Pfeff
    - This function just creates a docker image and connects to it with exec
    - If there is another parameter, that will be the image pulled by docker
    - Otherwise, just use the ubuntu image
*--------
'''
if [[ -z "$1" ]]; then
    echo "usage: fucking-init <container-name> <image>"
    return
fi

if [[ -z "$2" ]]; then
    docker run --name $1 -dt ubuntu
    docker exec -it $1 "/bin/bash"
    return
fi

docker run --name $1 -dt $2
docker exec -it $1 "/bin/bash"