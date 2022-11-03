package easy

import (
	"fmt"
	"github.com/xiye520/leetcode/util"
)

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 31; i >= 0; i-- {
		res += num & 1 << uint32(i)
		num = num >> 1
	}

	return uint32(res)
}

func reverseBits2(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res += num & 1 << uint32(31-i)
		num = num >> 1
	}

	return res
}

func test_rever() {
	fmt.Println(util.DecBin(4294967293))
	fmt.Println(3221225471 == reverseBits2(4294967293)) //3221225471
	fmt.Println(964176192 == reverseBits2(43261596))    //964176192
}
