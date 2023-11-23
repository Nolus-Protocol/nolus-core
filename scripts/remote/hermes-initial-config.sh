#!/bin/bash
# Install and Configure Hermes
#
# arg: Hermes working directory path, mandatory
# arg: chain1 id, mandatory
# arg: chain1 IP address, mandatory
# arg: chain1 RPC port, mandatory
# arg: chain1 GRPC port, mandatory
# arg: chain2 id, mandatory
# arg: chain2 IP address - rpc , mandatory
# arg: chain2 IP address - grpc, mandatory
# arg: chain2 account prefix, mandatory
# arg: chain2 price denom - grpc, mandatory
# arg: chain2 trusting period - grpc, mandatory
# arg: Hermes account seed, mandatory

set -euox pipefail

SCRIPTS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && cd .. && pwd)"
source "$SCRIPTS_DIR"/remote/lib/lib.sh
source "$SCRIPTS_DIR"/common/cmd.sh
source "$SCRIPTS_DIR"/common/rm-dir.sh
source "$SCRIPTS_DIR"/internal/add-dex-support.sh

declare -r hermes_root="$1"
declare -r chain1_id="$2"
declare -r chain1_ip_addr="$3"
declare -r chain1_rpc_port="$4"
declare -r chain1_grpc_port="$5"
declare -r chain2_id="$6"
declare -r chain2_ip_addr_RPC="$7"
declare -r chain2_ip_addr_gRPC="$8"
declare -r chain2_account_prefix="$9"
declare -r chain2_price_denom=${10}
declare -r chain2_trusting_period=${11}
declare -r hermes_mnemonic=${12}
declare if_interchain_security=${13}

# Install

declare -r archive_name="hermes-binary.tar.gz"
wget -O "$archive_name" https://github.com/informalsystems/hermes/releases/download/v1.7.0/hermes-v1.7.0-x86_64-unknown-linux-gnu.tar.gz

declare -r hermes_binary_dir="$hermes_root"/hermes
mkdir -p  "$hermes_binary_dir"
tar -C "$hermes_binary_dir" -vxzf "$archive_name"
rm "$archive_name"

# Configure

declare -r chain1_key_name="hermes-nolus"

declare -r hermes_config_dir="$HOME"/.hermes
rm_dir "$hermes_config_dir"
mkdir -p "$hermes_config_dir"
config_file="$hermes_config_dir"/config.toml
touch "$config_file"

update_config "$hermes_config_dir" '.mode.clients."enabled"' "true"
update_config "$hermes_config_dir" '.mode.clients."refresh"' "true"
update_config "$hermes_config_dir" '.mode.clients."misbehaviour"' "true"
update_config "$hermes_config_dir" '.mode.connections."enabled"' "true"
update_config "$hermes_config_dir" '.mode.channels."enabled"' "true"
update_config "$hermes_config_dir" '.mode.packets."enabled"' "true"

# Add Nolus chain configuration
update_config "$hermes_config_dir" '.chains[0]."id"' '"'"$chain1_id"'"'
update_config "$hermes_config_dir" '.chains[0]."rpc_addr"' '"http://'"$chain1_ip_addr"':'"$chain1_rpc_port"'"'
update_config "$hermes_config_dir" '.chains[0]."grpc_addr"' '"http://'"$chain1_ip_addr"':'"$chain1_grpc_port"'"'
update_config "$hermes_config_dir" '.chains[0]."rpc_timeout"' '"10s"'
update_config "$hermes_config_dir" '.chains[0]."account_prefix"' '"nolus"'
update_config "$hermes_config_dir" '.chains[0]."key_name"' '"'"$chain1_key_name"'"'
update_config "$hermes_config_dir" '.chains[0]."address_type"' '{ derivation : "cosmos" }'
update_config "$hermes_config_dir" '.chains[0]."store_prefix"' '"ibc"'
update_config "$hermes_config_dir" '.chains[0]."default_gas"' 1000000
update_config "$hermes_config_dir" '.chains[0]."max_gas"' 4000000
update_config "$hermes_config_dir" '.chains[0]."gas_price"' '{ price : 0.0025, denom : "unls" }'
update_config "$hermes_config_dir" '.chains[0]."gas_multiplier"' 1.1
update_config "$hermes_config_dir" '.chains[0]."max_msg_num"' 30
update_config "$hermes_config_dir" '.chains[0]."max_tx_size"' 2097152
update_config "$hermes_config_dir" '.chains[0]."clock_drift"' '"5s"'
update_config "$hermes_config_dir" '.chains[0]."max_block_time"' '"30s"'
update_config "$hermes_config_dir" '.chains[0]."trusting_period"' '"14days"'
update_config "$hermes_config_dir" '.chains[0]."trust_threshold"' '{ numerator : "1", denominator : "3" }'
update_config "$hermes_config_dir" '.chains[0]."memo_prefix"' '"''"'
update_config "$hermes_config_dir" '.chains[0]."event_source"' '{ mode : "push", url : "ws://127.0.0.1:'"$chain1_rpc_port"'/websocket", batch_delay : "500ms" }'

# Add DEX chain configuration
add_new_chain_hermes "$hermes_config_dir" "$chain2_id" "$chain2_ip_addr_RPC" "$chain2_ip_addr_gRPC" \
  "$chain2_account_prefix" "$chain2_price_denom" "$chain2_trusting_period" "$if_interchain_security"

# Account setup
declare hermes_mnemonic_file="$hermes_config_dir"/hermes.seed
echo "$hermes_mnemonic" > "$hermes_mnemonic_file"

dex_account_setup "$hermes_binary_dir" "$chain1_id" "$hermes_config_dir"/hermes.seed
dex_account_setup "$hermes_binary_dir" "$chain2_id" "$hermes_config_dir"/hermes.seed
