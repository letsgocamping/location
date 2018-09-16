if [ "$1" == "nocontainer" ]; then
  echo "Starting new container.."
else
  echo "Restarting container.."
  docker stop lgc-location && docker rm lgc-location
fi
docker build -t lgc-location . && docker run --name lgc-location -p 5678:8080 --network br0 --env-file .env -d lgc-location
