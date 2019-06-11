# technical-test-2
Technical test 2 for ANZ - API GoLang app with CI   

## About

This is a simple go json https api which will expose 3 different endpoints

- `/` will return: 
```
{
    "Welcome":"Hello World!"
}
```

- `/health` will return: 
```
{
    "health":"Still alive!"
}
```

- `/info` will return: 
```
{
    "myapplication":[ 
    {
        "version":"v1.0.0",
        "lastcommitsha":"f80ff8f841b23cc91511538a9411931b587970ce",
        "description":"pre-interview technical test"
        }
    ]
}
```

## Dependencies 
- Docker engine
- git 
- golang 1.12

## Building the API
To build this api via docker image run the following docker command

```bash
docker build . --build-arg version="v1.0.0" --build-arg last_commit_sha=$(git rev-parse --verify HEAD) -t <tag-name>
```

## Running the API
To run the this api via docker run the following docker command
```bash
docker run -p 80:8000 <tag-name>
```

## Testing
To test the 3 API endpoint run the following curl commands

```
curl http://localhost/
{"Welcome":"Hello World!"}

curl http://localhost/health
{"health":"Still alive!"}

curl http://localhost/info 
{"myapplication":[ {"version":"v1.0.0","lastcommitsha":"f80ff8f841b23cc91511538a9411931b587970ce", "description":"pre-interview technical test"}]}
```

## CI
this repo has a travis CI pipeline that will push to docker hub

## Risks
- No logging
- No monitoring
- Only http
- Version is set in docker file
- No deployment to anywhere

## Future Features
- add logging
- add monitoring
- add https
- add deployment to a kubernetes
- create auto versioning
