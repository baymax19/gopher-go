package main
//
// import (
// 	"encoding/binary"
// 	"encoding/json"
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
//
// 	"github.com/tendermint/go-amino"
// )
//
//
//
// var _ ID = (*SubscriptionID)(nil)
//
// type SubscriptionID uint64
//
// func (id SubscriptionID) String() string {
// 	return fmt.Sprintf("sess%X", id.Uint64())
// }
//
// func (id SubscriptionID) Uint64() uint64 {
// 	return uint64(id)
// }
//
// func (id SubscriptionID) Bytes() []byte {
// 	b := make([]byte, 8)
// 	binary.LittleEndian.PutUint64(b, id.Uint64())
// 	return b
// }
//
// func (id SubscriptionID) Marshal() ([]byte, error) {
// 	return id.Bytes(), nil
// }
//
// func (id *SubscriptionID) Unmarshal(data []byte) error {
//
// 	count := binary.LittleEndian.Uint64(data)
//
// 	*id = SubscriptionID(count)
//
// 	return nil
// }
//
// func (id SubscriptionID) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(id.String())
// }
//
// func (id *SubscriptionID) UnmarshalJSON(data []byte) error {
// 	var str string
//
// 	err := json.Unmarshal(data, &str)
// 	if err != nil {
// 		return err
// 	}
//
// 	strArray := strings.Split(id.String(), "sess")
// 	*id = SubscriptionID(NewSubscriptionIDFromString(strArray[1]).Uint64())
//
// 	return nil
// }
//
// func (id SubscriptionID) IsEqual(_id ID) bool {
// 	return id.Uint64() == _id.Uint64()
// }
//
// func NewSubscriptionID(id uint64) ID {
// 	return SubscriptionID(id)
// }
//
// func NewSubscriptionIDFromString(s string) ID {
// 	i, err := strconv.ParseUint(s, 16, 64)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return SubscriptionID(i)
// }
//
// func main() {
//
// 	cdc := amino.NewCodec()
// 	cdc.RegisterInterface((*ID)(nil), nil)
//
// 	sid := NewSessionID(12)
// 	fmt.Println(sid.String())
// 	marshal, err := sid.MarshalJSON()
// 	fmt.Println("MarshalJSON", marshal, err)
// 	fmt.Println("UnmarshalJSON", sid)
//
// 	cdcBytes, _ := sid.Marshal()
//
// 	fmt.Println("cdc Marshal", cdcBytes, sid)
//
// 	fmt.Println("Sid", sid.Bytes(), sid.String())
//
// 	newSid := NewSessionID(2)
// 	fmt.Println(newSid.IsEqual(sid))
//
// 	fmt.Println("Array Testing...")
// 	s1 := NewSessionID(19)
// 	s2 := NewSessionID(12)
// 	s3 := NewSessionID(23)
// 	s4 := NewSessionID(3)
//
// 	var sids IDs
// 	sids = sids.Append(s1)
// 	sids = sids.Append(s3)
// 	sids = sids.Append(s2)
// 	sids = sids.Append(s4)
//
// 	bytes, err := cdc.MarshalBinaryBare(sid)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Println("bytes of node ", bytes)
//
// 	cdc.UnmarshalBinaryBare(bytes, &sid)
// 	fmt.Println("sid is", sid)
//
// 	fmt.Println("SessionIDs", sids)
// 	fmt.Println("SessionIDs Length", sids.Len())
// 	fmt.Println("SessionIDs Sort", sids.Sort())
// 	sidAfterDeletion := sids.Delete(2)
// 	fmt.Println("SessionIDs Deleted", sidAfterDeletion)
// 	fmt.Println("SessionIDs After Deletion Sort", sidAfterDeletion.Sort())
//
// 	// fmt.Println("Strings", sids.String())
//
// 	fmt.Println("Subsctiption IDS..")
//
// 	sub := SubscriptionID(12)
// 	fmt.Println(sid.String())
// 	marshalSub, err1 := sid.MarshalJSON()
// 	fmt.Println("MarshalJSON", marshalSub, err1)
// 	fmt.Println("UnmarshalJSON", sub)
//
// 	Bytes, _ := sub.Marshal()
//
// 	fmt.Println("cdc Marshal", Bytes, sub)
// 	fmt.Println("Sid", sub.Bytes(), sub.String())
//
// 	newSub := NewSubscriptionID(2)
// 	fmt.Println(newSub.IsEqual(sid))
//
// 	fmt.Println("Array Testing...")
// 	sub1 := NewSubscriptionID(19)
// 	sub2 := NewSubscriptionID(12)
// 	sub3 := NewSubscriptionID(23)
// 	sub4 := NewSubscriptionID(3)
//
// 	var subs IDs
//
// 	subs = subs.Append(sub1)
// 	subs = subs.Append(sub3)
// 	subs = subs.Append(sub2)
// 	subs = subs.Append(sub4)
//
// 	fmt.Println("SubscriptionIDs", subs)
// 	fmt.Println("SubscriptionIDs Length", subs.Len())
// 	fmt.Println("SubscriptionID Sort", subs.Sort())
// 	subsAfterDeletion := subs.Delete(2)
// 	fmt.Println("SubscriptionID Deleted", subsAfterDeletion)
// 	fmt.Println("SubscriptionIDs After Deletion Sort", subsAfterDeletion.Sort())
//
// 	subsBytes, err := cdc.MarshalBinaryBare(subs)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	var newSubs IDs
// 	_ = cdc.UnmarshalBinaryBare(subsBytes, &newSubs)
//
// 	fmt.Println(subsBytes, "before marsh", subs)
// 	fmt.Println("after marshal", newSubs)
//
// }
// //
// // type IDs []ID
// //
// // func (ids IDs) Append(id ID) IDs { return append(ids, id) }
// // func (ids IDs) Len() int         { return len(ids) }
// //
// // func (ids IDs) Sort() IDs {
// // 	sort.Slice(ids, func(i, j int) bool {
// // 		return ids[i].Uint64() < ids[j].Uint64()
// // 	})
// //
// // 	return ids
// // }
// //
// // func (ids IDs) Delete(index int) IDs {
// // 	ids[index] = ids[ids.Len()-1]
// // 	return ids[:ids.Len()-1]
// // }
// //
// // func (ids IDs) Search(id ID) int {
// // 	index := sort.Search(len(ids), func(i int) bool {
// // 		return ids[i].Uint64() >= id.Uint64()
// // 	})
// //
// // 	if (index == ids.Len()) ||
// // 		(index < ids.Len() && ids[index].Uint64() != id.Uint64()) {
// // 		return ids.Len()
// // 	}
// //
// // 	return index
// // }
