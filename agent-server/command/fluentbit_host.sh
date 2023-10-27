#!/bin/bash
set -xe

###################
# how to use the script, example
# bash fluentbit_host.sh 11.22.33.44
###################


main() {
    param=$1
    h=$(echo $param | awk -F "="  "{print $2}")
    sed -i "s#host .*#host ${h}#g" fluent-bit.conf
}

# for example
# $1 : k1=v1,k2=v2
main "$@"
