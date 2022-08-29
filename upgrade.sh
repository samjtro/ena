#!/bin/bash
cd .. && rm -rf sn && cd ../../.. && sudo bash -c "rm -rf /etc/cron.hourly/run.sh /tmp/badger" && cd ~
git clone https://github.com/samjtro/sn && cd sn
sudo chmod +x setup.sh && ./setup.sh