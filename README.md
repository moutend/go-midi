go-midi
========

[![GitHub release](https://img.shields.io/github/release/moutend/go-midi.svg?style=flat-square)][release]
[![CircleCI](https://circleci.com/gh/moutend/go-midi.svg?style=svg&circle-token=e3db72ca7a0c643a8c7ed00d0d6b6ad36f4c70df)](https://circleci.com/gh/moutend/go-midi)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/moutend/go-midi/releases
[status]: https://circleci.com/gh/moutend/go-midi
[license]: https://github.com/moutend/go-midi/blob/master/LICENSE

WIP

Package midi implements reading and writing standard MIDI file.

## Installation

```console
go get github.com/moutend/go-midi
```

## Usage

WIP

## Numbering of notes

There are two conventions for notes in MIDI. The most common is where C3 is `0x3c` and the another is where C4 is `0x3c`. In this package, where C3 is `0x3c`.

## MIDI files for testing

The MIDI files located at `testdata` were composed by Nao. Check her great works:

- [星のカービィ　MIDI - みかんの旅](http://mikannotabi.blog31.fc2.com/blog-entry-6.html)

## About MIDI

- [MIDI File Format Specifications · colxi/midi-parser-js Wiki](https://github.com/colxi/midi-parser-js/wiki/MIDI-File-Format-Specifications)
- [0xff21 and 0xff20 in MIDI](https://groups.google.com/forum/#!topic/comp.music.midi/_MIjgi-8xQQ)
- [Computer Music/ MIDI Key Number Chart](http://computermusicresource.com/midikeys.html)
- [OpenMIDIProject - Documentations](http://openmidiproject.osdn.jp/documentations_en.html)

## Contributing

1. Fork ([https://github.com/moutend/go-midi/fork](https://github.com/moutend/go-midi/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
