# Random Value Generator Service

This service provides an API for generating random values and retrieving them using a unique ID. It is implemented in Go and supports various types of random values with customizable parameters.

Link to task from Avito: https://github.com/avito-tech/pro-backend-trainee-assignment

## Features
-[x] Generate Random Values: Generate random values of different types with customizable length.
-[x] Retrieve Values: Retrieve previously generated values using a unique ID.
-[x] Idempotent Requests: Ensure consistent results for requests with the same requestId.
-[x] Logging: Logs request details, including headers and generated values, to a relational database.
-[x] Docker Support: The service is containerized and available as a Docker image.
-[x] (Extra) Tested with K6: Load testing with K6 to ensure performance and scalability.
-[ ] Unit Testing: Includes unit tests for robust functionality.

## API Endpoints
### POST /api/generate/
Generates a random value with specified parameters.

**Request Parameters:**

- `type` (required): Type of the random value. Options:
    - `1`: Random string.
    - `2`: Random number.
    - `3`: Random UUID/GUID.
    - `4`: Random alphanumeric string.
- `length` : Length of the generated value.

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/generate/" -H "Content-Type: application/json" -d '{"type": "string", "length": 10}'
```

**Example Response:**
```json
{
  "id": "JScmkasd-asd912NJ-1asd3Cs4-1s12323sd-xaJCVasd",
  "value": "SA21njcdscsj#@nscscjn"
}
```

### POST /api/retrieve
Retrieves a previously generated value using its unique ID.

**Request Parameters:**

- `id` (required): ID of the random value. 

**Example Request:**
```bash
curl -X POST "http://localhost:8080/api/generate/" -H "Content-Type: application/json" -d '{"type": "string", "length": 10}'
```

**Example Response:**
```json
{
  "id": "JScmkasd-asd912NJ-1asd3Cs4-1s12323sd-xaJCVasd",
  "value": "SA21njcdscsj#@nscscjn"
}
```

## Idempotent Requests
To ensure idempotency, you can include a requestId as a query parameter or header. Multiple requests with the same requestId will return the same ID and value.

Example Request with Idempotency Key:

```bash
curl -X POST "http://localhost:8080/api/generate/?requestId=unique-request-id" -H "Content-Type: application/json" -d '{"type": "number"}'
```

## Docker
The service is available as a Docker image. To run it locally:

```bash
docker compose up --build -d
```

## Logging
Each POST and GET request is logged into a relational database. The following information is captured:

- User-Agent
- Request ID
- Token
- URL
- Count (number of retrieve requests)
