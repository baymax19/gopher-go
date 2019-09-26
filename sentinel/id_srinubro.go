package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	
	"github.com/tendermint/go-amino"
)

const (
	NodeIDPrefix       = "node"
	SessionIDPrefix    = "sess"
	SubScriptionPrefix = "subs"
)

type ID interface {
	String() string
	Uint64() uint64
	Bytes() []byte
	IsEqual(ID) bool
	Marshal() ([]byte, error)
	MarshalJSON() ([]byte, error)
}

// NodeID Methods

type NodeID []byte

func (id NodeID) String() string { return fmt.Sprintf("%s%x", NodeIDPrefix, id.Uint64()) }

func (id NodeID) Uint64() uint64 { return binary.BigEndian.Uint64(id) }

func (id NodeID) Bytes() []byte { return id }

func (id NodeID) IsEqual(_id ID) bool { return id.String() == _id.String() }

func (id NodeID) MarshalJSON() ([]byte, error) { return json.Marshal(id.String()) }

func (id *NodeID) UnmarshalJSON(bytes []byte) error {
	
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	
	if len(str) < 5 {
		return fmt.Errorf("invalid node id length")
	}
	
	*id = NewNodeIDFromString(str[4:])
	return nil
}

func NewNodeID(id uint64) NodeID {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, id)
	
	return bytes
}

func NewNodeIDFromString(id string) NodeID {
	_id, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		panic(err)
	}
	
	return NewNodeID(_id)
}

type SessionID []byte

func (id SessionID) String() string {
	return fmt.Sprintf("%s%x", SessionIDPrefix, id.Uint64())
}

func (id SessionID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id SessionID) Bytes() []byte {
	return id
}

func (id SessionID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id SessionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *SessionID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	
	if len(s) < 5 {
		return fmt.Errorf("invalid session id length")
	}
	
	*id = NewSessionIDFromString(s[4:])
	return nil
}

func (id SessionID) Marshal() ([]byte, error) {
	return id, nil
}

func (id *SessionID) Unmarshal(data []byte) error {
	*id = data
	return nil
}

func NewSessionID(id uint64) SessionID {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, id)
	
	return SessionID(b)
}

func NewSessionIDFromString(id string) SessionID {
	_id, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		panic(err)
	}
	
	return NewSessionID(_id)
}

type SubscriptionID []byte

func (id SubscriptionID) String() string {
	return fmt.Sprintf("%s%x", SubScriptionPrefix, id.Uint64())
}

func (id SubscriptionID) Uint64() uint64 {
	return binary.LittleEndian.Uint64(id)
}

func (id SubscriptionID) Bytes() []byte {
	return id
}

func (id SubscriptionID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id SubscriptionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *SubscriptionID) UnmarshalJSON(bytes []byte) error {
	
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	
	if len(str) < 5 {
		return fmt.Errorf("invalid subscription id length")
	}
	
	*id = NewSubscriptionIDFromString(str[4:])
	return nil
}

func (id SubscriptionID) Marshal() ([]byte, error) {
	return id, nil
}

func (id *SubscriptionID) Unmarshal(data []byte) error {
	*id = data
	return nil
}

func NewSubscriptionID(id uint64) SubscriptionID {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, id)
	
	return bytes
}

func NewSubscriptionIDFromString(id string) SubscriptionID {
	_id, err := strconv.ParseUint(id, 16, 64)
	if err != nil {
		panic(err)
	}
	
	return NewSubscriptionID(_id)
}

var _ sort.Interface = (*IDs)(nil)

type IDs []ID

func (ids IDs) Append(id ID) IDs { return append(ids, id) }

func (ids IDs) Len() int {
	return len(ids)
}

func (ids IDs) Less(x, y int) bool {
	return ids[x].Uint64() < ids[y].Uint64() // TODO check with bro
}

func (ids IDs) Swap(x, y int) {
	ids[x], ids[y] = ids[y], ids[x]
}

func (ids IDs) Sort() IDs {
	sort.Slice(ids, ids.Less)
	return ids
}

func (ids IDs) Delete(x int) IDs {
	ids[x] = ids[ids.Len()-1]
	return ids[:ids.Len()-1]
}

func (ids IDs) Search(_ids ID) int {
	s := _ids.Uint64()
	index := sort.Search(len(ids), func(x int) bool {
		return ids[x].Uint64() >= s
	})
	
	if (index == ids.Len()) ||
		(index < ids.Len() && ids[index].Uint64() != s) {
		return ids.Len()
	}
	
	return index
}

