# General
(*general*)

## Overview

### Available Operations

* [read_root_get](#read_root_get) - Read Root

## read_root_get

Root endpoint, returns a friendly greeting.

### Example Usage

```python
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.general.read_root_get()

    # Handle response
    print(res)

```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `retries`                                                           | [Optional[utils.RetryConfig]](../../models/utils/retryconfig.md)    | :heavy_minus_sign:                                                  | Configuration to override the default retry behavior of the client. |

### Response

**[Any](../../models/.md)**

### Errors

| Error Type      | Status Code     | Content Type    |
| --------------- | --------------- | --------------- |
| models.APIError | 4XX, 5XX        | \*/\*           |