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
init_flag = '' # -i
new_flag = '' # -n
restore_flag = '' # -r
list_flag = '' # -l
save_flag = '' # -s
saveAndClose_flag = '' # -x
listSaved_flag = '' # -a
kill_flag = '' # -k
verbose = '' # -v

print_usage() {
    echo "usage: dock <container-name> <image>"
    return
}

init() {
  docker run --name $1 -dt ubuntu
  docker exec -it $1 "/bin/bash"
}

new()


while getopts 'in:r:lsxak:' flag; do
  case "${flag}" in
    i) init_flag='true' ;;
    n) new_flag="${OPTARG}" ;;
    r) restore_flag="${OPTARG}" ;;
    l) list_flag = 'true' ;;
    s) save_flag = 'true' ;;
    x) saveAndClose_flag = 'true' ;;
    a) listSaved_flag = 'true' ;;
    k) kill_flag = "${OPTARG}" ;;
    v) verbose='true' ;;
    *) print_usage
       exit 1 ;;
  esac
done

docker run --name $1 -dt $2
docker exec -it $1 "/bin/bash"