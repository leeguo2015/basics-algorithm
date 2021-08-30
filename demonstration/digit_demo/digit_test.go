package digit_demo

import (
	"fmt"
	"testing"
)

func TestDigit(t *testing.T) {
	//fmt.Println(
	//	converToBianry(3),
	//	converToBianry(5),
	//	converToBianry(100),
	//)
}
func TestIntToBytes(t *testing.T) {
	fmt.Println(
		IntToBytes(3),
		IntToBytes(5),
		IntToBytes(253),
		//int([]byte()),
	)

	//    int  translate byte
	lenHandshake := 1111256
	var RecordLen = [4]byte{
		uint8(lenHandshake >> 24),
		uint8(lenHandshake >> 16),
		uint8(lenHandshake >> 8),
		uint8(lenHandshake),
	}
	//   byte  translate int
	fmt.Println(RecordLen)
	var data int
	data = int(RecordLen[0]) << 24
	data += int(RecordLen[1]) << 16
	data += int(RecordLen[2]) << 8
	data += int(RecordLen[3])
	fmt.Println("data: ", data)

	//fmt.Println(BytesToInt(RecordLen))

}
