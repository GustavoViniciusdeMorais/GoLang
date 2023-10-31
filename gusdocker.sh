 #!/bin/bash
echo "--- Docker Manager ---";
echo "";

if [ $1 = "-h" ]
then
    echo "Instructions:";
    echo "";
    echo " 1- --start                         will start container by docker-compose file config";
    echo " 1- --stop                          will stop container by docker-compose file config";
    echo " 3- [container_name] [sh, bash]     will enter container with provided mode";
    echo "";
fi

if [ -n "$1" ]
then
    if [ $1 = "--start" ]
    then
        docker container stop $(docker container ls -aq) && docker network prune -f
        docker compose up -d --build
    else
        if [ $1 = "--stop" ]
        then
            docker compose down
        fi

        if [ -n "$2" ]
        then
            docker exec -it -u 0 $1 $2
        fi
    fi
else
    echo "Please informe docker container name and enter mode [sh, bash]!";
    echo "";
fi