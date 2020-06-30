package hashtable

type HashTable interface {
	Insert(value int)
	Lookup(value int) bool
}

type linearList interface {
	Insert(value int)
	Lookup(value int) bool
}

type linearListImpl struct {
	values []int
}

func (list *linearListImpl) Insert(value int) {
	list.values = append(list.values, value)
}

func (list *linearListImpl) Lookup(value int) bool {
	for _, v := range list.values {
		if v == value {
			return true
		}
	}
	return false
}

const tableSize = 15

type hashTableImpl struct {
	table [tableSize]linearList
}

func New() HashTable {
	table := [tableSize]linearList{}
	for i := 0; i < tableSize; i++ {
		table[i] = &linearListImpl{}
	}
	return &hashTableImpl{table: table}
}

func (table *hashTableImpl) Insert(value int) {
	i := table.hashFunction(value)
	table.table[i].Insert(value)
}

func (table *hashTableImpl) Lookup(value int) bool {
	i := table.hashFunction(value)
	return table.table[i].Lookup(value)
}

func (table *hashTableImpl) hashFunction(value int) int {
	return value % tableSize
}
