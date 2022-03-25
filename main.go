package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Byte2BinStr([]byte{0x2F, 0xFF, 0x43, 0x90}))

	fmt.Printf("%x", BinStr2Byte(Byte2BinStr([]byte{0x2F, 0xFF, 0x43, 0x90})))
}

func BinStr2Byte(ss []string) []byte {
	if len(ss)%2 != 0 {
		return nil
	}
	var ret []byte
	for i := 0; i < len(ss); i += 2 {
		ui, _ := strconv.ParseUint(ss[i], 2, 64)
		tmpstr := fmt.Sprintf("%x", ui)
		ui, _ = strconv.ParseUint(ss[i+1], 2, 64)
		tmpstr += fmt.Sprintf("%x", ui)
		btmp, _ := hex.DecodeString(tmpstr)
		ret = append(ret, btmp[0])
	}
	return ret
}

func Byte2BinStr(b []byte) []string {
	var ret []string
	for _, v := range b {
		str := fmt.Sprintf("%x", v)
		b0, _ := hex.DecodeString("0" + string(str[0]))
		ret = append(ret, patch(strconv.FormatInt(int64(b0[0]), 2)))
		b0, _ = hex.DecodeString("0" + string(str[1]))
		strconv.FormatInt(int64(b0[0]), 2)
		ret = append(ret, patch(strconv.FormatInt(int64(b0[0]), 2)))

	}
	return ret
}

func patch(s string) string {
	switch s {
	case "0":
		return "000" + s
	case "1":
		return "000" + s
	case "10":
		return "00" + s
	case "11":
		return "00" + s
	case "100":
		return "0" + s
	case "101":
		return "0" + s
	case "110":
		return "0" + s
	case "111":
		return "0" + s
	}
	return s
}
