docker build . -t named-api
docker run -d -p 8080:8080 named-api