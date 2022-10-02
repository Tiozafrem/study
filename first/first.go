package first

type MonthValue struct {
	Month int
	Value float32
}

func SubstractMK(m float32, k float32, p float32, q float32) <-chan MonthValue {
	c := make(chan MonthValue)
	go func() {
		defer close(c)

		for i := 1; i <= 12; i++ {
			c <- MonthValue{
				i,
				m - k,
			}
			k = k - ((k * q) / 100)
			m = m + ((m * p) / 100)
		}

	}()

	return c
}

func SubstractAdvancedMK(m float32, k float32, p float32, q float32) <-chan MonthValue {
	c := make(chan MonthValue)

	go func() {
		defer close(c)
		for i := 1; i <= 12; i++ {
			c <- MonthValue{
				i,
				m - k,
			}
			k = k - ((k * q) / 100)
			q = q / 2
			m = m + ((m * p) / 100)
		}
	}()

	return c
}

func AverageValueYear(k float32, p float32) float32 {
	sum := k
	var n float32
	t := 1

	for i := 2; i <= 12; i++ {
		n = float32((float32((i - 2)) / 10.0))
		k = k + ((k * (p + n)) / 100)
		sum += k
		t++
	}

	return sum / float32(t)
}
