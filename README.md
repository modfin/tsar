# TSAR - Text Search All Rows
> A library for specific search use case


tsar is a go library that creates a sorted reverse word index that uses binary search to look up. 
It also provides a small DSL for querying files directly

tsar works directly on text files and creates an index, this can be saved directly to a file and used for querying the text.

see example/simple.go