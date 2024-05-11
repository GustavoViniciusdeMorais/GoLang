# API Config and request

```sh
sudo go run api.go
```
```sh
sudo curl -X GET localhost:9003/products

sudo curl -X GET localhost:9003/products/1

sudo curl -d '{"name":"test","inventory":10,"price":13}' -X POST localhost:9003/products
```