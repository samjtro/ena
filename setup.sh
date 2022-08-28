#!/bin/bash
go build .
echo '0 * * * * .$HOME/sn/sn' >> /var/spool/cron/crontabs/$HOME