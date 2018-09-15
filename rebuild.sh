if [ "$1" == "nocontainer" ]; then
  echo "Starting new container.."
else
  echo "Restarting container..!"
  docker stop location_service && docker rm location_service
fi
docker build -t location_service . && docker run --network br0 -d --name location_service --env-file .env location_service
