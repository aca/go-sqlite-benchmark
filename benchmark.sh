#!/usr/bin/env bash

(cd gosqlite && go test -bench . -benchtime 1000x)
echo
(cd mattn && go test -bench . -benchtime 1000x)
echo
(cd ncruces && go test -bench . -benchtime 1000x)
