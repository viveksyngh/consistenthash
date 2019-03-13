package consistenthash

//Hasher hash function type to hash the key
type Hasher func([]byte) uint32

//ConsistentHash consistent hash type
type ConsistentHash struct {
	//Hash function to hash key and nodes
	Hash Hasher
	//Replicas Number of replicas for each node
	Replicas int
	//HashRing a sorted array of hashes
	HashRing []uint32
	//HashMap a map of hash and it's value to find server
	HashMap map[uint32]string
}
