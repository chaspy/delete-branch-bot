#!/bin/bash
SCRIPT_DIR=$(cd $(dirname $0); pwd)
export GITHUB_APP_IDENTIFIER=17630
export GITHUB_APP_INSTLLATIONID=348213
export GITHUB_PRIVATE_KEY=`awk '{printf "%s\\n", $0}' $SCRIPT_DIR/private-key.pem`
