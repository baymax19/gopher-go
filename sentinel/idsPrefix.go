package main

// //
// // import (
// // 	"fmt"
// // )
// //
// // // type ID uint64
// // //
// // // func (id ID) String(prefix string) string {
// // //
// // // 	bech32ID, err := bech32.ConvertAndEncode(prefix, id.Bytes())
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}
// // // 	// bech32ID = prefix + fmt.Sprintf("%X", id)
// // //
// // // 	return bech32ID
// // // }
// // //
// // // func (id ID) Bytes() []byte {
// // // 	fmt.Println(id)
// // // 	return []byte(fmt.Sprintf("%X", id))
// // // }
// //
// // type ID uint64
// //
// // type Prefix interface {
// // 	String() string
// // }
// //
// // var _ Prefix = (*Node)(nil)
// //
// // type Node struct {
// // 	ID ID
// // }
// //
// // func (n Node) String() string {
// // 	return "Node" + fmt.Sprintf("%X", n.ID)
// // }
// //
// // type Sessions struct {
// // 	ID uint64
// // }
// //
// // func (s Sessions) String() string {
// // 	return "Session" + fmt.Sprintf("%X", s.ID)
// // }
// //
// // func main() {
// // 	var a Node
// // 	a.ID = 12
// // 	fmt.Println(a.String())
// //
// // 	var s Sessions
// // 	s.ID = 12
// // 	fmt.Println(s.String())
// // }


// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// )

// type BaseID uint64

// func NewIDFromUInt64(i uint64) BaseID { return BaseID(i) }

// func NewIDFromString(s string) BaseID {
// 	i, err := strconv.ParseUint(s, 16, 64)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	return NewIDFromUInt64(i)
// }

// type SessionID uint64
// type SubscriptionID uint64
// type NodeID uint64

// type ID interface {
// 	String() string
// 	Uint64() uint64
// 	IsEqual(id ID) bool
// }

// var _ ID = (*SessionID)(nil)

// func (id SessionID) String() string      { return "session" + fmt.Sprintf("%X", id.Uint64()) }
// func (id SessionID) Uint64() uint64      { return uint64(id) }
// func (id SessionID) IsEqual(_id ID) bool { return id.Uint64() == _id.Uint64() }

// var _ ID = (*SubscriptionID)(nil)

// func (id SubscriptionID) String() string      { return "subscription" + fmt.Sprintf("%X", id) }
// func (id SubscriptionID) Uint64() uint64      { return uint64(id) }
// func (id SubscriptionID) IsEqual(_id ID) bool { return id.Uint64() == _id.Uint64() }

// var _ ID = (*NodeID)(nil)

// func (id NodeID) String() string      { return "session" + fmt.Sprintf("%X", id) }
// func (id NodeID) Uint64() uint64      { return uint64(id) }
// func (id NodeID) IsEqual(_id ID) bool { return id.Uint64() == _id.Uint64() }

// type IDs []BaseID

// func (ids IDs) Append(id ...BaseID) IDs { return append(ids, id...) }
// func (ids IDs) Len() int                { return len(ids) }

// func (ids IDs) Sort() IDs {
// 	sort.Slice(ids, func(i, j int) bool {
// 		return ids[i] < ids[j]
// 	})
	
// 	return ids
// }

// func (ids IDs) Delete(index int) IDs {
// 	ids[index] = ids[ids.Len()-1]
// 	return ids[:ids.Len()-1]
// }

// func (ids IDs) Search(id BaseID) int {
// 	index := sort.Search(len(ids), func(i int) bool {
// 		return ids[i] >= id
// 	})
	
// 	if (index == ids.Len()) ||
// 		(index < ids.Len() && ids[index] != id) {
// 		return ids.Len()
// 	}
	
// 	return index
// }
