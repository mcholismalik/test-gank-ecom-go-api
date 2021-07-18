# run elasticsearch
sudo sysctl -w vm.max_map_count=262144
docker compose -f dependencies/elasticsearch/docker-compose.yml up -d

# run postgres
docker compose -f dependencies/postgres/docker-compose.yml up -d
