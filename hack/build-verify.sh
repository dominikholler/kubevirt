#!/usr/bin/env bash

set -e

function report_dirty_build() {
    set +e
    echo "Build is not clean:"
    kubevirtci/cluster-up/virtctl.sh version
    git status
    exit 1
}

# Check that "clean" is reported at least once
if [ -z "$(kubevirtci/cluster-up/virtctl.sh version | grep clean)" ]; then
    report_dirty_build
fi

# Check that "dirty" is never reported
if [ -n "$(kubevirtci/cluster-up/virtctl.sh version | grep dirty)" ]; then
    report_dirty_build
fi
