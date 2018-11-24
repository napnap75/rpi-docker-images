#!/bin/sh

set -e

crond -b -L /var/log/cron.log && tail -f /var/log/cron.log
