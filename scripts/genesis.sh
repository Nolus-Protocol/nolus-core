#!/bin/bash
set -euox pipefail

SCRIPT_GENESIS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_GENESIS_DIR"/internal/genesis.sh
source "$SCRIPT_GENESIS_DIR"/internal/verify.sh

cleanup() {
  cleanup_genesis_sh
  exit
}
trap cleanup INT TERM EXIT

__print_usage() {
    printf \
    "Usage: %s
    <$COMMAND_FULL_GEN>
    [-c|--chain-id <string>]
    [--currency <native_currency>]
    [--accounts <accounts_spec_json>]
    [--wasm-script-path <wasm_script_path>]
    [--wasm-code-path <wasm_code_path>]
    [--treasury-nls-u128 <init treasury amount of uNLS>
    [--validator-node-urls-pubkeys <validator_node_urls_and_validator_pubkeys>]
    [--validator-accounts-dir <validator_accounts_dir>]
    [--validator-tokens <validators_initial_tokens>]
    [--validator-stake <tokens_validator_stakes>]
    [--gov-voting-period <voting_period>]
    [--gov-max-deposit-period <max_deposit_period - XXs>]
    [--staking-max-validators <staking_max_validators>]
    [--feerefunder-ack-fee-min <feerefunder_ack_fee_min_amount>]
    [--feerefunder-timeout-fee-min <feerefunder_timeout_fee_min_amount>]
    [--dex-admin-mnemonic <dex_admin_account_mnemonic>]
    [--store-code-privileged-account-mnemonic <store_code_privileged_account_mnemonic>]
    [-o|--output <genesis_file_path>]" \
     "$1"
}

COMMAND_FULL_GEN="full-gen"
CHAIN_ID=""
NATIVE_CURRENCY="unls"
ACCOUNTS_SPEC=""
WASM_SCRIPT_PATH=""
WASM_CODE_PATH=""
TREASURY_INIT_TOKENS_U128=""
VAL_NODE_URLS_AND_VAL_PUBKEYS=""
VAL_ACCOUNTS_DIR="val-accounts"
VAL_TOKENS="1000000000""$NATIVE_CURRENCY"
VAL_STAKE="1000000""$NATIVE_CURRENCY"
OUTPUT_FILE=""
GOV_VOTING_PERIOD="43200s"
GOV_MAX_DEPOSIT_PERIOD="43200s"
STAKING_MAX_VALIDATORS="40"
FEEREFUNDER_ACK_FEE_MIN="1"
FEEREFUNDER_TIMEOUT_FEE_MIN="1"
DEX_ADMIN_MNEMONIC=""
STORE_CODE_PRIVILEGED_ACCOUNT_MNEMONIC=""
ADMINS_TOKENS="10000000""$NATIVE_CURRENCY"

if [[ $# -lt 1 ]]; then
  echo "Missing command!"
  __print_usage "$0"
  exit 1
fi
COMMAND="$1"
shift

while [[ $# -gt 0 ]]; do
  key="$1"

  case $key in

  -h | --help)
    __print_usage "$0"
    exit 0
    ;;

  -c | --chain-id)
    CHAIN_ID="$2"
    shift
    shift
    ;;

  --currency)
    NATIVE_CURRENCY="$2"
    shift
    shift
    ;;

  --accounts)
    ACCOUNTS_SPEC="$2"
    shift
    shift
    ;;

  --wasm-script-path)
    WASM_SCRIPT_PATH="$2"
    shift
    shift
    ;;

  --wasm-code-path)
    WASM_CODE_PATH="$2"
    shift
    shift
    ;;

  --treasury-nls-u128)
    TREASURY_INIT_TOKENS_U128="$2"
    shift
    shift
    ;;

  --validator-node-urls-pubkeys)
    VAL_NODE_URLS_AND_VAL_PUBKEYS="$2"
    shift
    shift
    ;;

  --validator-accounts-dir)
    VAL_ACCOUNTS_DIR="$2"
    shift
    shift
    ;;

  --validator-tokens)
    VAL_TOKENS="$2"
    shift
    shift
    ;;

  --validator-stake)
    VAL_STAKE="$2"
    shift
    shift
    ;;

  --gov-voting-period)
    GOV_VOTING_PERIOD="$2"
    shift
    shift
    ;;

  --gov-max-deposit-period)
    GOV_MAX_DEPOSIT_PERIOD="$2"
    shift
    shift
    ;;

  --staking-max-validators)
    STAKING_MAX_VALIDATORS="$2"
    shift
    shift
    ;;

  --feerefunder-ack-fee-min)
    FEEREFUNDER_ACK_FEE_MIN="$2"
    shift
    shift
    ;;

  --feerefunder-timeout-fee-min)
    FEEREFUNDER_TIMEOUT_FEE_MIN="$2"
    shift
    shift
    ;;

  --dex-admin-mnemonic)
    DEX_ADMIN_MNEMONIC="$2"
    shift
    shift
    ;;

  --store-code-privileged-account-mnemonic)
    STORE_CODE_PRIVILEGED_ACCOUNT_MNEMONIC="$2"
    shift
    shift
    ;;

  -o | --output)
    OUTPUT_FILE="$2"
    shift
    shift
    ;;

  *)
    echo "unknown option '$key'"
    exit 1
    ;;

  esac
done

if [[ "$COMMAND" == "$COMMAND_FULL_GEN" ]]; then
  verify_mandatory "$CHAIN_ID" "Nolus chain identifier"
  verify_mandatory "$ACCOUNTS_SPEC" "Nolus genesis accounts spec"
  verify_mandatory "$WASM_SCRIPT_PATH" "Wasm script path"
  verify_mandatory "$WASM_CODE_PATH" "Wasm code path"
  verify_mandatory "$TREASURY_INIT_TOKENS_U128" "Treasury init tokens"
  verify_mandatory "$VAL_NODE_URLS_AND_VAL_PUBKEYS" "Validator URLs and validator public keys spec"
  verify_mandatory "$OUTPUT_FILE" "Genesis output file"
  verify_mandatory "$DEX_ADMIN_MNEMONIC" "DEX-admin account mnemonic"
  verify_mandatory "$STORE_CODE_PRIVILEGED_ACCOUNT_MNEMONIC" "WASM store-code privileged account mnemonic"

  genesis_file=$(generate_genesis "$CHAIN_ID" "$NATIVE_CURRENCY" "$VAL_TOKENS" "$VAL_STAKE" \
                                  "$VAL_ACCOUNTS_DIR" "$ACCOUNTS_SPEC" "$WASM_SCRIPT_PATH" \
                                  "$WASM_CODE_PATH" "$TREASURY_INIT_TOKENS_U128" \
                                  "$VAL_NODE_URLS_AND_VAL_PUBKEYS" "$GOV_VOTING_PERIOD" \
                                  "$GOV_MAX_DEPOSIT_PERIOD" "$STAKING_MAX_VALIDATORS" \
                                  "$FEEREFUNDER_ACK_FEE_MIN" "$FEEREFUNDER_TIMEOUT_FEE_MIN"  \
                                  "$DEX_ADMIN_MNEMONIC" "$STORE_CODE_PRIVILEGED_ACCOUNT_MNEMONIC" \
                                  "$ADMINS_TOKENS")
  mv "$genesis_file" "$OUTPUT_FILE"
# elif [[ "$COMMAND" == "$COMMAND_SETUP" ]]; then
#
else
  echo "Unknown command!"
  exit 1
fi