package shapes

// Perimeter 長方形の周囲の長さ
func Perimeter(w float64, h float64) float64 {
	return 2 * (w + h)
}

// Area 長方形の面積
func Area(w float64, h float64) float64 {
	return w * h
}
