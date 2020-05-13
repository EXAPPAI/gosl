#!/bin/bash

FILES="*.go"

echo
echo "monitoring:"
echo $FILES
echo
echo

while true; do
    inotifywait -q -e modify $FILES
    echo
    echo
    echo
    echo
    # go test -test.run="InitDefault"
    # go test -test.run="InitWithParams"
    # go test -test.run="HideBorders"
    go test -test.run="PlotSimpleCurve"
done
