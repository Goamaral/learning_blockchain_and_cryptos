package hashlist

import (
  "crypto/sha256"
  "node"
)

type Hashlist struct {
  Checksum [sha256.Size]byte
  Root *node.Node
}

func New(data []byte) *Hashlist {
  length := len(data)
  i := 0
  var list Hashlist
  var hashChain []byte
  var tailNode *node.Node

  for i < length {
    upper := i + 32
    if upper > length {
      upper = length
    }

    node := node.New(data[i:upper])

    if list.Root == nil {
      list.Root = node
    } else {
      tailNode.Next = node
    }

    tailNode = node
    hashChain = append(hashChain, node.Checksum[:]...)

    i += 32
  }

  list.Checksum = sha256.Sum256(hashChain)

  return &list
}