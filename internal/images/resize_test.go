package images

import (
	"fmt"
	"image"
	"math"
	"testing"
)

func TestAspectRatio(t *testing.T) {
	testTolerance := 0.00001
	aspectTests := []struct {
		width  int
		height int
		aspect float64
	}{
		{10, 5, 2.0},
		{1920, 1080, 1.777777778},
		{640, 480, 1.333333333},
		{320, 240, 1.333333333},
	}
	for _, tt := range aspectTests {
		t.Run(
			fmt.Sprintf("%d/%d=%f", tt.width, tt.height, tt.aspect),
			func(t *testing.T) {
				rect := image.Rect(0, 0, tt.width, tt.height)

				res := aspectRatio(rect)
				if math.Abs(res-tt.aspect) > testTolerance {
					t.Errorf("got %f, want %f", res, tt.aspect)
				}
			})
	}
}

func TestSourceRect(t *testing.T) {

	rectTests := []struct {
		width  int
		height int
		aspect float64
		rect   image.Rectangle
	}{
		{10, 5, 2.0, image.Rect(0, 0, 10, 5)},
		{1920, 1080, 1.777777778, image.Rect(0, 0, 1920, 1080)},
		{1920, 1080, 1.333333333, image.Rect(240, 0, 1680, 1080)},
		{640, 480, 1.333333333, image.Rect(0, 0, 640, 480)},
		{640, 480, 1.777777778, image.Rect(0, 60, 640, 420)},
		{1, 1, 1.333333333, image.Rect(0, 0, 1, 1)},
		{1, 1, 1.777777778, image.Rect(0, 0, 1, 1)},
		{1, 1, 0.75, image.Rect(0, 0, 1, 1)},
	}
	for _, tt := range rectTests {
		t.Run(
			fmt.Sprintf("%dx%d:%f", tt.width, tt.height, tt.aspect),
			func(t *testing.T) {
				res := sourceRect(image.Rect(0, 0, tt.width, tt.height), tt.aspect)
				if res != tt.rect {
					t.Errorf("got %q, want %q", res, tt.rect)
				}
			})
	}
}
