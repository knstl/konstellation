# Network deployment

## 1. Network files generation
The folder config should have testnet.yml in order to generate properly network files.
Network files are configs, genesis, private keys, etc, that are needed to run more than one node. 
When network consists of one node, starport chain serve can be used instead.

We have special utility cosmodrome that generates network basing on the current app modules.

Folder `config` also should have keys.json to read the mnemonics for validators.
The example of keys.json. Keys.json is also added to .gitignore to avoid hitting to a public repo.
```
{
  "keys": {
    "darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx": {
      "name": "hawking",
      "password": "hawking1",
      "mnemonic": "disorder squirrel cage garlic oyster leaf segment casual siren shiver lecture among either wool improve head thunder walnut cram force crystal advice slab sail",
      "keystore": {
        "address": "darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx",
        "crypto": {
          "cipher": "aes-128-ctr",
          "ciphertext": "07ee1d1f64f0b719b3e8b6788971d87f1f2137b349d5d047c9818cf3bc79872ffcbd50bfbd",
          "cipherparams": {
            "iv": "14037fcf066213200bf2695da591cc47"
          },
          "kdf": "scrypt",
          "kdfparams": {
            "dklen": 32,
            "n": 8192,
            "p": 1,
            "r": 8,
            "salt": "4c463a8fefbfb73e11e514a245454f42cea8e45b980120d89c71ebbc581a033c"
          },
          "mac": "bd211eab1d8bab79da1cb7b15793ed459dea2c215457e7c2bbe07093fbad9106"
        },
        "id": "93a2c265-bdbc-444e-ae56-1276f82801f0",
        "version": 3
      }
    }
  }
}
```
To generate network files run:
```
make install
cosmodrome gn --chain-id darchub -n ./config/testnet.yml -o ./testnet --keyring-backend test
```
Note: to generate localnet for using with docker, simply rename testnet into localnet

## 2. Network files delivering
Copy network node files to the appropriate server using ssh. 
Run 
```
jq -r '
    .validators[] |
    "echo \(.ip);
     ssh -i ~/.ssh/<YOUR_SSH_KEY> root@\(.ip) \"rm -rdf /root/.knstld\";
     scp -i ~/.ssh/<YOUR_SSH_KEY> -r ./testnet/\(.name)/.knstld root@\(.ip):/root;
     echo "
    ' ./config/testnet.json
```
## 3. Network launching
Login to each server that will run node and run
```
knstld start
```
As far as blockchain requires 2/3 nodes to support consensus, 
first nodes will wait for the 2/3+1 node to start blockchain.

Assuming that you have 3 nodes in total, first 2 will wait for the third one. And after that consensus will initiated