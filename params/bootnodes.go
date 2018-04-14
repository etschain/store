// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
     "enode://470a2147d6b56061bb4f785c09787911790e009086ff74e008ae2654048f5d6f08ff422a31d66bec352efd9fef76119376c1c90fa3f7f4f70fee585def2b9793@::10401",
     "enode://50faa1e6af14d0ddf2b9dc845c1b2d69a2b710c599cc3c5d609c5009014085c285b92cef25b005f34266b43c7f6295a3aed0f0cd4460f878008b7d99121f5243@::10403",
     "enode://c621385a2e0eeb0e0b9d38f577bba25547d66b3cbf951b6ee7a4c0b7f882320f7f8f0568a9268581ea1ba3542abe1c95b83def8a041be7cc4a14225d54cafade@::10401",
     "enode://02d52190e7b19d33021249fe121fdb0c32203a91b0173930c88cf1fb925ff539fb5959ab3e86e0c84a89476b87a0466dde7ab1eb2b3d826619810a0e359e3f80@::10403",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{

}
