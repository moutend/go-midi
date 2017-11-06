go-midi
========

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/moutend/mediumctl/blob/master/LICENSE

WIP

Read and write standard MIDI file.

```console
go get golang.org/x/tools/cmd/stringer
```


## Numbering of keys (notes)

> There are TWO conventions for numbering keys (notes) in MIDI. The most common is the one below where MIDDLE C (note #60; $3C) is C3 (C in the 3rd octave). However, another convention was adopted by Yamaha Corp. for their synthesizer products which parallels the Octave Designation System used in Music Education formulated by the Acoustical Society of America. In that convention, Middle C is designated "C4". The "C3 Convention" is the most commonly used octave designation system on standard MIDI keyboards and this is the convention we will use for this class.

# Test Data

[星のカービィ　MIDI - みかんの旅]()http://mikannotabi.blog31.fc2.com/blog-entry-6.html

## Documents

- [Computer Music/ MIDI Key Number Chart](http://computermusicresource.com/midikeys.html)
- https://www.cs.cmu.edu/~music/cmsip/readings/MIDI%20tutorial%20for%20programmers.html
- http://openmidiproject.osdn.jp/documentations_en.html

## Contributing

1. Fork ([https://github.com/moutend/go-midi/fork](https://github.com/moutend/go-midi/fork))
1. Create a feature branch
1. Add changes
1. Run `go fmt`
1. Commit your changes
1. Open a new Pull Request

## Author

[Yoshiyuki Koyanagi](https://github.com/moutend)
