# containerChain
build a supply chain of *physical Node --> virtual Node --> pod --> container --> application --> virtualApplication*.

# Overview
<div >
<img width="700" src="https://github.com/songbinliu/containerChain/blob/master/conf/supplyChain.png">
</div>

# Commodities between layers
<div>
<img width="500" src="https://github.com/songbinliu/containerChain/blob/master/conf/commodity.png">
</div>

# Supported Actions
|SE type| Move | Resize|
|-|-|-|
|ContainerPod| Yes | No |
|Container | No | WIP |
| VirtualMachine |Yes | WIP|

note *WIP* = work in progress.

# Run it

```bash
#1. get source code
go get github.com/songbinliu/containerChain

#2. compile it
cd $GOPATH/src/github.com/songbinliu/containerChain
make build

#3. run it
turbo=./conf/turbo.json
topology=./conf/topology.conf
target=./conf/target.json
./_output/containerChain --topologyConf $topology --turboConf $turbo --targetConf $target --logtostderr --v 3 
```

**turbo** is a json file about the settings of the OpsMgr, [example](https://github.com/songbinliu/containerChain/blob/master/conf/turbo.json);

**target** is a json file about settings of generated cluster for OpsMgr, [example](https://github.com/songbinliu/containerChain/blob/master/conf/target.json);

**topology** is the configuration file about the virtual cluster to be generated, [example](https://github.com/songbinliu/containerChain/blob/master/conf/topology.conf).
