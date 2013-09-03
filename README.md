# piglow-fader

Make the LEDs on the piglow fade in a pleasing way.

## Usage

Cross compile for the Raspberry pi with:

````GOOS=linux GOARM=6 GOARCH=arm go build````

The binary produced takes a delay in milliseconds, defaulting to 1000:

./piglow_fader -delay 500

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## Author

* Will Jessop, @will_j, will@willj.net