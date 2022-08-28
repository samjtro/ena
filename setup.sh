#!/bin/bash
git clone https://github.com/dgraph-io/badger && cd badger && go install && cd .. && rm -rf badger
go build .
sudo bash -c "mv run.sh /etc/cron.hourly/ && chmod +x /etc/cron.hourly/run.sh && mv /home/$USER/sn/config.env /etc/cron.hourly/"