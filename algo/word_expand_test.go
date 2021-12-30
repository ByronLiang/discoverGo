package algo

import "testing"

func TestByteExpand(t *testing.T) {
	t.Log(ByteExpand("abc1[a2[k]x2[mm]]zz"))
}
