package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestStrToGBK1(t *testing.T) {
	str := "hello 华夏10进制转36进制："
	fmt.Println(Autngo.EncodeHelper.StrToGBK(str))
	fmt.Println(str)
	fmt.Println(Autngo.EncodeHelper.StrHex(str))
	hexSTT := "68656c6c6f20e58d8ee5a48f3130e8bf9be588b6e8bdac3336e8bf9be588b6efbc9a"
	str = Autngo.EncodeHelper.HexStr(hexSTT)
	fmt.Println(str)
	strs := "幽篁独坐,长啸鸣琴。 禅寂入定,毒龙遁形。 我心无窍,天道酬勤。 我义凛然,鬼魅皆惊。 我情豪溢,天地归心。 我志扬迈,水起风生! 天高地阔,流水行云。 清新治本,直道谋身。"
	ss, err := Autngo.EncodeHelper.ZipEncode(strs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(strs), "===>", len(ss))
	strs2, err2 := Autngo.EncodeHelper.ZipDecode(ss)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(len(ss), "===>", strs2)
}
