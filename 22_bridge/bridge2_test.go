package bridge

import "testing"

func TestStringToJson(t *testing.T) {
	fileMaker := &JsonFileMaker{&StringAdapter{"this is data"}}
	fileMaker.MakeFile()
}

func TestByteToNumber(t *testing.T) {
	fileMaker := &NumberFileMaker{&ByteAdapter{[]byte("this is data")}}
	fileMaker.MakeFile()
}