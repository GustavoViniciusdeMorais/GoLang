# API Config and request

```sh
sudo go run api.go
```
```sh
sudo curl -X GET localhost:9003/products

sudo curl -X POST localhost:9003/products

sudo curl -X DELETE localhost:9003/products

sudo curl -X GET localhost:9003/products/1

sudo curl -d '{"id":7,"name":"test","inventory":"test","price":13}' -X POST localhost:9003/products
```