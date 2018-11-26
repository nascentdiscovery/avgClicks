# avgClicks

This application provides an average number of clicks per day by
Country for the default group associated with a Bitly access token.
It can easily be extended to support reporting this data by individual
bitlink as well.

Please create a bitly access token via https://bitly.is/accesstoken
before proceeding.

The environment variable BITLY_API_TOKEN should be populated with your
API token before launching the program.

--help for usage; a int can be provided via the argument --lookback to
override the default of 30 days of historical data.
