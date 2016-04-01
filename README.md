# tz
What time is it at Lat, Lon?

This function translates date/time information between time zones.
Provide in input:
- A Unix timestamp, i.e. number of seconds since 1970-01-01 00:00 UTC
- A latitude and longitude
The returned string contains wall time at the place of interest at
the moment of the provided time stamp. To convert between (lat,lon) and
an offset in seconds, a call is made to Google Map API to determine
in which time zone the place is located, and if daylight saving time
is observed at the moment of the requested time stamp.

MIT License (c) nicolas314