func main() {
	sid1 := NewSessionID(170000008)
	
	fmt.Println("SessionID", sid1)
	fmt.Println("SessionID String", sid1.String())
	fmt.Println("SessionID Bytes", sid1.Bytes())
	fmt.Println("SesssionID Uint64", sid1.Uint64())
	marshal, err := sid1.Marshal()
	fmt.Println("SesssionID Marshal", marshal, err)
	
	var newSid1 SessionID
	
	err = newSid1.Unmarshal(marshal)
	fmt.Println("SessionID UnMarshal", newSid1, err)
	
	fmt.Println("SessionID Equals", newSid1.IsEqual(NewSessionID(12)))
	
	marshalJSON, err := newSid1.MarshalJSON()
	fmt.Println("SessionID Marshal JSON", marshalJSON, err)
	
	var newSid2 SessionID
	err = newSid2.UnmarshalJSON(marshalJSON)
	fmt.Println("SessionID UnmarshalJSON", newSid2, err)
	
	fmt.Println("Slice Operations...")
	
	sess1 := NewSessionID(1)
	sess2 := NewSessionID(13)
	sess3 := NewSessionID(4)
	sess4 := NewSessionID(19)
	
	var sessSlice IDs
	
	sessSlice = sessSlice.Append(sess1)
	sessSlice = sessSlice.Append(sess2)
	sessSlice = sessSlice.Append(sess3)
	sessSlice = sessSlice.Append(sess4)
	
	fmt.Println("SessionSlice", sessSlice)
	fmt.Println("SessionSlice Len", sessSlice.Len())
	
	fmt.Println("SessionSlice Sort", sessSlice.Sort())
	fmt.Println("SessionSlice Search", sessSlice.Search(sess2))
	
	newSlice := sessSlice.Delete(2)
	fmt.Println("newSessionSlice", newSlice)
	
	fmt.Println("newSessionSlice Sort", newSlice.Sort())
	fmt.Println("newSessionSlice Search", newSlice.Search(sess2))
	
	// Subscriptin
	
	sub1 := NewSubscriptionID(170000008)
	
	fmt.Println("Subscription", sub1)
	fmt.Println("Subscription String", sub1.String())
	fmt.Println("Subscription Bytes", sub1.Bytes())
	fmt.Println("Subscription Uint64", sub1.Uint64())
	marshal, err = sub1.Marshal()
	fmt.Println("Subscription Marshal", marshal, err)
	
	var newSub1 SubscriptionID
	
	err = newSub1.Unmarshal(marshal)
	fmt.Println("Subscription UnMarshal", newSub1, err)
	
	fmt.Println("Subscription Equals", newSub1.IsEqual(NewSubscriptionID(170000008)))
	
	marshalJSON, err = newSub1.MarshalJSON()
	fmt.Println("Subscription Marshal JSON", marshalJSON, err)
	
	var newSub2 SubscriptionID
	err = newSub2.UnmarshalJSON(marshalJSON)
	fmt.Println("Subscription UnmarshalJSON", newSub2, err)
	
	fmt.Println("Slice Operations...")
	
	sub1 = NewSubscriptionID(19)
	sub2 := NewSubscriptionID(13)
	sub3 := NewSubscriptionID(14)
	sub4 := NewSubscriptionID(19)
	
	var subSlice IDs
	
	subSlice = subSlice.Append(sub1)
	subSlice = subSlice.Append(sub2)
	subSlice = subSlice.Append(sub3)
	subSlice = subSlice.Append(sub4)
	
	fmt.Println("SessionSlice", subSlice)
	fmt.Println("SessionSlice Len", subSlice.Len())
	
	fmt.Println("SessionSlice Sort", subSlice.Sort())
	fmt.Println("SessionSlice Search", subSlice.Search(sub1)) // TODO Searching
	
	newSubSlice := subSlice.Delete(2)
	fmt.Println("newSessionSlice", newSubSlice)
	
	fmt.Println("newSessionSlice Sort", newSubSlice.Sort())
	fmt.Println("newSessionSlice Search", newSubSlice.Search(sub2))
	
	// Codec Marshal
	
	cdc := amino.NewCodec()
	cdc.RegisterInterface((*ID)(nil), nil)
	cdc.RegisterConcrete(SessionID{}, "sessionID", nil)
	cdc.RegisterConcrete(SubscriptionID{}, "subID", nil)
	
	bz, err := cdc.MarshalBinaryBare(sessSlice)
	if err != nil {
		panic(err)
	}
	fmt.Println(bz)
	
	var newSesslice IDs
	err = cdc.UnmarshalBinaryBare(bz, &newSesslice)
	fmt.Println("newsessSlice", newSesslice)
	
	bz, err = cdc.MarshalBinaryBare(newSubSlice)
	if err != nil {
		panic(err)
	}
	fmt.Println(bz)
	
	var newSubSlice1 IDs
	err = cdc.UnmarshalBinaryBare(bz, &newSubSlice1)
	fmt.Println("newSubSlice", newSubSlice1)
	
}
