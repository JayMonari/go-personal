#!/usr/bin/env bash
set -euxo pipefail

echo '{new_user|list_accounts|get_account} <username>'

basepath='http://localhost:9001'
token=$(mktemp)

create_user() {
curl localhost:9001/users -d \
  '{"username":'"\"$1\""',"full_name":"CURL LURC","email":"curl.test@cmail.com","password":"password123"}'
}

login() {
  curl "$basepath"/users/login \
    -d '{"username":'"\"$1\""',"password":"password123"}' \
    | jq -r '.AccessToken' > "$token"
}

make_accounts() {
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$basepath"'/accounts' \
    -d '{"currency":"CAD"}'
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$basepath"'/accounts' \
    -d '{"currency":"USD"}'
  curl -H "Authorization:Bearer $(cat "$token")" \
    "$basepath"'/accounts' \
    -d '{"currency":"EUR"}'
}

case "$1" in
  new_user)
    create_user "$2"
    login "$2"
    make_accounts
    ;;
  list_accounts)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$basepath"'/accounts?page_size=5&page_id=1' | jq
    ;;
  get_account)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$basepath/accounts/$3" | jq
    ;;
  transfer)
    login "$2"
    curl -H "Authorization:Bearer $(cat "$token")" \
      "$basepath/transfers" \
      -d "$(printf \
          '{"from_account_id":%s,"to_account_id":%s,"amount":%s,"currency":"%s"}' \
          "$3" "$4" "$5" "$6")" \
      | jq
    ;;
esac

