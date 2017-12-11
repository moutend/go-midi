# Maximize Velocity of Note On Events

In this example, reads a given standard MIDI file, maximize the velocity of all Note ON events which value is greater than 0 and creates a new file named `output.mid`.

## Usage

```console
go run main.go music.mid
```

Then `output.mid` will be created. Listen the midi file and find out the all velocity of Note On events are set to 127.

## LICENSE

MIT
