# BlockChain using Go

Pre-requisite is

```bash
mkdir -p ./tmp/blocks
```

# Commands

## createblockchain

```sh
$ go run main.go createblockchain -address "Vinamra"
2021/04/01 11:36:29 Replaying from value pointer: {Fid:0 Len:0 Offset:0}
2021/04/01 11:36:29 Iterating file id: 0
2021/04/01 11:36:29 Iteration took: 13.253µs
00000c7149982119142803dba19f29f0d51c4cdd400408c83fcfb6996d250288
Genesis created
Finished
```

## printchain

```sh
$ go run main.go printchain
2021/04/01 11:37:14 Replaying from value pointer: {Fid:0 Len:42 Offset:590}
2021/04/01 11:37:14 Iterating file id: 0
2021/04/01 11:37:14 Iteration took: 11.409µs
Previous Hash:
Hash: 00000c7149982119142803dba19f29f0d51c4cdd400408c83fcfb6996d250288
PoW: true
```

## getbalance

```sh
$ go run main.go getbalance -address "Vinamra"
2021/04/01 11:37:40 Replaying from value pointer: {Fid:0 Len:42 Offset:590}
2021/04/01 11:37:40 Iterating file id: 0
2021/04/01 11:37:40 Iteration took: 10.277µs
Balance of Vinamra: 100
```

## send

```sh
$ go run main.go send -from "Vinamra" -to "Fred" -amount 50
2021/04/01 11:39:57 Replaying from value pointer: {Fid:0 Len:42 Offset:590}
2021/04/01 11:39:57 Iterating file id: 0
2021/04/01 11:39:57 Iteration took: 15.403µs
00003e6e91de63f3b169b855d96e2360f30a19a630e87c6315c18f9b29b09bd9
Success
```
