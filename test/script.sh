#!/bin/bash

# Format command
mkdir test/output/format -p
go run main.go format test/input/whiteBox_100x100.png test/output/format/whiteBox_100x100.png
go run main.go format test/input/flower_720x720.png test/output/format/flower_720x720.png
go run main.go format test/input/blackBox_100x100.jpg test/output/format/blackBox_100x100.jpg

# Resize command
mkdir test/output/resize -p
go run main.go resize test/input/text_109x40.png test/output/resize/text_218x80.png 218 80

# Crop command
mkdir test/output/crop -p
go run main.go crop test/input/flower_720x720.jpg test/output/crop/flower_450x500.jpg 450 500
go run main.go crop test/input/bird_1280x720.png test/output/crop/bird_600x256.png 600 256

# Multi(Multiple) command
mkdir test/output/multi -p
go run main.go crop test/input/flower_720x720.jpg test/output/multi/flower_450x500.png 450 500

# TODO !!! Add tests here

# Compare Results
go test test/imgproc_test.go -v ./..