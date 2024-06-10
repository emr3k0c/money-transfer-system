## Installation

### 1. Clone the Repository

Clone the project repository:

```sh
git clone https://github.com/emr3k0c/money-transfer-system.git
cd money-transfer-system
go mod tidy
export JWT_SECRET='u8$FpD3!vG7*Qw4^zRtLj9&eB2Kx#YnM'
go run main.go
```

### 2. Access the API

Create a Customer:

```sh
curl -X POST http://localhost:8080/api/customers -H "Content-Type: application/json" -d '{
  "name": "Emre Koc",
  "username": "emrekoc",
  "password": "secret"
}'
```

Login:

```sh
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
  "username": "emrekoc",
  "password": "password123"
}'
```

Create an Account:

```sh
TOKEN=<jwt_token>
curl -X POST http://localhost:8080/api/accounts -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "customer_id": 1,
  "balance": 1000
}'
```

Update an Account:

```sh
TOKEN=<jwt_token>
curl -X PUT http://localhost:8080/api/accounts/1 -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "balance": 1500
}'
```

List Accounts:

```sh
TOKEN=<jwt_token>
curl -X GET http://localhost:8080/api/accounts -H "Authorization: $TOKEN"
```

Transfer Money:

```sh
TOKEN=<jwt_token>
curl -X POST http://localhost:8080/api/accounts/transfer -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 100
}'
```
