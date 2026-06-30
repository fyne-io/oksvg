package oksvg

import "testing"

func TestSetTarget(t *testing.T) {
	type Box struct {
		X, Y, W, H float64
	}

	tests := []struct {
		ViewBox Box
		Target  Box
	}{
		// direct translation
		{ViewBox: Box{0, 0, 24, 24}, Target: Box{0, 0, 24, 24}},
		{ViewBox: Box{0, -24, 24, 24}, Target: Box{0, 0, 24, 24}},
		{ViewBox: Box{-24, 0, 24, 24}, Target: Box{0, 0, 24, 24}},
		{ViewBox: Box{-24, -24, 24, 24}, Target: Box{0, 0, 24, 24}},
		{ViewBox: Box{24, 24, 24, 24}, Target: Box{0, 0, 24, 24}},

		// alignment and proportional scaling
		{ViewBox: Box{0, 0, 24, 24}, Target: Box{0, 0, 36, 36}},
		{ViewBox: Box{0, -24, 24, 24}, Target: Box{0, 0, 36, 36}},
		{ViewBox: Box{-24, 0, 24, 24}, Target: Box{0, 0, 36, 36}},
		{ViewBox: Box{-24, -24, 24, 24}, Target: Box{0, 0, 36, 36}},

		// arbitrary transformations
		{ViewBox: Box{8, 0, 24, 26}, Target: Box{-16, 16, 36, 26}},
		{ViewBox: Box{0, 8, 24, 26}, Target: Box{-8, -8, 64, 8}},
		{ViewBox: Box{-16, -16, 24, 26}, Target: Box{-4, 8, 12, 12}},
		{ViewBox: Box{-4.5, 2.0, 24.5, 26.125}, Target: Box{-4, 8, 16.125, 12.5}},
	}

	for i, test := range tests {
		var icon SvgIcon
		icon.ViewBox = test.ViewBox
		icon.SetTarget(test.Target.X, test.Target.Y, test.Target.W, test.Target.H)

		var tx, ty [4]float64
		tx[0], ty[0] = test.Target.X, test.Target.Y
		tx[1], ty[1] = test.Target.X+test.Target.W, test.Target.Y
		tx[2], ty[2] = test.Target.X+test.Target.W, test.Target.Y+test.Target.H
		tx[3], ty[3] = test.Target.X, test.Target.Y+test.Target.H

		var px, py [4]float64
		px[0], py[0] = icon.Transform.Transform(test.ViewBox.X, test.ViewBox.Y)
		px[1], py[1] = icon.Transform.Transform(test.ViewBox.X+test.ViewBox.W, test.ViewBox.Y)
		px[2], py[2] = icon.Transform.Transform(test.ViewBox.X+test.ViewBox.W, test.ViewBox.Y+test.ViewBox.H)
		px[3], py[3] = icon.Transform.Transform(test.ViewBox.X, test.ViewBox.Y+test.ViewBox.H)

		for n := 0; n < 4; n++ {
			if px[n] != tx[n] || py[n] != ty[n] {
				t.Fatalf(
					"test#%d: point %d: expected (%f, %f), got (%f, %f) %v",
					i, n, tx[n], ty[n], px[n], py[n], test,
				)
			}
		}
	}
}
