# Destinations
(*destinations*)

## Overview

### Available Operations

* [get_destinations_destinations_get](#get_destinations_destinations_get) - Get Destinations
* [create_destination_destinations_post](#create_destination_destinations_post) - Create Destination
* [delete_destination_destinations_destination_id_delete](#delete_destination_destinations_destination_id_delete) - Delete Destination
* [get_destination_by_id_destinations_destination_id_get](#get_destination_by_id_destinations_destination_id_get) - Get Destination By Id
* [update_destination_destinations_destination_id_put](#update_destination_destinations_destination_id_put) - Update Destination

## get_destinations_destinations_get

Retrieve a list of all holiday destinations in the database, optionally filtered by minimum rating.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.destinations.get_destinations_destinations_get()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `min_rating`                                                        | *Optional[float]*                                                   | :heavy_minus_sign:                                                  | Filter destinations by minimum rating (0 to 5).                     |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[List[models.Destination]](../../models/.md)**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| models.HTTPValidationError | 422                        | application/json           |
| models.APIError            | 4XX, 5XX                   | \*/\*                      |

## create_destination_destinations_post

Add a new holiday destination to the database.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.destinations.create_destination_destinations_post(country="Indonesia", description="Beautiful beaches and vibrant culture.", name="Bali", rating=4.8)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `country`                                                           | *str*                                                               | :heavy_check_mark:                                                  | The country where the destination is located.                       | Indonesia                                                           |
| `description`                                                       | *str*                                                               | :heavy_check_mark:                                                  | A brief description of the destination.                             | Beautiful beaches and vibrant culture.                              |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | The name of the holiday destination.                                | Bali                                                                |
| `rating`                                                            | *float*                                                             | :heavy_check_mark:                                                  | Rating of the destination (0 to 5).                                 | 4.8                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Response

**[models.Destination](../../models/destination.md)**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| models.HTTPValidationError | 422                        | application/json           |
| models.APIError            | 4XX, 5XX                   | \*/\*                      |

## delete_destination_destinations_destination_id_delete

Remove a holiday destination from the database by its ID.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    holiday_destinations.destinations.delete_destination_destinations_destination_id_delete(destination_id=0)

    # Use the SDK ...

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `destination_id`                                                    | *int*                                                               | :heavy_check_mark:                                                  | The ID of the destination to delete.                                | 0                                                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| models.HTTPValidationError | 422                        | application/json           |
| models.APIError            | 4XX, 5XX                   | \*/\*                      |

## get_destination_by_id_destinations_destination_id_get

Retrieve details of a specific holiday destination by its ID.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.destinations.get_destination_by_id_destinations_destination_id_get(destination_id=0)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `destination_id`                                                    | *int*                                                               | :heavy_check_mark:                                                  | The ID of the destination to retrieve.                              | 0                                                                   |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Response

**[models.Destination](../../models/destination.md)**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| models.HTTPValidationError | 422                        | application/json           |
| models.APIError            | 4XX, 5XX                   | \*/\*                      |

## update_destination_destinations_destination_id_put

Update details of a specific holiday destination by its ID.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.destinations.update_destination_destinations_destination_id_put(destination_id=1, country="Indonesia", description="Beautiful beaches and vibrant culture.", name="Bali", rating=4.8)

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `destination_id`                                                    | *int*                                                               | :heavy_check_mark:                                                  | The ID of the destination to update.                                | 1                                                                   |
| `country`                                                           | *str*                                                               | :heavy_check_mark:                                                  | The country where the destination is located.                       | Indonesia                                                           |
| `description`                                                       | *str*                                                               | :heavy_check_mark:                                                  | A brief description of the destination.                             | Beautiful beaches and vibrant culture.                              |
| `name`                                                              | *str*                                                               | :heavy_check_mark:                                                  | The name of the holiday destination.                                | Bali                                                                |
| `rating`                                                            | *float*                                                             | :heavy_check_mark:                                                  | Rating of the destination (0 to 5).                                 | 4.8                                                                 |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |                                                                     |

### Response

**[models.Destination](../../models/destination.md)**

### Errors

| Error Type                 | Status Code                | Content Type               |
| -------------------------- | -------------------------- | -------------------------- |
| models.HTTPValidationError | 422                        | application/json           |
| models.APIError            | 4XX, 5XX                   | \*/\*                      |