#!/bin/bash
go build .
mv run.sh /etc/cron.hourly/
sudo chmod +x /etc/cron.hourly/run.sh