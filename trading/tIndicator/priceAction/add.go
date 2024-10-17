package priceAction

import "github.com/yasseldg/go-simple/trading/tCandle"

func (pa *Base) Add(candle tCandle.Inter) {
	pa.mu.Lock()
	defer pa.mu.Unlock()

	if pa.last == nil {
		pa.last = candle

		pa.new_low = 0
		pa.new_high = 0

		pa.last_low = pa.funcLow(candle)
		pa.last_high = pa.funcHigh(candle)

		pa.Increase()
		return
	}

	if pa.last.Ts() >= candle.Ts() {
		return
	}

	pa.Increase()

	if pa.funcLow(candle) < pa.funcLow(pa.last) {
		if pa.swing {
			if pa.new_low == 0 || pa.funcLow(candle) < pa.new_low {
				pa.new_low = pa.funcLow(candle)
			}
		} else {
			pa.new_low = pa.funcLow(candle)
		}
	}

	if pa.funcHigh(candle) > pa.funcHigh(pa.last) {
		if pa.swing {
			if pa.funcHigh(candle) > pa.new_high {
				pa.new_high = pa.funcHigh(candle)
			}
		} else {
			pa.new_high = pa.funcHigh(candle)
		}
	}

	pa.last = candle

	switch pa.state {
	case neutral_down, neutral_up:
		pa.neutral(candle)

	case uptrend:
		pa.uptrend(candle)

	case downtrend:
		pa.downtrend(candle)
	}
}

// private methods

func (pa *Base) neutral(candle tCandle.Inter) {

	if candle.Close() > pa.last_high {

		pa.last_high = pa.funcHigh(candle)

		if pa.new_low > 0 {

			pa.last_low = pa.new_low
			pa.new_low = 0

			pa.state = uptrend

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
		return
	}

	if pa.high_low && candle.High() > pa.last_high {
		pa.last_high = candle.High()
	}

	if candle.Close() < pa.last_low {

		pa.last_low = pa.funcLow(candle)

		if pa.new_high > 0 {

			pa.last_high = pa.new_high
			pa.new_high = 0

			pa.state = downtrend

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
		return
	}

	if pa.high_low && candle.Low() < pa.last_low {
		pa.last_low = candle.Low()
	}
}

func (pa *Base) uptrend(candle tCandle.Inter) {

	if candle.Close() > pa.last_high {

		pa.last_high = pa.funcHigh(candle)

		if pa.new_low > 0 {

			pa.last_low = pa.new_low
			pa.new_low = 0

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
		return
	}

	if pa.high_low && candle.High() > pa.last_high {
		pa.last_high = candle.High()
	}

	if pa.funcLow(candle) < pa.last_low {

		pa.last_low = pa.funcLow(candle)

		if pa.new_high > 0 {

			pa.last_high = pa.new_high
			pa.new_high = 0

			pa.state = neutral_up

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
	}
}

func (pa *Base) downtrend(candle tCandle.Inter) {

	if candle.Close() < pa.last_low {

		pa.last_low = pa.funcLow(candle)

		if pa.new_high > 0 {

			pa.last_high = pa.new_high
			pa.new_high = 0

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
		return
	}

	if pa.high_low && candle.Low() < pa.last_low {
		pa.last_low = candle.Low()
	}

	if pa.funcHigh(candle) > pa.last_high {

		pa.last_high = pa.funcHigh(candle)

		if pa.new_low > 0 {

			pa.last_low = pa.new_low
			pa.new_low = 0

			pa.state = neutral_down

			pa.values = append(pa.values, newInd(candle.Ts(), pa.state))
		}
	}
}
