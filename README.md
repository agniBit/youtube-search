# youtube-search


Build and run API server
```sh
docker build -t agnibit-api-server -f docker/Dockerfile-api  .
docker run -e DATABASE_URL=<database_url> -e YOUTUBE_API_KEYS=<youtube_api_keys> -d -p 8081:8080 agnibit-api-server
```


Build and run cron server for fetching youtube videos and saving into database

```sh
docker build -t agnibit-cron-server -f docker/Dockerfile-cron  .
docker run -e DATABASE_URL=<database_url> -e YOUTUBE_API_KEYS=<youtube_api_keys> agnibit-cron-server
```
