#!/bin/sh
docker build ./ -t sudare_contents
docker run -t --init --name sudare_contents -v `pwd`:/go/src/sudare_contents/ -p 5563:5563/tcp sudare_contents