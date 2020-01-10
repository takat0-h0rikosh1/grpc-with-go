# Cafe CLI App (gRPC with Go)

## Using

```sh
$ go run main.go

$ grpcurl -plaintext localhost:50052 cafe.Cafe/GetMenus
{
  "menus": [
    {
      "name": "coffee",
      "price": 100
    },
    {
      "name": "late",
      "price": 110
    },
    {
      "name": "mocha",
      "price": 120
    }
  ]
}

$ grpcurl -plaintext -d '{"name":"coffee"}'  localhost:50052 cafe.Cafe/Order
{
  "price": 100
}
```


## Docker run

```sh
$ docker build -t cafe .

$ docker run -it --rm -p 50052:50052 cafe
```
