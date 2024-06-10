#!/bin/bash

# Create a customer
curl -X POST http://localhost:8080/api/customers -H "Content-Type: application/json" -d '{
  "name": "Emre Koc",
  "username": "emrekoc",
  "password": "password123"
}'

# Login and get the token
TOKEN=$(curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
  "username": "emrekoc",
  "password": "password123"
}' | jq -r '.token')

# Create an account
curl -X POST http://localhost:8080/api/accounts -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "customer_id": 1,
  "balance": 1000
}'

# Create another account
curl -X POST http://localhost:8080/api/accounts -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "customer_id": 1,
  "balance": 1000
}'

# Update the account
curl -X PUT http://localhost:8080/api/accounts/1 -H "Authorization: $TOKEN" -H "Content-Type: application/json" -d '{
  "balance": 1500
}'

# List accounts
curl -X GET http://localhost:8080/api/accounts -H "Authorization: $TOKEN"

# Transfer money
curl -X POST http://localhost:8080/api/accounts/transfer -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" -d '{
  "from_account_id": 1,
  "to_account_id": 2,
  "amount": 100
}'
