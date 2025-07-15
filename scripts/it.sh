#!/bin/bash

if [[ $@ == *"--debug"* ]]; then
    export TF_LOG="DEBUG"
    export TF_LOG_PATH=$PWD/tf.log
    export TF_ACC_LOG_PATH=$PWD/tf.log
    echo "[*] To see debug logs: tail -f $PWD/tf.log"
fi

if [ "" == "$TEST_FILTER" ]; then
    TEST_FILTER="TestAcc"
fi

export TMPDIR=/tmp
export TF_ACC=1

go run gotest.tools/gotestsum@latest --junitfile junit.xml --format short-verbose \
    --raw-command go test -v -json \
    -test.timeout 30m \
    -run $TEST_FILTER ./... 2>&1 | tee out.log

FAILURES=$(grep "\-\-\- FAIL" out.log | sed 's/--- FAIL: / \* \[ \] /g' | sort)
PASSES=$(grep PASS out.log | grep Test | sort | sed 's/PASS/ \* \[x\]/')

cat <<-EOF > test-report.log
Test run: $TEST_FILTER
---
${FAILURES}
${PASSES}
EOF

cat test-report.log
