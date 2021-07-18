# run deployment

# run elasticsearch
sudo sysctl -w vm.max_map_count=262144
docker compose -f elasticsearch/docker-compose.yml up -d

# run postgres
docker compose -f postgres/docker-compose.yml up -d

# run services
docker compose -f services/auth/docker-compose.yml up -d
docker compose -f services/product/docker-compose.yml up -d
docker compose -f services/order/docker-compose.yml up -d
