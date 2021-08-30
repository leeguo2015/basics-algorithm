package digit_demo

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

// 数字转换成二进制
func IntToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := uint8(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint8
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
