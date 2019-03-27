to get docker ip address(local) for specified container use:

docker ps -a
docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' containerID
