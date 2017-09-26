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

func BenchmarkArabicToChinese(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArabicToChinese(i)
	}
}

func TestChineseToArabic(t *testing.T) {
	for _, s := range []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "一十"} {
		testChineseToArabic(t, s)
	}
	testChineseToArabic(t, "四十三")
	testChineseToArabic(t, "二百零三")
	testChineseToArabic(t, "一万零二十三")
	//testChineseToArabic(t, "一万零啊二十三") // !!! panic !!!
	testChineseToArabic(t, "五万四千三百二十一")
	testChineseToArabic(t, "一亿零五百零七")
	testChineseToArabic(t, "一百亿五千")
	testChineseToArabic(t, "一百亿零一百万")
}

func testChineseToArabic(t *testing.T, s string) {
	t.Log(s, "=", ChineseToArabic(s))
}

func BenchmarkChineseToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChineseToArabic(ArabicToChinese(i))
	}
}
