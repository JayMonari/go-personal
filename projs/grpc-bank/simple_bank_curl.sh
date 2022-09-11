#!/usr/bin/env bash
set -euo pipefail

token=$(mktemp)

create_user() {
  curl localhost:9001/users -d \
    "$(printf \
    '{"username":"%s","full_name":"CURL LURC","email":"%s@cmail.com","password":"password123"}' \
    "$1" "$1")"
}

make_accounts() {
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$schemehost"'/accounts' \
    -d '{"currency":"CAD"}'
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$schemehost"'/accounts' \
    -d '{"currency":"USD"}'
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$schemehost"'/accounts' \
    -d '{"currency":"EUR"}'
}

case "$1" in
  renew)
    curl "$schemehost"/tokens/renew_access \
      -d "$(printf '{"refresh_token":"%s"}', "$2")" | jq
    ;;
  new_user)
    create_user "$2"
    login "$2"
    make_accounts
    ;;
  list_accounts)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$schemehost"'/accounts?page_size=5&page_id=1' | jq
    ;;
  get_account)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$schemehost/accounts/$3" | jq
    ;;
  transfer)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$schemehost/transfers" \
      -d "$(printf \
'{"from_account_id":%s,"to_account_id":%s,"amount":%s,"currency":"%s"}' \
          "$3" "$4" "$5" "$6")" \
      | jq
    ;;
esac

