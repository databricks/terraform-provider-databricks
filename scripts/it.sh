#!/bin/bash

if [[ $@ == *"--debug"* ]]; then
    export TF_LOG="DEBUG"
    export TF_LOG_PATH=$PWD/tf.log
    export TF_ACC_LOG_PATH=$PWD/tf.log
    echo "[*] To see debug logs: tail -f $PWD/tf.log"
fi

TMPDIR=/tmp TF_ACC=1 gotestsum \
    --format short-verbose \
    --raw-command go test -v -json \
    -test.timeout 35m \
    -run $1 ../../... 2>&1 | tee out.log

echo "✓ To output of existing tests: less $PWD/out.log"

FAILURES=$(grep "\-\-\- FAIL" out.log | sed 's/--- FAIL: / \* \[ \] /g' | sort)
PASSES=$(grep PASS out.log | grep Test | sort | sed 's/PASS/ \* \[x\]/')

cat <<-EOF > test-report.log
$1
---
${FAILURES}
${PASSES}
EOF

cat test-report.log