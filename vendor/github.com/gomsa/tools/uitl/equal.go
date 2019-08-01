package uitl

import "encoding/json"

// Data2Data 两个结构之间数据架构相同进行数据转换赋值
// type User struct {
// 	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
// 	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
// }
// to
// type User struct {
// 	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
// 	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
// }
func Data2Data(data interface{}, newData interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(j, newData)
	if err != nil {
		return err
	}
	return nil
}
