#!/bin/zsh

(cd benchmarks; go test -bench=. -benchmem)
