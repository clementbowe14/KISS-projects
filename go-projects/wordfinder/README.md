# Wordfinder

Wordfinder builds a trie from a word list. Each node has one child slot for each lowercase letter from `a` to `z`. Inserting `apple` and `apricot` reuses the same path for `ap` before the words branch at the third letter.

## Run the demo

Use Go 1.26.5 or later. The demo writes its built-in list to `words.txt` on each run. If you cloned the repo elsewhere, change the absolute path in the `os.Open` call in `main.go` to `"words.txt"` so it reads the file it just wrote.

From this directory, run `go run .`. The program searches for `apple` and `banana` and prints `true` on two lines.

Try changing one search to `app`. It also returns `true` because the search succeeds as soon as it follows all the supplied letters. An end-of-word marker in `trie.go` would let the trie distinguish a complete word from a prefix.
