# run services
docker compose -f services/auth/docker-compose.yml up -d
docker compose -f services/product/docker-compose.yml up -d
docker compose -f services/order/docker-compose.yml up -d
