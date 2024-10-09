package direction_test

import (
	"testing"

	"github.com/dwethmar/vork/direction"
)

func TestGetDirection(t *testing.T) {
	type args struct {
		sX, sY, dX, dY int
	}
	tests := []struct {
		name string
		args args
		want direction.Direction
	}{
		{
			name: "Top",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: -1,
			},
			want: direction.North,
		},
		{
			name: "TopRight",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: -1,
			},
			want: direction.NorthEast,
		},
		{
			name: "Right",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: 0,
			},
			want: direction.East,
		},
		{
			name: "BottomRight",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: 1,
			},
			want: direction.SouthEast,
		},
		{
			name: "Bottom",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: 1,
			},
			want: direction.South,
		},
		{
			name: "BottomLeft",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: 1,
			},
			want: direction.SouthWest,
		},
		{
			name: "Left",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: 0,
			},
			want: direction.West,
		},
		{
			name: "TopLeft",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: -1,
			},
			want: direction.NorthWest,
		},
		{
			name: "NoDirection",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: 0,
			},
			want: direction.None,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := direction.Get(tt.args.sX, tt.args.sY, tt.args.dX, tt.args.dY); got != tt.want {
				t.Errorf("GetDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
