package direction

import "testing"

func TestGetDirection(t *testing.T) {
	type args struct {
		sX, sY, dX, dY int
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "Top",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: -1,
			},
			want: North,
		},
		{
			name: "TopRight",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: -1,
			},
			want: NorthEast,
		},
		{
			name: "Right",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: 0,
			},
			want: East,
		},
		{
			name: "BottomRight",
			args: args{
				sX: 0,
				sY: 0,
				dX: 1,
				dY: 1,
			},
			want: SouthEast,
		},
		{
			name: "Bottom",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: 1,
			},
			want: South,
		},
		{
			name: "BottomLeft",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: 1,
			},
			want: SouthWest,
		},
		{
			name: "Left",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: 0,
			},
			want: West,
		},
		{
			name: "TopLeft",
			args: args{
				sX: 0,
				sY: 0,
				dX: -1,
				dY: -1,
			},
			want: NorthWest,
		},
		{
			name: "NoDirection",
			args: args{
				sX: 0,
				sY: 0,
				dX: 0,
				dY: 0,
			},
			want: None,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.sX, tt.args.sY, tt.args.dX, tt.args.dY); got != tt.want {
				t.Errorf("GetDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}
