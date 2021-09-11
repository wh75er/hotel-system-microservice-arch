# Markdown **Hotel Service** API Specification

Simple and small markdown api documentation

## Contents:

List of contents:

- [`GET /api/v1/hotels?isReady=<boolValue>&offset={offset}&limit={limit}` - Get all hotels(light data)](#get-all-hotels)
- [`GET /api/v1/hotels/<hotelUuid>` - get info about one hotel(heavy data)](#get-hotel)
- [`POST /api/v1/hotels` - add hotel](#add-hotel)
- [`PATCH /api/v1/hotels/<hotelUuid>` - change hotel information](#patch-hotel)
- [`DELETE /api/v1/hotels/<hotelUuid>` - delete specified hotel(need to make request to reservation service)](#delete-hotel)

---

- [`GET /api/v1/hotels/<hotelUuid>/reviews?offset={offset}&limit={limit}` - get all reviews for specified hotel](#get-all-reviews)
- [`GET /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - get info about specified review](#get-review)
- [`PATCH /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - change review text](#patch-review)
- [`DELETE /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - delete review](#delete-review)
- [`POST /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - add review(need to make request to authorization service)](#add-review)

---

- [`GET /api/v1/hotels/<hotelUuid>/rooms` - get all rooms(heavy data)](#get-all-rooms)
- [`PATCH /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>` - change room data](#patch-rooms)
- [`POST /api/v1/hotels/<hotelUuid>/rooms` - add room](#add-room)
- [`DELETE /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>` - delete room](#delete-room)

## Light Hotel Data

```json
{
    "name": "string",
    "hotelUuid": "string",
    "photos": ["uuid", "uuid", "uuid"],
    "description": "string",
    "creationDate": "timestamp",
    "country": "string",
    "city": "string",
    "address": "string",
    "isReady": "bool"
}
```

## Room Data

```json
{
    "roomType": "string",
    "amount": "int",
    "beds": "int",
    "hotelUuid": "uuid",
    "roomUuid": "uuid",
    "offers": ["string"],
    "nightPrice": "int",
    "creationDate": "timestamp"
}
```

## Review Data

```json
{
    "reviewUuid": "uuid",
    "userUuid": "uuid",
    "hotelUuid": "uuid",
    "text": "string",
    "isAnonymous": "bool",
    "photos": ["uuid"],
    "creationDate": "timestamp"
}
```

## Heavy Hotel Data

```json
{
    "name": "string",
    "hotelUuid": "string",
    "photos": ["uuid", "uuid", "uuid"],
    "description": "string",
    "creationDate": "timestamp",
    "country": "string",
    "city": "string",
    "address": "string",
    "isReady": "bool",
    "rooms": ["Room Data"],
    "reviews": ["Review Data"]
}
```

### Get all hotels

1. `GET /api/v1/hotels?isReady=<boolValue>&offset={offset}&limit={limit}` - Get all hotels(light data)
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "name": "best hotel",
                        "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                        "photos": ["dd6f4022-c663-436c-a5f5-3ad003c894e8", "dd6f4022-c663-436c-a5f5-3ad003c894e8"],
                        "description": "best hotel in the universe",
                        "creationDate": 1630753631,
                        "country": "japan",
                        "city": "okinawa",
                        "address": "bumbum, 2",
                        "isReady": "true"
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "offset and limit should be integers > 0"
            }
            ```
        - `404`
            ```json
            {
                "error": "no hotels found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Get hotel

2. `GET /api/v1/hotels/<hotelUuid>` - get info about one hotel(heavy data)
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "name": "best hotel",
                    "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "photos": ["dd6f4022-c663-436c-a5f5-3ad003c894e8", "dd6f4022-c663-436c-a5f5-3ad003c894e8"],
                    "description": "best hotel in the universe",
                    "creationDate": 1630753631,
                    "country": "japan",
                    "city": "okinawa",
                    "address": "bumbum, 2",
                    "isReady": "true",
                    "rooms": [
                        {
                            "roomUuid": "22d65c61-21c9-423f-a154-947992a83a6e",
                            "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                            "beds": 5,
                            "amount": 10,
                            "creationDate": 1631381372,
                            "packages": [
                                {
                                    "packageUuid": "5b7b0ef0-9c97-4adc-9712-fd976ecdee31",
                                    "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                                    "sleeps": 2,
                                    "offers": ["free cancelation", "food included", "diving suits"],
                                    "nightPrice": 500.8,
                                    "creationDate": 1631381372
                                }
                            ]
                        }
                    ],
                    "reviews": [
                        {
                            "reviewUuid": "0581c2f5-f32c-4299-921a-4c10bd9d62b4",
                            "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                            "userUuid": "3d2dcd12-a0cc-412d-9714-23e875850340",
                            "text": "Awesome hotel!",
                            "isAnonymous": "false",
                            "photos": ["0fe34b71-9a11-4bc3-b2d7-7857e5da7925", "8ee021a9-cbe9-4d4c-9ff4-9db1562492da"]
                        }
                    ]
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "provided uuid is not valid"
            }
            ```
        - `404`

            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Add hotel

3. `POST /api/v1/hotels` - add hotel
    - Request Body:
        ```json
        {
            "name": "best hotel",
            "photos": ["dd6f4022-c663-436c-a5f5-3ad003c894e8", "dd6f4022-c663-436c-a5f5-3ad003c894e8"],
            "description": "best hotel in the universe",
            "country": "japan",
            "city": "okinawa",
            "address": "bumbum, 2"
        }
        ```
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "name": "best hotel",
                    "hotelUuid": "13570e16-4d98-4823-b266-eeb3f4776eda",
                    "photos": ["dd6f4022-c663-436c-a5f5-3ad003c894e8", "dd6f4022-c663-436c-a5f5-3ad003c894e8"],
                    "description": "best hotel in the universe",
                    "creationDate": 1630753631,
                    "country": "japan",
                    "city": "okinawa",
                    "address": "bumbum, 2",
                    "isReady": "false"
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "city name is longer than 100 characters"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Patch hotel

4. `PATCH /api/v1/hotels/<hotelUuid>` - change hotel information
    - Request Body:
        ```json
        {
            "photos": ["dd6f4022-c663-436c-a5f5-3ad003c894e8", "dd6f4022-c663-436c-a5f5-3ad003c894e8"]
        }
        ```
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "uuid is incorrect"
            }
            ```
        - `404`
            ```json
            {
                "error": "not found hotel"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Delete hotel

5. `DELETE /api/v1/hotels/<hotelUuid>` - delete specified hotel(need to make request to reservation service)
    - Request Body:
    (empty)
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "uuid is incorrect"
            }
            ```
        - `404`
            ```json
            {
                "error": "not found hotel"
            }
            ```
        - `422`
            ```json
            {
                "error": "service unavailable",
                "description": "unable to reach reservations service"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

---

### Get all reviews

1. `GET /api/v1/hotels/<hotelUuid>/reviews?offset={offset}&limit={limit}` - get all reviews for specified hotel(#get-all-reviews)
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "reviewUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                        "userUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                        "hotelUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                        "text": "Awesome hotel",
                        "isAnonymous": "true",
                        "photos": ["fbd49c54-209e-4b3a-a811-ebf5c937a697", "fbd49c54-209e-4b3a-a811-ebf5c937a697"],
                        "creationDate": "1631368874"
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "no reviews found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Get review

2. `GET /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - get info about specified review
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "reviewUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "userUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "hotelUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "text": "Awesome hotel",
                    "isAnonymous": "true",
                    "photos": ["fbd49c54-209e-4b3a-a811-ebf5c937a697", "fbd49c54-209e-4b3a-a811-ebf5c937a697"],
                    "creationDate": "1631368874"
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "no reviews found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### PATCH review

3. `PATCH /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - change review text
    - Request Body:
        ```json
        {
            "photos": ["fbd49c54-209e-4b3a-a811-ebf5c937a697", "fbd49c54-209e-4b3a-a811-ebf5c937a697"]
        }
        ```
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "not found review"
            }
            ```
        - `422`
            ```json
            {
                "error": "service unavailable",
                "description": "unable to reach users service"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Delete review

4. `DELETE /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - delete review
    - Request Body:
    (empty)
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "not found review"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Add review

5. `POST /api/v1/hotels/<hotelUuid>/reviews/<reviewUuid>` - add review(need to make request to authorization service)
    - Request Body:
        ```json
        {
            "userUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
            "hotelUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
            "text": "Awesome hotel",
            "isAnonymous": "true",
            "photos": ["fbd49c54-209e-4b3a-a811-ebf5c937a697", "fbd49c54-209e-4b3a-a811-ebf5c937a697"]
        }
        ```
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "reviewUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "userUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "hotelUuid": "fbd49c54-209e-4b3a-a811-ebf5c937a697",
                    "text": "Awesome hotel",
                    "isAnonymous": "true",
                    "photos": ["fbd49c54-209e-4b3a-a811-ebf5c937a697", "fbd49c54-209e-4b3a-a811-ebf5c937a697"],
                    "creationDate": "1631368874"
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "no reviews found"
            }
            ```
        - `422`
            ```json
            {
                "error": "service unavailable",
                "description": "unable to reach users service"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

---

### Get all packages

1. `GET /api/v1/hotels/<hotelUuid>/packages` - get all packages for hotel
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "sleeps": "int",
                        "offers": ["string"],
                        "hotelUuid": "uuid",
                        "nightPrice": "int",
                        "packageUuid": "uuid",
                        "creationDate": "timestamp"
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "no packages found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Get package

2. `GET /api/v1/hotels/<hotelUuid>/packages/<packagesUuid>` - get packages
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "sleeps": "int",
                    "offers": ["string"],
                    "hotelUuid": "uuid",
                    "nightPrice": "int",
                    "packageUuid": "uuid",
                    "creationDate": "timestamp"
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "package not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Patch package

3. `PATCH /api/v1/hotels/<hotelUuid>/packages/<packageUuid>` - patch package
    - Request Body:
    ```json
    {
        "nightPrice": "int"
    }
    ```
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "package not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Delete package

4. `DELETE /api/v1/hotels/<hotelUuid>/packages/<packageUuid>` - delete package
    - Request Body:
    (empty)
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "package not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Add package

5. `POST /api/v1/hotels/<hotelUuid>/packages` - add package
    - Request Body:
        ```json
        {
            "sleeps": "int",
            "offers": ["string"],
            "hotelUuid": "uuid",
            "nightPrice": "int"
        }
        ```
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "sleeps": "int",
                        "offers": ["string"],
                        "hotelUuid": "uuid",
                        "nightPrice": "int",
                        "packageUuid": "uuid",
                        "creationDate": 1631381372
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

---

### Get all rooms

1. `GET /api/v1/hotels/<hotelUuid>/rooms` - get all rooms(heavy data)
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "roomType": "family room",
                        "amount": 34,
                        "beds": 4,
                        "hotelUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                        "roomUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                        "creationDate": "timestamp",
                        "packages": [
                            {
                                "sleeps": 3,
                                "offers": ["diving", "food included", "free cancelation"],
                                "hotelUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                                "nightPrice": 504.33,
                                "packageUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                                "creationDate": 1631381372
                            }
                        ]
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Get all room packages

2. `GET /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>/packages` - get all room's pacakges
    - Request Body:
    (empty)
    - Responses:
        - `200`
            ```json
            {
                "data": [
                    {
                        "sleeps": 3,
                        "offers": ["diving", "food included", "free cancelation"],
                        "hotelUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                        "nightPrice": 504.33,
                        "packageUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                        "creationDate": 1631381372
                    }
                ]
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "no packages for required room found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Add package for room

3. `POST /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>/packages` - add existing package to room
    - Request Body:
        ```json
        {
            "packageUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8"
        }
        ```
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Delete package from room

4. `DELETE /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>/packages/<packageUuid>` - delete existing package from room
    - Request Body:
    (empty)
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Patch rooms

5. `PATCH /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>` - change room data
    - Request Body:
        ```json
        {
            "beds": 5
        }
        ```
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Add room

6. `POST /api/v1/hotels/<hotelUuid>/rooms` - add room
    - Request Body:
        ```json
        {
            "roomType": "chill couple",
            "amount": 30,
            "beds": 2,
            "hotelUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8"
        }
        ```
    - Responses:
        - `200`
            ```json
            {
                "data": {
                    "roomType": "chill couple",
                    "amount": 30,
                    "beds": 2,
                    "hotelUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                    "roomUuid": "a88186fb-ffea-4a29-881e-0ecfe9231bb8",
                    "creationDate": "timestamp"
                }
            }
            ```
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```

### Delete room

7. `DELETE /api/v1/hotels/<hotelUuid>/rooms/<roomUuid>` - delete room
    - Request Body:
    (empty)
    - Responses:
        - `204`
        - `400`
            ```json
            {
                "error": "invalid data format",
                "description": "invalid hotel uuid"
            }
            ```
        - `404`
            ```json
            {
                "error": "hotel not found"
            }
            ```
        - `500`
            ```json
            {
                "error": "something went wrong with the repository",
                "description": "database connection is down"
            }
            ```
