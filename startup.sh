#!/bin/bash
echo Wait for servers to be up
sleep 10

HOSTPARAMS="--host database --insecure"
SQL="/cockroach/cockroach.sh sql $HOSTPARAMS"

$SQL -e "CREATE USER IF NOT EXISTS maxroach;"
$SQL -e "CREATE DATABASE IF NOT EXISTS spero;"
$SQL -e "GRANT ALL ON DATABASE spero TO maxroach;"
