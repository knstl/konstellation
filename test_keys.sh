# --------------------------------------
# Keys
# --------------------------------------

# "address": "darc1x68e44c90hy49ppxqr2n0lawkv5w5vzp98uw4j"
# "pubkey": "darcpub1addwnpepq2evdf2hzflltqamfkqld3sc59ex40m68d7udzhj93569p5ltq4s54ntuxa"
# shellcheck disable=SC2034
declare -A key1=(
  [name]="mask"
  [password]="mask1234"
  [mnemonic]="remain local lens arena card squeeze fire head ill win legend practice omit broccoli road rocket honey promote drink stamp wet sign van picnic"
)

# "address": "darc1d22ccl8xpzzzldl28l9gs9htrgaatkaxjwskkl"
# "pubkey": "darcpub1addwnpepqvz0qh2yha89c4shy5nseptsfqmwd3m9xudkkzagc9jkxdnkctwl56sqvxf"
# shellcheck disable=SC2034
declare -A key2=(
  [name]="satoshi"
  [password]="satoshi1"
  [mnemonic]="creek round dish idea parade balance embody rate sail wheel casino suspect cute transfer pizza square cabbage valid divert gap chimney legal hurt interest"
)

# "address": "darc1z6u8hddknlp33z79rra7q4gjlw3thwqrdzyplq
# "pubkey": "darcpub1addwnpepqf3eyh3tq4gf5n8eh7xg8vf5q3k6k7c00q2nx4lrp6r0e9zrz66lj4d5rf9"
# shellcheck disable=SC2034
declare -A key3=(
  [name]="nakamoto"
  [password]="nakamoto"
  [mnemonic]="bachelor miracle cook decade chalk table blouse learn vanish menu general transfer engage coast scout rookie glove affair target glide own manual consider scrub"
)

#  "address": "darc15h6uhzufhe0d0avuk54zcqw0t66qefeyc3vttf"
#  "pubkey": "darcpub1addwnpepqdszyhp5ypp5lx27gr933wwn2yek4u9m7twzwjge8ycf2jg7wyry6nncd73"
# shellcheck disable=SC2034
declare -A key4=(
  [name]="vitalik"
  [password]="vitalik1"
  [mnemonic]="hair reunion riot fish hawk arch buzz debris arch easily search glove private rival boat resist chaos cause panda icon pave shock egg ability"
)

#  "address": "darc1ejgxhtvj6c9n7d7g29jmsxhnn6wh2j8rll0vfc"
#  "pubkey": "darcpub1addwnpepqdgyr7nptmtu9fkqcqyhatq6dg2zat53lynfd8tkzfurgtj82np7qx00dxc"
# shellcheck disable=SC2034
declare -A key5=(
  [name]="buterin"
  [password]="buterin1"
  [mnemonic]="other original trouble craft hard loan ostrich aim drastic team absent kiwi matrix dose engage cup novel humor brave budget stage label future exile"
)
declare -a keys=(key1 key2 key3 key4 key5)

for key in "${keys[@]}"; do
  KEY_NAME="${key}[name]"
  KEY_PASSWORD="${key}[password]"
  KEY_MNEMONIC="${key}[mnemonic]"

  if [ -n "${!KEY_NAME}" ]
  then
    echo "--------------------------------------"
    echo "${!KEY_NAME}"
    echo "${!KEY_PASSWORD}"
    echo "${!KEY_MNEMONIC}"

    {
      #    echo "${!KEY_PASSWORD}";
      echo "${!KEY_MNEMONIC}"
      echo
    } | konstellationcli keys add "${!KEY_NAME}" --dry-run --interactive
  fi
done
