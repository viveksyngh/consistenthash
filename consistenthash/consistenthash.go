package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

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
	//HashMap a map of hash and it's value
	HashMap map[uint32]string
}

//New returns a instance of consistent hash
func New(replicas int, hasher Hasher) *ConsistentHash {
	chash := &ConsistentHash{
		Hash:     hasher,
		Replicas: replicas,
		HashMap:  make(map[uint32]string),
	}

	//Using CRC32 hash becuase It has pretty good peformance compared other
	//https://softwareengineering.stackexchange.com/questions/49550/which-hashing-algorithm-is-best-for-uniqueness-and-speed
	if chash.Hash == nil {
		chash.Hash = crc32.ChecksumIEEE
	}

	return chash
}

//Add add keys to consistent hashing
func (c *ConsistentHash) Add(keys ...string) {
	for _, key := range keys {

		//Create replicas of nodes to uniform distribution of keys
		for i := 0; i < c.Replicas; i++ {
			hash := c.Hash([]byte(key + strconv.Itoa(i)))
			c.HashRing = append(c.HashRing, hash)
			c.HashMap[hash] = key
		}
	}

	//Sort hash keys on the hashring
	sort.Slice(c.HashRing, func(i, j int) bool {
		return c.HashRing[i] < c.HashRing[j]
	})
}

//IsEmpty checks if hash map is empty or not
func (c *ConsistentHash) IsEmpty() bool {
	if len(c.HashRing) == 0 {
		return true
	}

	return false
}

//Get get nearest server to store
func (c *ConsistentHash) Get(key string) string {
	if c.IsEmpty() {
		return ""
	}

	hash := c.Hash([]byte(key))

	for _, item := range c.HashRing {
		if item >= hash {
			return c.HashMap[item]
		}
	}

	return c.HashMap[c.HashRing[0]]
}
