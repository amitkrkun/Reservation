#!/bin/bash

go build -o bookings cmd/web/*.go

./bookings -dbname=bookings -dbuser=amitkun -cache=false -production=false
