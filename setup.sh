#!/bin/bash
go build .
#echo '0 * * * * .$HOME/sn/sn' | sudo tee -a /var/spool/cron/crontabs/$USER
#sudo bash -c "echo '0 * * * * .$HOME/sn/sn' >> /var/spool/cron/crontabs/$USER"