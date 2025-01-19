# wizcraft
`**W**orking **I**n the microservice **Z**one: **CRAFT**ing Scalable Architectures`

Backend API Server for MSaaS

## Prerequisites
- For Dev (택일)
  - Docker / Docker-compose
  - Go 1.23.2 or later
- For Production ( To be decided )

## How to start?

### 1. Clone the repository
```bash
$ git clone [git-repo-url]
$ cd MSaaS_back
```

#### setup .env
- create .env file
    ```bash
    # This is the local path where the microservices repository created and the database data saved.
    $ mkdir -p /home/user/msaas/db-volume /home/user/msaas/projects-volume
    $ vi .env
    ```
- .env example
    ```text
    POSTGRES_USER=root
    POSTGRES_PASSWORD=admin123
    POSTGRES_DB=msaas
    POSTGRES_HOST=localhost
    POSTGRES_PORT=5432
    POSTGRES_VOLUME=/home/user/msaas/db-volume
    WIZCRAFT_PORT=8080
    WIZCRAFT_VOLUME=/home/user/msaas/projects-volume
    ```

### 2.1. Run the server by Docker-compose
#### dev mode (supporting hot-reload)
백엔드 코드 변경 시 반영됩니다.
- 생성
```bash
$ docker compose -f docker-compose.dev.yml up -d
```
- 종료
```bash
$ docker compose -f docker-compose.dev.yml down
```
- 동작 확인
```bash
$ docker ps -a
$ curl localhost:${WIZCRAFT_PORT}
```

#### static dev version
백엔드 코드 변경 시 반영되지 않습니다.
```bash
$ docker compose -f docker-compose.yml up -d
$ docker ps -a
```

### 2.2. Run the server by Go