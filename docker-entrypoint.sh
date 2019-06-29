#!/bin/sh
rm /usr/bin/sudare_contents
go build -o /usr/bin/sudare_contents
exec /usr/bin/sudare_contents
