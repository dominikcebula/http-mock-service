[![CI Pipeline](https://github.com/dominikcebula/http-mock-service/actions/workflows/go.yml/badge.svg)](https://github.com/dominikcebula/http-mock-service/actions/workflows/go.yml)

# HTTP Mock Service

## Intro

This repository contains HTTP Mock Service application that can be used to easily create declarative HTTP Service Mock.
You can use this project to easily create mocks for REST JSON APIs, REST XML APIs using declarative approach in YAML:

```yaml
  - request:
      path: ^/hello_json$
    response:
      code: 200
      headers:
        Content-Type: application/json
        Cache-Control: no-cache
        X-Powered-By: http-mock-service
      body: |
        {
            "product_id": "12345",
            "name": "Wireless Bluetooth Headphones",
            "brand": "TechSound",
            "price": 79.99,
        }
```

## Features

* Request matching based on:
  * Request Path using regexp
  * HTTP Method (GET, POST, PUT, DELETE, ...)
* Response generation, including:
  * Response code
  * Response headers
  * Response body
* Ability to simulate response delays
* HTTP Server Configuration
  * Listen address
  * Listen port

## Running the project

### Docker

Run the following command to execute HTTP Mock Service using Docker:

```shell
docker run --rm -p 8080:8080 dominikcebula/http-mock-service:v0.5.1
```

After running above command you will see on the screen:

```text
2024/05/25 20:34:06 Reading config file...
2024/05/25 20:34:06 Config file loaded.
2024/05/25 20:34:06 Creating request handler(s)...
2024/05/25 20:34:06 Created 4 request handler(s).
2024/05/25 20:34:06 Creating HTTP Mock Server...
2024/05/25 20:34:06 Listening for connections on address "0.0.0.0:8080"...
```

You can now execute HTTP Requests against executed HTTP Mock Service.

For example to test it using `curl` you can execute the following command:

```shell
curl localhost:8080/hello_json
```

Which will return following results:

```text
{
    "product_id": "12345",
    "name": "Wireless Bluetooth Headphones",
    "brand": "TechSound",
    "price": 79.99,
}
```

You can declare your own behaviors using your own `config.yaml` file.

For example, below `my_config.yaml` file will generate new json response available under `/hello_order` endpoint:

```yaml
server:
  host: 0.0.0.0
  port: 8080

rules:
  - request:
      path: ^/hello_order$
    response:
      code: 200
      headers:
        Content-Type: application/json
        Cache-Control: no-cache
        X-Powered-By: http-mock-service
      body: |
        {
          "orderId": "A003",
          "date": "2024-05-25",
          "userId": "98765",
          "shippingAddress": {
            "street": "123 Main St",
            "city": "Anytown",
            "state": "CA",
            "zip": "12345",
            "country": "USA"
          }
        }
```

Having new `my_config.yaml` file created, HTTP Mock Service needs to be executed with the new `my_config.yaml` file
using the
following command:

```shell
docker run --mount type=bind,source=./my_config.yaml,target=/app/config.yaml --rm -p 8080:8080 dominikcebula/http-mock-service:v0.5.1
```

Introduced API endpoint can be queried using curl:

```shell
curl localhost:8080/hello_order
```

which will produce:

```text
{
  "orderId": "A003",
  "date": "2024-05-25",
  "userId": "98765",
  "shippingAddress": {
    "street": "123 Main St",
    "city": "Anytown",
    "state": "CA",
    "zip": "12345",
    "country": "USA"
  }
}
```

## Author

Dominik Cebula

- https://dominikcebula.com/
- https://blog.dominikcebula.com/
- https://www.udemy.com/user/dominik-cebula/
