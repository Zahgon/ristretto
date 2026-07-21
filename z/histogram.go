package z

func HistogramBounds(minExponent, maxExponent uint32) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func Fibonacci(num int) []float64 { _ = "STUB: not implemented"; return nil }

type HistogramData struct {
	Bounds         []float64
	Count          int64
	CountPerBucket []int64
	Min            int64
	Max            int64
	Sum            int64
}

func NewHistogramData(bounds []float64) *HistogramData { _ = "STUB: not implemented"; return nil }

func (histogram *HistogramData) Copy() *HistogramData { _ = "STUB: not implemented"; return nil }

func (histogram *HistogramData) Update(value int64) { _ = "STUB: not implemented"; return }

func (histogram *HistogramData) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (histogram *HistogramData) String() string { _ = "STUB: not implemented"; return "" }

func (histogram *HistogramData) Percentile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (histogram *HistogramData) Clear() { _ = "STUB: not implemented"; return }
