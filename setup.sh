#!/bin/bash
git clone https://github.com/dgraph-io/badger && cd badger && go install && cd .. && rm -rf badger
go build .
mv run.sh /etc/cron.hourly/
sudo chmod +x /etc/cron.hourly/run.sh