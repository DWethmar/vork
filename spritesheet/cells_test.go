package spritesheet_test

import (
	"image"
	"reflect"
	"testing"

	"github.com/dwethmar/vork/spritesheet"
)

func TestCreateRectangleGrid(t *testing.T) {
	type args struct {
		columns int
		rows    int
		width   int
		height  int
	}
	tests := []struct {
		name string
		args args
		want [][]image.Rectangle
	}{
		{
			name: "create cells",
			args: args{
				columns: 6,
				rows:    1,
				width:   64,
				height:  64,
			},
			want: [][]image.Rectangle{
				{
					image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 64, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 64, Y: 0}, Max: image.Point{X: 128, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 128, Y: 0}, Max: image.Point{X: 192, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 192, Y: 0}, Max: image.Point{X: 256, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 256, Y: 0}, Max: image.Point{X: 320, Y: 64}},
				},
				{
					image.Rectangle{Min: image.Point{X: 320, Y: 0}, Max: image.Point{X: 384, Y: 64}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spritesheet.CreateRectangleGrid(tt.args.columns, tt.args.rows, tt.args.width, tt.args.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
