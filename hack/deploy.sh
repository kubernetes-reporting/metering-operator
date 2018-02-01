#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
source ${DIR}/util.sh

export CHARGEBACK_NAMESPACE=${CHARGEBACK_NAMESPACE:-chargeback-ci}
export SKIP_DELETE_CRDS=true

CHARGEBACK_NAMESPACE="$(sanetize_namespace "$CHARGEBACK_NAMESPACE")"

: "${CUSTOM_CHARGEBACK_SETTINGS_FILE:=}"
: "${UNINSTALL_CHARGEBACK:=true}"
: "${INSTALL_CHARGEBACK:=true}"

while true; do
    echo "Checking namespace status"
    NS="$(kubectl get ns "$CHARGEBACK_NAMESPACE" -o json --ignore-not-found)"
    if [ "$NS" == "" ]; then
        echo "Namespace ${NAMESPACE} does not exist"
        break
    ***REMOVED***
    if [ "$(echo "$NS" | jq -r '.status.phase')" == "Terminating" ]; then
        echo "Waiting for namespace $CHARGEBACK_NAMESPACE termination to complete before continuing"
    ***REMOVED***
    sleep 2
done

echo "Creating namespace $CHARGEBACK_NAMESPACE"
kubectl create ns "$CHARGEBACK_NAMESPACE" || true

if [ "$UNINSTALL_CHARGEBACK" == "true" ]; then
    echo "Uninstalling chargeback"
    ./hack/uninstall.sh
***REMOVED***
    echo "Skipping uninstall"
***REMOVED***

until [ "$(kubectl -n $CHARGEBACK_NAMESPACE get deployments -l app=chargeback-helm-operator -o json | jq '.items | length' -r)" == "0" ]; do
    echo 'waiting for chargeback-helm-operator deployment to be deleted'
    sleep 5
done

until [ "$(kubectl -n $CHARGEBACK_NAMESPACE get pods -o json | jq '.items | length' -r)" == "0" ]; do
    echo 'waiting for chargeback pods to be deleted'
    sleep 5
done

if [ "$INSTALL_CHARGEBACK" == "true" ]; then
    echo "Installing chargeback"
    ./hack/install.sh
***REMOVED***
    echo "Skipping install"
***REMOVED***

until [ "$(kubectl -n $CHARGEBACK_NAMESPACE get pods -l app=chargeback-helm-operator -o json | jq '.items | map(.status.containerStatuses[] | .ready) | all' -r)" == "true" ]; do
    echo 'waiting for chargeback-helm-operator pods to be ready'
    sleep 5
done
echo "chargeback helm-operator is ready"

EXPECTED_POD_COUNT=6
until [ "$(kubectl -n $CHARGEBACK_NAMESPACE get pods -o json | jq '.items | length' -r)" == "$EXPECTED_POD_COUNT" ]; do
    echo 'waiting for chargeback pods to be created'
    sleep 10
done
echo "all of the chargeback pods have been started"

until [ "$(kubectl -n $CHARGEBACK_NAMESPACE get pods  -o json | jq '.items | map(.status.containerStatuses | if . == null then [{ready: false}] ***REMOVED*** . end | .[].ready) | all' -r)" == "true" ]; do
    echo 'waiting for all pods to be ready'
    sleep 10
done
echo "chargeback pods are all ready"
