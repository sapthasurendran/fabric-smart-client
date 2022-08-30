#!/bin/bash
# Copyright the Hyperledger Fabric contributors. All rights reserved.
#
# SPDX-License-Identifier: Apache-2.0
set -euo pipefail

CONFIGPATH=$HOME/.config/softhsm2

mkdir -p $CONFIGPATH

if [[ "$OSTYPE" == "darwin"* ]]; then
    # Install softhsm
    brew install softhsm

    # Move softhsm config file inside config path
    export SOFTHSM2_CONF=$CONFIGPATH/softhsm2.conf

    # Create tokens inside tmp folder
    echo directories.tokendir = /tmp > $CONFIGPATH/softhsm2.conf

    #Initialize tokens
    softhsm2-util --init-token --slot 0 --label "ForFSC" --so-pin 1234 --pin 98765432

else
    #Install softhsm
    sudo apt-get install -y softhsm2

    #Create directory to store hsm tokens
    sudo mkdir -p /var/lib/softhsm/tokens
    sudo chmod -R 777 /var/lib/softhsm

    cp /usr/share/softhsm/softhsm2.conf $CONFIGPATH

    #Initialize tokens
    sudo softhsm2-util --init-token --slot 0 --label "ForFSC" --so-pin 1234 --pin 98765432

fi


