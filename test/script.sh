#!/bin/bash

# Format command
mkdir test/output/format -p
go run main.go format test/input/whiteBox_100x100.png test/output/format/whiteBox_100x100.png
go run main.go format test/input/flower_720x720.png test/output/format/flower_720x720.png
go run main.go format test/input/blackBox_100x100.jpg test/output/format/blackBox_100x100.jpg
###  testing jpg is difficult due to compression

# resize
mkdir test/output/resize -p
go run main.go resize test/input/text_109x40.png test/output/resize/text_218x80.png 218 80


go test test/imgproc_test.go -v ./..