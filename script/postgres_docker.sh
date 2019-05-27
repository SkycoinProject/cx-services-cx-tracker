docker pull postgres
docker run --name cx-tracker-db -e POSTGRES_PASSWORD=supersecretpass -p 5434:5432 -d postgres