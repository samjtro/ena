#!/bin/bash
BADGER=/home/$USER/sn/badger.bak
if [ ! -f "$BADGER" ]; then
    sudo badger restore --dir $BADGER
fi
./home/$USER/sn/sn
sudo badger backup --dir /tmp/badger -f $BADGER