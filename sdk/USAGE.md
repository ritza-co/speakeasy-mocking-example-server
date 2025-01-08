<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
from holiday_destinations import HolidayDestinations

with HolidayDestinations() as holiday_destinations:

    res = holiday_destinations.general.read_root_get()

    # Handle response
    print(res)
```

</br>

The same SDK client can also be used to make asychronous requests by importing asyncio.
```python
# Asynchronous Example
import asyncio
from holiday_destinations import HolidayDestinations

async def main():
    async with HolidayDestinations() as holiday_destinations:

        res = await holiday_destinations.general.read_root_get_async()

        # Handle response
        print(res)

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->