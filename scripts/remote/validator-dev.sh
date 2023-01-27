#!/bin/bash
# Configure a locally installed validator node for the needs of the dev net
# The Nolus binary should be accessible on the system path.
# TBD what part of the scripts should be available next to this script
#
# arg1: home directory of the validator node, mandatory
# arg2: node's moniker, mandatory
# arg3: base port, mandatory. Used to determine the endpoint ports.
# arg4: timeout commit, mandatory. Example: "3s".
# arg5: first node's identificator, optional. Empty, if this is the first node.
#
# Returns the node identificator in the form of "node-id@host:p2p-port" followed
# by the node public key in JSON.
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR"/lib/lib.sh
source "$SCRIPT_DIR"/../common/cmd.sh

home_dir="$1"
node_moniker="$2"
base_port="$3"
timeout_commit="$4"
if [[ $# -gt 4 ]]
then first_node_id="$5"
else first_node_id=""
fi

HOST="127.0.0.1"
P2P_PORT=$((base_port))
RPC_PORT=$((base_port+1))
MONITORING_PORT=$((base_port+2))
API_PORT=$((base_port+3))
GRPC_PORT=$((base_port+4))

rm -fr "$home_dir"
mkdir -p "$home_dir"

run_cmd "$home_dir" init "$node_moniker" >/dev/null

declare -r config_dir="$home_dir"/config

# although the API endpoint is deprecated it is still required by Keplr
# TBD reevaluate the necessity to remain open
# the grpc endpoint is required by the market data feeder
update_app "$config_dir" '."api"."enable"' "true" >/dev/null
update_app "$config_dir" '."api"."address"' '"tcp://0.0.0.0:'"$API_PORT"'"' >/dev/null
update_app "$config_dir" '."api"."enabled-unsafe-cors"' "true" >/dev/null
update_app "$config_dir" '."grpc"."enable"' "true" >/dev/null
update_app "$config_dir" '."grpc"."address"' '"0.0.0.0:'"$GRPC_PORT"'"' >/dev/null
update_app "$config_dir" '."grpc-web"."enable"' "false" >/dev/null
update_app "$config_dir" '."minimum-gas-prices"' '"'"0.0025unls"'"' >/dev/null
update_app "$config_dir" '."telemetry"."enabled"' "true" >/dev/null
update_app "$config_dir" '."telemetry"."prometheus-retention-time"' "1" >/dev/null
update_app "$config_dir" '."wasm"."query_gas_limit"' "3500000" >/dev/null

update_config "$config_dir" '."rpc"."laddr"' '"tcp://0.0.0.0:'"$RPC_PORT"'"' >/dev/null
update_config "$config_dir" '."rpc"."cors_allowed_origins"' '["*"]' >/dev/null
update_config "$config_dir" '."p2p"."laddr"' '"tcp://'"$HOST:$P2P_PORT"'"' >/dev/null
update_config "$config_dir" '."p2p"."addr_book_strict"' 'false' >/dev/null
update_config "$config_dir" '."p2p"."allow_duplicate_ip"' 'true' >/dev/null
update_config "$config_dir" '."p2p"."persistent_peers"' '"'"$first_node_id"'"' >/dev/null
update_config "$config_dir" '."proxy_app"' '""' >/dev/null
update_config "$config_dir" '."consensus"."timeout_commit"' '"'"$timeout_commit"'"' >/dev/null
update_config "$config_dir" '."instrumentation"."prometheus"' "true" >/dev/null
update_config "$config_dir" '."instrumentation"."prometheus_listen_addr"' '"'":$MONITORING_PORT"'"' >/dev/null
update_config "$config_dir" '."log_format"' '"json"'

tendermint_node_id=$(run_cmd "$home_dir" tendermint show-node-id)
validator_pub_key=$(run_cmd "$home_dir" tendermint show-validator)
echo "$tendermint_node_id@$HOST:$P2P_PORT $validator_pub_key"
