set -e

# Use testnet for deployment with legacy transaction format
forge create --rpc-url https://testrpc.xlayer.tech --private-key $PRIVATE_KEY --legacy --broadcast contracts/SmartAccount.sol:SmartAccount --constructor-args $OWNER
