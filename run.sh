#!/bin/bash
BADGER=/home/$USER/sn/badger.bak
if [ ! -f "$badger" ]; then
    sudo badger load --dir $BADGER
fi
./home/$USER/sn/sn
sudo badger backup --dir /tmp/badger -f BADGER