#!/bin/sh
docker build ./ -t sudare_contents
docker run -t --init --name sudare_contents -v `pwd`:/go/src/sudare_contents/ sudare_contents