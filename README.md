# avgClicks

Repository also available at https://github.com/nascentdiscovery/avgClicks

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

To run from binary on an armvl6/7 32bit system, simply launch the provided
binary after setting the BITLY_API_TOKEN env var.

To build from source, please install the most recent version of golang, copy
this repository locally, cd into it, and then `go build .`

If you are stuck with an older version of go, please set your GOPATH env var,
copy this repository into $GOPATH/src/github.com/nascentdiscovery/, cd into it, and
then `go get github.com/spf13/pflag; go get github.com/bitly/go-simplejson; go build .`
