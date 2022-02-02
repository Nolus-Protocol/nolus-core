#!/bin/bash
set -euxo pipefail

# start "instance" variables
setup_validator_dev_scripts_home_dir=""
setup_validator_dev_binary_url=""
setup_validator_dev_root_dir=""
setup_validator_dev_tmp_dir=""
setup_validator_dev_prev_node_id=""
# end "instance" variables
SETUP_VALIDATOR_DEV_BASE_PORT=26606
ARTIFACT_S3_BUCKET="nolus-artifact-bucket/dev"

init_setup_validator_dev_sh() {
  setup_validator_dev_scripts_home_dir="$1"
  setup_validator_dev_binary_url="$2"
  setup_validator_dev_root_dir="$3"
  setup_validator_dev_tmp_dir=$(mktemp -d)
}

cleanup_setup_validator_dev_sh() {
  rm -fr "$setup_validator_dev_tmp_dir"
}

# Setup validator nodes and collect their ids and validator public keys
#
# The nodes are installed and configured depending on the sourced implementation script.
# The node ids and validator public keys are printed on the standard output one at a line.
setup_all() {
  set -euxo pipefail
  local validators_nb="$1"

  __deploy
  for i in $(seq "$validators_nb"); do
    config "$i"
  done
}

propagate_genesis_all() {
  local genesis_file="$1"
  local validators_nb="$2"

  for i in $(seq "$validators_nb"); do
    propagate_genesis "$i" "$genesis_file"
  done
}

#
# Return the node ids and validator public keys printed on the standard output delimited with a space.
#
config() {
  set -euxo pipefail
  local node_index="$1"

  local home_dir
  home_dir=$(__home_dir "$node_index")
  local node_moniker
  node_moniker=$(__node_moniker "$node_index")
  local node_base_port
  node_base_port=$(__node_base_port "$node_index")

  local node_id_val_pub_key
  node_id_val_pub_key=$("$setup_validator_dev_scripts_home_dir"/remote/validator-dev.sh "$home_dir" "$node_moniker" \
                                          "$node_base_port" "$setup_validator_dev_prev_node_id")
  read -r setup_validator_dev_prev_node_id __val_pub_key <<< "$node_id_val_pub_key"
  echo "$node_id_val_pub_key"
}

propagate_genesis() {
  local node_index="$1"
  local genesis_file="$2"

  cp "$genesis_file" "$(__home_dir "$node_index")/config/genesis.json"
}

#####################
# private functions #
#####################
__home_dir() {
  local node_index=$1
  local node_id
  node_id=$(__node_moniker "$node_index")
  echo "$setup_validator_dev_root_dir/$node_id"
}

__node_moniker() {
  echo "dev-validator-$1"
}

__node_base_port() {
  local node_index=$1
  echo $((SETUP_VALIDATOR_DEV_BASE_PORT + node_index*5))
}

__deploy() {
  local bin_file="nolusd.zip"
  local local_bin="$setup_validator_dev_tmp_dir"/"$bin_file"

  # certificate checks are skiped due to a potential misconfiguration of the HTTP server
  # echo | openssl s_client --showcerts  -connect  gitlab-nomo.credissimo.net:443 returns an error
  # wget as well
  # TBD enable it once the server configuration got fixed
  wget --no-check-certificate -O "$local_bin" "$setup_validator_dev_binary_url" 2>/dev/null
  aws s3 cp "$local_bin" s3://$ARTIFACT_S3_BUCKET/"$bin_file" 1>&2 2>/dev/null
}