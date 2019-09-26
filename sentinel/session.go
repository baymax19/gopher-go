package main

import (
	"fmt"
	"strings"
)

//
// type ID interface {
// 	String() string
// 	Uint64() uint64
// 	Bytes() []byte
// 	Marshal() ([]byte, error)
// 	MarshalJSON() ([]byte, error)
// 	IsEqual(ID) bool
// }
//
// type SessionID []byte
//
// func (id SessionID) String() string {
// 	return fmt.Sprintf("sess%X", id.Uint64())
// }
//
// func (id SessionID) Uint64() uint64 {
// 	return binary.LittleEndian.Uint64(id.Bytes())
// }
//
// func (id SessionID) Bytes() []byte {
// 	return id
// }
//
// func (id SessionID) Marshal() ([]byte, error) {
// 	return id.Bytes(), nil
// }
//
// func (id *SessionID) Unmarshal(data []byte) error {
// 	*id = data
// 	return nil
// }
//
// func (id SessionID) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(id.String())
// }
//
// func (id *SessionID) UnmarshalJSON(data []byte) error {
// 	var str string
// 	if err := json.Unmarshal(data, &str); err != nil {
// 		panic(err)
// 	}
//
// 	// fmt.Println("bytes ", string(data))
// 	//
// 	strArray := strings.Split(str, "sess")
// 	// fmt.Println(strArray[1])
// 	// uinte := binary.LittleEndian.Uint64([]byte(strArray[1]))
// 	// fmt.Println("unite", uinte)
// 	//
//
// 	*id = SessionID([]byte(strArray[1]))
// 	fmt.Println("id", id)
// 	return nil
// }
//
// func (id SessionID) IsEqual(_id ID) bool {
// 	return id.Uint64() == _id.Uint64()
// }
//
// func NewSessionID(id uint64) SessionID {
// 	b := make([]byte, 8)
// 	binary.LittleEndian.PutUint64(b, id)
//
// 	return SessionID(b)
// }
//
// func NewSessionIDFromString(s string) SessionID {
// 	i, err := strconv.ParseUint(s, 16, 64)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return NewSessionID(i)
// }
//
// func main() {
// 	sid1 := NewSessionID(170000008)
//
// 	fmt.Println("SessionID", sid1)
// 	fmt.Println("SessionID String", sid1.String())
// 	fmt.Println("SessionID Bytes", sid1.Bytes())
// 	fmt.Println("SesssionID Uint64", sid1.Uint64())
// 	marshal, err := sid1.Marshal()
// 	fmt.Println("SesssionID Marshal", marshal, err)
//
// 	var newSid1 SessionID
//
// 	err = newSid1.Unmarshal(marshal)
// 	fmt.Println("UnMarshal", newSid1, err)
//
// 	fmt.Println("SessionID Equals", newSid1.IsEqual(NewSessionID(12)))
//
// 	marshalJSON, err := newSid1.MarshalJSON()
// 	fmt.Println("Marshal JSON", marshalJSON, err)
//
// 	var newSid2 SessionID
// 	err = newSid2.UnmarshalJSON(marshalJSON)
// 	fmt.Println("UnmarshalJSON", newSid2, err)
// }

func main() {
	fmt.Println("Equal", Equal("a", "A"))
}

func Equal(a, b string) bool {
	return strings.Compare(a, b) < 0
}
