# Warehouse

The following repository contains a warehouse application written in Go. This readme will provide you with instruction on how to use the api, run and test it.

## Prerequisite

- [Go 1.17](https://go.dev/dl/)
- [Make](https://www.gnu.org/software/make/)

## Build

To build this service, run `make build`. This will download dependencies and produce an executable.

## Start

The server can be started by running: `make start`. This will start the executable described in the previous section. The application loads products and inventory provided in `products.json`, `inventory.json` and starts a server on port `8080`.

## Test

Run tests by running `make test`. This will run all tests located in `warehouse_test.go`.

### API

The service provides a simple REST api for fetching products and buying products. See `warehouse.postman_collection.json` for a postman collection.

#### Get products

Request

```http
GET localhost:8080/product
```

Response

```json
[
  {
    "name": "Dining Chair",
    "articles": [
      {
        "art_id": "1",
        "amount_of": "4"
      },
      {
        "art_id": "2",
        "amount_of": "8"
      },
      {
        "art_id": "3",
        "amount_of": "1"
      }
    ],
    "quantity": 2,
    "productId": 0
  },
  {
    "name": "Dinning Table",
    "articles": [
      {
        "art_id": "1",
        "amount_of": "4"
      },
      {
        "art_id": "2",
        "amount_of": "8"
      },
      {
        "art_id": "4",
        "amount_of": "1"
      }
    ],
    "quantity": 1,
    "productId": 1
  }
]
```

#### buy products

```http request
POST localhost:8080/product/:productId/buy
```

Response 200:

```json
"purchase successful"
```

Response 400:

```json

```

## Choices/compromises

- Storage: In this iteration, I have used an in-memory struct to store articles and products. In a future version it would make sense to replace this storage with a database.
- Buy/sell more items at once: Currently one can only buy/sell one product. It would make sense to support buying/selling of more products since that would be an expected use case in the future.
- Update product stock: Product stock is calculated on get products. This could have been a separate function and be reused on buy/sell.
- Multiple requests: The current implementation does not handle multiple requests well. One could be in a state were a product is potentially sold twice.
- Docker: The current implementation rely on the user having Go installed in order to run it. It would make sense to use Docker to containerise the service for portability and ease of deployments.
- Response: the current responses could be more information. Forsell/buy products one could elaborate on why the request failed and use more http statuses.
