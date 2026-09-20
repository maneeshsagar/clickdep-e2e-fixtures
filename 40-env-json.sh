#!/bin/sh
{
  printf '{"role":"frontend-ui","has_database_url":%s,"env":{' "$( [ -n "$DATABASE_URL" ] && echo true || echo false )"
  first=1
  for kv in $(env | grep '^E2E_' | sed 's/"/\\"/g' | awk -F= '{print $1}'); do
    v=$(printenv "$kv" | sed -e 's/\\/\\\\/g' -e 's/"/\\"/g' | sed ':a;N;$!ba;s/\n/\\n/g')
    [ $first -eq 1 ] || printf ','
    printf '"%s":"%s"' "$kv" "$v"; first=0
  done
  printf '}}'
} > /usr/share/nginx/html/env.json
