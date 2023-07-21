# Web server with MariaDB and GoFiber


## Test Locally

### 1. Clone Repository to local machine
```shell
git clone https://github.com/amalmadhu06/mariadb-fiber-go.git
```
### 2. Change directory
```shell
cd mariadb-fiber-go
```
### 3. Download dependencies
```shell
go mod download
```
### 4. Set up a MariaDB database
Set up a MariaDB database and add schema and data in the `pkg/db/sql` folder
### 5. Configure .env
Rename the `.env.example` to `dev.env` and update the content according to your database set up.
### 6. Run the server locally
```shell
make run 
```
### Endpoints
```shell
curl -X GET http://127.0.0.1:3000/offer/us
```
```shell
curl -X GET http://127.0.0.1:3000/offer/fr
```
```shell
curl -X GET http://127.0.0.1:3000/offer/ca
```
```shell
curl -X GET http://127.0.0.1:3000/offer/br
```