# Greenfield BSC Archiver

## Description

The Greenfield BSC Archiver is tasked with storing headers and bodies of BSC historical blocks. It uploads this block information to Greenfield via the bundle service. The components of this archiver include:

### Syncer

Responsible for fetching block information via Ethereum calls, storing it in a database, and uploading it to Greenfield through the bundle service. The current setting uploads data every 100 blocks.

### Verifier

Regularly checks the integrity of uploaded data on Greenfield. If errors are found, it corrects and re-uploads the data, and updates the block bundle name in the database.

### Server

Provides RPC interfaces compatible with Ethereum calls. Variables are generated using a Swagger.yaml file to ensure adaptability.
1. eth_getBlockByNumber
2. eth_getBundledBlockByNumber
3. eth_blockNumber
4. eth_getBlockByHash
5. /bsc/v1/blocks/{block_id}/bundle/name

### Example

**eth_getBlockByNumber**
request
```json
{
    "method": "eth_getBlockByNumber",
    "params": [
        "0x2625A00"
    ],
    "id": 100,
    "jsonrpc": "2.0"
}
```


response
```json
{
    "id": 100,
    "jsonrpc": "2.0",
    "result": {
        "difficulty": "0x2",
        "extraData": "0xd883010002846765746888676f312e31332e34856c696e757800000000000000924cd67a1565fdd24dd59327a298f1d702d6b7a721440c063713cecb7229f4e162ae38be78f6f71aa5badeaaef35cea25061ee2100622a4a1631a07e862b517401",
        "gasLimit": "0x25ff7a7",
        "gasUsed": "0x300b37",
        "hash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
        "logsBloom": "0x08000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000",
        "miner": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
        "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
        "nonce": "0x0000000000000000",
        "number": "0x1",
        "parentHash": "0x0d21840abff46b96c84b2ac9e10e4f5cdaeb5693cb665db62a2f3b02d2d57b5b",
        "receiptsRoot": "0xfc7c0fda97e67ed8ae06e7a160218b3df995560dfcb209a3b0dddde969ec6b00",
        "sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
        "size": "0x558",
        "stateRoot": "0x1db428ea79cb2e8cc233ae7f4db7c3567adfcb699af668a9f583fdae98e95588",
        "timestamp": "0x5f49ca59",
        "totalDifficulty": "0x3",
        "transactions": [
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0xbaf8ffa0b475a67cfeac3992d24422804452f0982e4e21a8816db2e0c9e5f224",
                "input": "0xe1c7392a",
                "nonce": "0x0",
                "r": "0xba64c6614028e6f7fa4cf2b218719160b9850c591279218def6894586ab73157",
                "s": "0x69cd1b3041eaa9b282a522c49b294e3c51801a1347bb8999015a0bd29bed3bbf",
                "to": "0x0000000000000000000000000000000000001000",
                "transactionIndex": "0x0",
                "type": "0x0",
                "v": "0x93",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0x8ea486df4eafaf713fbbe3b4b0b4196e50fbd1ea93daf66675accf3bf3f59d00",
                "input": "0xe1c7392a",
                "nonce": "0x1",
                "r": "0xd6541593c8d1bce8024f2bccfa6f73fc3a107f08e02b39a4571a027863558b3b",
                "s": "0x3448e68c416c115137e9cd1b04a9a09e10a53eadd958d9606962b4170ffffbe7",
                "to": "0x0000000000000000000000000000000000001001",
                "transactionIndex": "0x1",
                "type": "0x0",
                "v": "0x93",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0x9ebc5237eabb339a103a34daf280db3d9498142b49fa47f1af71f64a605acffa",
                "input": "0xe1c7392a",
                "nonce": "0x2",
                "r": "0x151f5a54ed865ed852eafe84be6557c13c6d23535ff0ec212e043f48ba5c275d",
                "s": "0x17fa5ca73322133191df498ecb7a9b7fa57a81c403312704314716d81666f848",
                "to": "0x0000000000000000000000000000000000001003",
                "transactionIndex": "0x2",
                "type": "0x0",
                "v": "0x93",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0xc043c5d33f8c3a6d6c0853ff8cbe88ebdf746f8092cb763b18de65db45246a6e",
                "input": "0xe1c7392a",
                "nonce": "0x3",
                "r": "0xf1d2414ecd63f567980be648645197d8a0d940921e6e8cde82f9da0ba7259be",
                "s": "0x5add78f836ad168f37cf8fe38432ca01c7827ffbccb8b1981b151f4463651b90",
                "to": "0x0000000000000000000000000000000000001006",
                "transactionIndex": "0x3",
                "type": "0x0",
                "v": "0x94",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0x2f64d7e926e6fb62f906e18258097af179c213f0c87a717476cce1b334049797",
                "input": "0xe1c7392a",
                "nonce": "0x4",
                "r": "0x51c8f79cbe3eabe7b8c90fff02b6186c26a31681acea887557450916eaf7b67b",
                "s": "0x63d36286391f2976842e9f541a0ab07aa7e662b79b420f78e59a1e9a487658c9",
                "to": "0x0000000000000000000000000000000000001004",
                "transactionIndex": "0x4",
                "type": "0x0",
                "v": "0x93",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0x463f0a179a89f47b055df14897dd7c55a2d819351568045dcb0496f2875c71ee",
                "input": "0xe1c7392a",
                "nonce": "0x5",
                "r": "0x36ffe2a7b51ea50b94dc75e02171e494e1e876902371586f24708a8c8ddb1b1b",
                "s": "0x3164d01a18795962b3a317dfec2ed4393dc0a6c34a2dc2bbb0a986d348f134df",
                "to": "0x0000000000000000000000000000000000001005",
                "transactionIndex": "0x5",
                "type": "0x0",
                "v": "0x93",
                "value": "0x0"
            },
            {
                "accessList": null,
                "blockHash": "0x04055304e432294a65ff31069c4d3092ff8b58f009cdb50eba5351e0332ad0f6",
                "blockNumber": "0x1",
                "chainId": "0x38",
                "from": "0x2a7cdd959bFe8D9487B2a43B33565295a698F7e2",
                "gas": "0x7fffffffffffffff",
                "gasPrice": "0x0",
                "hash": "0xc02fd5fc71fe8bdc4fec3f97a019a4dc9961eb95e5251c55fcb3da76f5cb5bca",
                "input": "0xe1c7392a",
                "nonce": "0x6",
                "r": "0x410b03142901b97ff5c61ed962de42f2aa052799eb1d196a7c1a2aca839e012a",
                "s": "0x599d8118ea6c2bbfe6a2338d21b46667b6ebd7d698221c8c3c580b59f3d49a06",
                "to": "0x0000000000000000000000000000000000002000",
                "transactionIndex": "0x6",
                "type": "0x0",
                "v": "0x94",
                "value": "0x0"
            }
        ],
        "transactionsRoot": "0x53a8743b873570daa630948b1858eaf5dc9bb0bca2093a197e507b2466c110a0",
        "uncles": [],
        "withdrawals": null
    }
}
```

