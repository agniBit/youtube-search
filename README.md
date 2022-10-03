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


API request
```
curl --location --request GET '127.0.0.1:8081/v1/youtube/videos?title=cricket&limit=20&offset=10'
```

API usage 

| Query | Description |
| ------ | ------ |
| title | search video by title (partial match) |
| description | search video by description (partial match) |
| search | search video by title or description (partial match) |
| offset | skip result |
| limit | total desired results |


### Examples :- 

Get top 10 results order by publish date
```
curl --location --request GET '127.0.0.1:8081/v1/youtube/videos?limit=10'
```


Get top 10 results order by publish date match by title of video
```
curl --location --request GET '127.0.0.1:8081/v1/youtube/videos?limit=10&title=virat'
```


Get 20 results, skip first 10 result order by publish date match by title of video
```
curl --location --request GET '127.0.0.1:8081/v1/youtube/videos?offset=10&limit=10&title=virat'
```
