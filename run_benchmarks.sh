#!/bin/bash

(cd benchmarks; go test -bench=. -benchmem)
