package numconv

import "testing"

func TestArabicToChinese(t *testing.T) {
	for i := 0; i < 20; i++ {
		testArabicToChinese(t, i)
	}
	testArabicToChinese(t, 1234)
	testArabicToChinese(t, 10234)
	testArabicToChinese(t, 12034)
	testArabicToChinese(t, 54321)
	testArabicToChinese(t, 123456)
	testArabicToChinese(t, 10000100)
	testArabicToChinese(t, 101000000)
}

func testArabicToChinese(t *testing.T, n int) {
	t.Log(n, "=", ArabicToChinese(n))
}
