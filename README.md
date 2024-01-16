# Planets API

planets api for information of planets.

### Requirements

Install [docker desktop](https://www.docker.com/products/docker-desktop/)

Just run build once. If it stacks, restart docker desktop.

```shell
docker build --no-cache -t planets-info/api .
```

Can run as many times as you want.

```shell
docker run -d -p 8081:80 planets-info/api
```

Now goto: http://localhost:8081/ to access swagger.

## License

MIT
