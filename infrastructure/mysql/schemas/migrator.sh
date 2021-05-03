#!/bin/bash

skeema diff local -p$MYSQL_ROOT_PASSWORD

read -p "Migrateしますか？ (y/n) :" YN
if [ "${YN}" = "y" ]; then
  skeema push local -p$MYSQL_ROOT_PASSWORD
else
  echo "bye!";
fi