### Get a Bundler Account

Request a bundle account from the Bundle Service, you need to grant the bundle account permission in next step, so that bundle service
can create object behave of your account.

```shell
curl -X POST  https://gnfd-mainnet-bundle.nodereal.io/v1/bundlerAccount/0xf74d8897D8BeafDF4b766E19A62078DE84570656

{"address":"0x4605BFc98E0a5EA63D9D5a4a1Df549732a6963f3"}
```

### Grant fee and permission to the bundle address for creating bundled objects under the bucket


#### use go-sdk

```go
package main

import (
    "context"
    "log"
    "testing"

    "cosmossdk.io/math"
    "github.com/bnb-chain/greenfield-go-sdk/client"
    "github.com/bnb-chain/greenfield-go-sdk/pkg/utils"
    types3 "github.com/bnb-chain/greenfield-go-sdk/types"
    types4 "github.com/bnb-chain/greenfield/sdk/types"
    "github.com/bnb-chain/greenfield/x/permission/types"
    types2 "github.com/cosmos/cosmos-sdk/types"
)

func TestClient_ListBuckets(t *testing.T) {
    defaultAccount, _ := types3.NewAccountFromPrivateKey("barry", "private-key")
    cli, _ := client.New("greenfield_5600-1", "https://gnfd-testnet-fullnode-tendermint-ap.bnbchain.org:443", client.Option{
    DefaultAccount: defaultAccount,
    })
    bucketActions := []types.ActionType{types.ACTION_CREATE_OBJECT}
    statements := utils.NewStatement(bucketActions, types.EFFECT_ALLOW, nil, types3.NewStatementOptions{
    StatementExpireTime: nil,
    LimitSize:           0,
    })
    bundleAgentPrincipal, err := utils.NewPrincipalWithAccount(types2.MustAccAddressFromHex("0x4605BFc98E0a5EA63D9D5a4a1Df549732a6963f3"))
    if err != nil {
    log.Fatalf("NewPrincipalWithAccount: %v", err)
    return
    }

    _, err = cli.PutBucketPolicy(context.Background(), "your-bucket-name", bundleAgentPrincipal, []*types.Statement{&statements}, types3.PutPolicyOption{})
    if err != nil {
    log.Fatalf("put policy failed: %v", err)
    return
    }

    allowanceAmount := math.NewIntWithDecimal(1, 19)
    _, err = cli.GrantBasicAllowance(context.Background(), "granteeAddress", allowanceAmount, nil, types4.TxOption{})
    if err != nil {
    log.Fatalf("grant fee allowance failed: %v", err)
    }
}

```

You can find similar example [permission](https://github.com/bnb-chain/greenfield-go-sdk/blob/master/examples/permission.go):

#### Use provided sript
or you can use the script, replace `GRANTEE_BUNDLE_ACCOUNT` with the addr got from Bundle Service in step 2, and modify the `ALLOWANCE` to your expect amount, so that transacion
gas will be paid your account:

```shell
bash scripts/set_up.sh --grant
```

After above steps are done, you can start running the Greenfield BSC Archiver.



## Build

### Build all

```shell
make build
```

### Build blob syncer

```shell
make build_syncer
```

### Build Blob Hub api server

```shell
make build_server
```


### Run the api server

```shell
./build/server --config-path config/local/config-server.json --port 8080
```