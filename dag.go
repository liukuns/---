package merkledag

import "hash"

func Add(store KVStore, node Node, h hash.Hash) []byte {
	// TODO 将分片写入到KVStore中，并返回Merkle Root
	return nil
}

// package merkledag

// import (
// 	"hash"
// )

// func Add(store KVStore, node Node, h hash.Hash) []byte {
// 	switch node.Type() {
// 	case FILE:
// 		// 对于文件，我们直接将其内容保存到KVStore中
// 		file := node.(File)
// 		content := file.Bytes()
// 		h.Write(content)
// 		hash := h.Sum(nil)
// 		store.Put(hash, content)
// 		return hash
// 	case DIR:
// 		// 对于目录，我们需要遍历其下的所有节点，并递归地调用Add函数
// 		dir := node.(Dir)
// 		it := dir.It()
// 		var hashes []byte
// 		for it.Next() {
// 			hash := Add(store, it.Node(), h)
// 			hashes = append(hashes, hash...)
// 		}
// 		// 最后，我们将所有子节点的哈希值连接起来，并计算其哈希值
// 		h.Reset()
// 		h.Write(hashes)
// 		hash := h.Sum(nil)
// 		store.Put(hash, hashes)
// 		return hash
// 	default:
// 		// 未知的节点类型
// 		return nil
// 	}
// }



package merkledag

import (
	"hash"
)

// 定义一个新的哈希结构
type MyHash struct {
	data []byte
}

// 为新的哈希结构实现 hash.Hash 接口的所有方法
func (h *MyHash) Write(p []byte) (n int, err error) {
	h.data = append(h.data, p...)
	return len(p), nil
}

func (h *MyHash) Sum(b []byte) []byte {
	return append(b, h.data...)
}

func (h *MyHash) Reset() {
	h.data = nil
}

func (h *MyHash) Size() int {
	return len(h.data)
}

func (h *MyHash) BlockSize() int {
	return 1
}

// 在 Add 函数中使用新的哈希结构
func Add(store KVStore, node Node) []byte {
	h := &MyHash{}
	switch node.Type() {
	case FILE:
		// 对于文件，我们直接将其内容保存到KVStore中
		file := node.(File)
		content := file.Bytes()
		h.Write(content)
		hash := h.Sum(nil)
		store.Put(hash, content)
		return hash
	case DIR:
		// 对于目录，我们需要遍历其下的所有节点，并递归地调用Add函数
		dir := node.(Dir)
		it := dir.It()
		var hashes []byte
		for it.Next() {
			hash := Add(store, it.Node())
			hashes = append(hashes, hash...)
		}
		// 最后，我们将所有子节点的哈希值连接起来，并计算其哈希值
		h.Reset()
		h.Write(hashes)
		hash := h.Sum(nil)
		store.Put(hash, hashes)
		return hash
	default:
		// 未知的节点类型
		return nil
	}
}
