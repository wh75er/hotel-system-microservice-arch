# Markdown **Hotel Service** API Specification

Simple and small markdown api documentation

## Contents:

List of contents:

- `[GET /api/v1/hotels?isReady=<boolValue>&Offset={offset}&Limit={limit}` - Get all hotels(light data)](#get-all-hotels)
- `[GET /api/v1/hotels/<hotelUuid>` - get info about one hotel(heavy data)](#get-hotel)
- `[POST /api/v1/hotels` - add hotel](#add-hotel)
- `[PATCH /api/v1/hotels/<hotelUuid>` - change hotel information](#patch-hotel)
- `[DELETE /api/v1/hotels/<hotelUuid>` - delete specified hotel(need to make request to reservation service)](#delete hotel)

---

- `[GET /api/v1/hotels/<hotelUuid>/reviews` - get all reviews for specified hotel](#get-all-reviews)
- `[GET /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - get info about specified review](#get-review)
- `[PATCH /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - change review text](#patch-review)
- `[DELETE /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - delete review](#delete-review)
- `[POST /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - add review](#add-review)

---

- `[GET /api/v1/hotels/<hotelUuid>/packages` - get all packages for hotel](#get-all-packages)
- `[GET /api/v1/hotels/<hotelUuid>/packages/<packagesUuid>` - get packages](#get-package)
- `[PATCH /api/v1/hotels/<hotelUuid>/packages/<packageUuid>` - patch package](#patch-package)
- `[DELETE /api/v1/hotels/<hotelUuid>/packages/<packageUuid>` - delete package](#delete-package)
- `[POST /api/v1/hotels/<hotelUuid>/packages` - add package](#add-package)

## Light Hotel Data

```json
{
    "name": "string",
    "hotelUuid": "string",
    "photos": ["uuid", "uuid", "uuid"],
    description: "string",
    "registrationDate": "timestamp"
    "country": "string",
    "city": "string",
    "address": "string",
    "latitude": "real",
    "longitude": "real"
    "isReady": "bool"
}
```

## Light Room Data

```json
{
    "roomType": "string",
    "amount": "int",
    "beds": int,
    "hotelUuid": "uuid",
    "roomUuid": "uuid"
}
```

## Package Data

```json
{
    "sleeps": "int",
    "offers": ["string"],
    "hotelUuid": "uuid",
    "nightPrice": "int",
    "packageUuid": "uuid"
}
```

## Review Data

```json
{
    "creationDate": "timestamp",
    "test": "string",
    "userUuid": "uuid",
    "isAnonymous": "bool",
    "photos": ["uuid"],
    "hotelUuid": "uuid"
}
```

## Heavy Room Data

```json
{
    "roomType": "string",
    "amount": "int",
    "beds": int,
    "hotelUuid": "uuid",
    "roomUuid": "uuid"
    "packages": ["Package Data"]
}
```

## Heavy Hotel Data

```json
{
    "name": "string",
    "hotelUuid": "string",
    "photos": ["uuid", "uuid", "uuid"],
    description: "string",
    "registrationDate": "timestamp"
    "country": "string",
    "city": "string",
    "address": "string",
    "latitude": "real",
    "longitude": "real"
    "isReady": "bool"
    "rooms": ["Heavy Room Data"]
    "reviews": ["Review Data"]
}
```

### Get user deposit info

1. `GET /api/v1/deposits/{user-uuid}` - Get user's deposit info
    - Request Body:
    (empty)
    - Responses:
        - `200`

            ```json
            {
                "data": {
                    "id": 1,
                    "userUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "deposit": 0,
                    "creationDate": 1630753631
                }
            }

            ```

        - `400`

            ```json
            {
                "error": "invalid data format",
                "description": "failed to convert provided userUuid field into UUID"
            }

            ```

        - `404`

            ```json
            {
                "error": "user not found",
            }

            ```

        - `500`

            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }

            ```
