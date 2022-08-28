#!/bin/bash
cd .. && rm -rf sn && cd ../../.. && sudo rm -rf /etc/cron.hourly/run.sh && cd ~
git clone https://github.com/samjtro/sn && cd sn
sudo chmod +x setup.sh && ./setup.sh