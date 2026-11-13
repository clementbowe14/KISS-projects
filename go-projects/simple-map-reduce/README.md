# Simple MapReduce

This project rebuilds MapReduce as a local word counter. It uses goroutines and files to explore how workers count words independently and then combine counts for the same word.

1. The generator in `main.go` writes 50 random word files to `files/`.
2. Five map workers read the input and hash each word to choose its reducer. They write partial counts to `mapped-files/`. Every occurrence of a word goes to the same reducer, even when different workers read it.
3. Once the map workers finish, five reducers combine the partial counts for their partitions. Each writes a JSON file in `reduced-files/`.

## Run the demo

Use Go 1.26.5 or later. From this directory, run `go run .`. If you cloned the repo elsewhere, first change the generator’s `parentDir` in `main.go` to `"files/wordfile"` so it writes to the directory the coordinator reads.

The run rewrites the sample inputs and reducer outputs. For counts from a single run, start with an empty `mapped-files/` directory. The mapper reads existing intermediate counts, including the checked-in examples, and adds to them.

To trace a word through the pipeline, pick one from a reducer’s JSON output and look for it in the mapper files with the same reducer number.
