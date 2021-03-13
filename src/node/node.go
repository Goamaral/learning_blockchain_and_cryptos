package node

import "crypto/sha256"

type Node struct {
  Checksum [sha256.Size]byte
  Next *Node
}

func New(data []byte) *Node {
  checksum := sha256.Sum256(data[:32])
  return &Node{Checksum: checksum, Next: nil}
}