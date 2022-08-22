# Multi-language errors with Golang

### Commands

Run server

```bash
go mod vendor
go run .
```

Call API

```bash
curl --location --request POST 'localhost:8080/trans/users?locale=vi' \
--header 'Content-Type: application/json' \
--data-raw '{
    "age": 100,
    "email": "liam@gmailcom"
}'
# {"message":"Tên là trường bắt buộc. Tuổi phải bằng hoặc nhỏ hơn 90. Email không hợp lệ"}

curl --location --request POST 'localhost:8080/trans/users?locale=en' \
--header 'Content-Type: application/json' \
--data-raw '{
    "age": 100,
    "email": "liam@gmailcom"
}'
# {"message":"Name is a required field. Age must be 90 or smaller. Email is invalid"}
```
