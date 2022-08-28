#!/bin/bash
BADGER=/home/$USER/sn/badger.bak
if [ ! -f "$badger" ]; then
    badger load --dir $BADGER
fi
./home/$USER/sn/sn
badger backup --dir /tmp/badger -f BADGER