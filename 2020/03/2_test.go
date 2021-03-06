package _2020

import "testing"

func TestMultiplySlopes(t *testing.T) {
	type args struct {
		rawInput string
		slopes   []slope
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				rawInput: `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
				slopes: []slope{
					{incRow: 1, incCol: 1},
					{incRow: 1, incCol: 3},
					{incRow: 1, incCol: 5},
					{incRow: 1, incCol: 7},
					{incRow: 2, incCol: 1},
				},
			},
			want: 336,
		}, {
			name: "given input",
			args: args{
				rawInput: `.........#.#.#.........#.#.....
...#......#...#.....#.....#....
#.....#.....#.....#.#........#.
......#..#......#.......#......
.#..........#.............#...#
............#..##.......##...##
....#.....#..#....#............
.#..#.........#....#.#....#....
#.#...#...##..##.#..##..#....#.
.#.......#.#...#..........#....
...#...#........##...#..#.....#
..................#..........#.
.....#.##..............#.......
........#....##..##....#.......
...#.....#.##..........#...##..
.......#.#....#............#...
..............#......#......#..
#.......#...........#........##
.......#.......##......#.......
................#....##...#.#.#
#.......#....................#.
.##.#..##..#..#.#.....#.....#..
#...#............#......##....#
.#....##.#......#.#......#..#..
..........#........#.#.#.......
...#...#..........#..#....#....
..#.#...#...#...##...##......#.
......#...#........#.......###.
....#...............#.###...#.#
..................#.....#..#.#.
.#...#..#..........#........#..
#..........##................##
...#.....#...#......#.#......#.
......#..........#.#......#..#.
..#......#.....................
............#.........##.......
......#.......#........#.......
#.#...#...........#.......#....
.#.#........#.#.#....#........#
#.....##........#.#.....#.#....
.#...#..........##...#.....#..#
.........#....###............#.
..#...#..............#.#.#.....
.....#.#.#..#.#.#.###......#.#.
.#.##...#.......###..#.........
.....##....#.##..#.##..#.......
..#...........##......#..#.....
................##...#.........
##.....................#..#.###
...#..#....#...........#.......
.#.............##.#.....#.#....
.......#.#.#....##..#....#...#.
...##..#..........#....#.......
....##......#.........#........
.##....#...........#.#.......#.
...#...#.##.......#.#..........
..#.........#.##...........#...
....#.##........#.......#...##.
...................#..#..#...##
#...#......###..##.#..###......
#.............#.#........#...#.
...#...#..#..#..............#..
#.....#..#...#...#......#.#..#.
...##.............#........##.#
......#.#.........#.#....#...#.
........##............#...#....
............#.#...#......#.....
...#...........#...#...........
.........#.#......#............
....#.............#..#.....#..#
#.....#...........#.....#.#.#.#
.............#.....##......#...
...................###.#......#
#.##.....#...#..#..#...#....#..
............#.....#....#.#.....
#....#..#..........#....#..#...
..........##..................#
....#.......###..#......###....
.......#...#.##.##..#....##....
...##...............#.....#...#
#...........#...#......#..#..#.
..##....#.......#...#.....#....
.......##..............#.##..#.
.#......#..........#.......#...
....##...................#.#..#
......#....###................#
.#........#...........#........
.#.....##....#..##...........#.
##..............##...#.......#.
......#..........#..........##.
..#.....#.#.........####....#..
.............#......#......#...
..#.............#...........##.
#.#...#........#..........##...
...#....#.....#.....#.....##...
......#........................
......#..#...#......#.....#....
......#....##.....#....#.......
#.#......#.##..#...............
..........#.#.##..##..#........
......#.#..#....###.#..........
........................#....#.
#.#.............#....#.....##.#
.#.#..........#.......#..#....#
..#...#......#..#..#...#.#...#.
...#.##...###..................
........#.#...........#...#....
........#..........#....#......
#....#........#.......##.....#.
......###...##...#......#......
............#.......#.....##..#
....#..#.......##......#.......
#............##................
...............#..#......#...#.
.#....##...#......#............
#...#................#.........
..#....#..#........##......#...
....#....###......##.#.......#.
......#.#..#.#..#....#..#......
....#..........#..#..#.........
.#..#.......#......#........#..
.......#..#..#............#....
.............#...#....#..#.....
..............#.#.#.........#..
#.....##.......#....#..........
...#.....#......#..............
...##.#..#.#........#..##....#.
.......#.#.....##..#...........
.....#.#....#..................
.#......#.###..............##..
..#....#...#..#...##...##....#.
..........##..#...#..#.........
..#............#........#.#...#
.........................#.#.#.
......#........#.#..#......##.#
#.#......#..#.........#........
.....#........#......#...#.#...
........##....##....#.#........
....#....#.#...#...##....#..#..
#.............#.....#..........
#............##..#............#
..#.#......#........#..........
.#......#......#.#.##.##.......
..#.....#..........#......##...
...#......#...#.##....#.....##.
......#......#...........#.#...
....#........#..#..#........#.#
....#.........#.....#...#.#.#..
....#.....###........#.........
.............##........#.#.....
...#............#........#.#.#.
......#....#.......#.#.........
.....#................#........
.#....#..#.#.............#...#.
#..##...#............#......#..
...#..#........................
.#.#...........#.......#.......
#....###............##.........
...##....#.#............##.....
.........####......#........#..
.....#.......#.#...............
.......#...#.###..#....#...#..#
...#.....##..#....#..#.#...###.
.............#........#.#.#..#.
................#..........##..
.......####..##..##........##.#
..#......#..#..#.#.##..........
#....#........#....#...###.#.#.
........##........##.....#.....
...........#.#...........#....#
#.............#...........#....
...#.........#...#....#.....#..
..##......#...#...............#
.............#.........#....#..
..#...........#..#........#.##.
.#.#......#.............##...#.
.#......#.....##.#..#..#.......
....##......#..................
.#.#..##............#....#....#
........#...##.............#..#
........#....##.....#......###.
.........#....#.#..............
#.....#........................
.#..#....#.....#......#.###..#.
..........#...#....##....#..#..
...#.#.#...##..#..#......#..#.#
#............#.....#....#......
#.###...#.#......###..#....#..#
...#.###........#......#....#..
..#.##...#.........#.#.........
............##.................
....#..........#.#.#.#.#....#..
...##.#...#.......#.......##..#
....##.#........#....#...#.....
.............#.#....#...#.....#
...#......................##...
..#...#.....#.....#........#..#
..#..#.......#....#..##.....#..
..#..#.#.......................
.......##..#....#....#..#......
....#......##....#............#
.#...#..#..#.##...#.#...#......
.....#......#....#.........#...
.##......##.........#....#.....
#...........#...##.....#......#
.....#.#.......#.........#.....
.........#..........#..####.##.
............#..#......#.#......
.#.............#........#.#....
......#......#...#..#....#####.
.........##.#..##...###..#....#
....#.#....#.#..#.........#....
..#.............#...##...##....
........#..........#.##..#....#
.....#...#..##........#.#..#...
##..#.#.....#............#.....
.............#........##...##..
#......####.....##.............
..##.....##....###..#.#....#...
......##.##.#...#..#.#..##.....
......#.................#......
#.....#.#...#......#.#....#....
....#.#........#..............#
##........#.......##.#...##...#
..#..................#.#....#..
...........#..........#.#.....#
........##.#.....#......#..#..#
.....#....#..#.....#.........##
#.#..#..#...#......#..........#
#...##.....#..#.#.......#.##...
..#....##...............#......
#..........#.#.........#.#....#
..............#......#....#....
.....#...........#...#...#...#.
...#......#....#....#..........
.#..........#.#....##..##....#.
..............#.........#.#....
.......#.....#.....#...##....#.
##.#.........#....#.....#.#....
....#..#......#................
......##.....#.......##........
.....##...#........#...#...#...
..#...#...#..#..#.#......#..#..
....#...#.......#..............
....#..#.........###........#..
....#.............##..#........
..........##.#.......##..##....
#.##..................#.....#..
#........#........#.....#......
.#...#......#..................
#....##.##......#...#.........#
......#.##..##................#
............#.........##.......
..........####.#........#.....#
.##...#...#....#..#............
.#.##...#..#...#......#......##
.....#.#....#..###......#.#.#..
...#.......................##..
......................#.......#
..#....#.........#..#.#.....#..
.#....#..#....#...#............
..........#...##.....#.#..#....
........#..#..#....#...#...#...
.....#......#.#................
.....#...........#...#.........
.....#...##..#.#....#..#.....#.
#.......#.............##.......
................#....#.#..#....
.#..##...#.#........#......#.#.
.#.##..........#...............
....##......#....#........#....
....#..#....#.##.#.............
.......#..#......##.#.....#....
.......#.....#.............#...
.....#....#.......#............
........#.#...##..##..##.......
#.........##....##...##........
........#..#.#..........###.#..
..........................#.#..
#.....#.......#..#........#....
...##.....#.......#......#.....
.#.#..#...........#...........#
.....##..#........#...####.....
.#.#...##.#.#..#..#.#..#.......
..#.##.#...#.#.#...#..#........
............#..........#..#....
...............#..##.#.........
.............#.....#....#......
...##..##......##..........#...
..#.......#....#..........#...#
.##................#.#.#.......
.....##.....#..#.....#.........
......#.#.......#......#..#....
.....#.....#........#.......##.
......#.......##......#...#...#
....#...........###.........#..
...#.....#.........##........#.
..#.....#..............#.......
....#.......#...#....#....#..##
......#...........#...........#
.##......#......#.#.....#.##...
....#..##......#...#..#.#.###..
.......#.#....#......#..#......
..........#........#...........
#.##.........#.#.#...#...#.#...
.#......###.....#....#.#....#..
...................##..#.......
....#..#..............#.#.....#
#..................#.....#.....
...........##.##.......#..#.#..
........#.#......#...........#.
#..#.......#...#...........#.#.
......##...........#...........
.........#.#........#........#.
#......#....#.#.....#..#.......
............#..#.....##...#....
.#......#..#......#.........#..
.......#...#.........#.##.....#
........................#..#...
.###..............#.#..#.......
.....#.........#.......#......#
..##..##....#.....#.......#.#..
...###.#..#.##............#....`,
				slopes: []slope{
					{incRow: 1, incCol: 1},
					{incRow: 1, incCol: 3},
					{incRow: 1, incCol: 5},
					{incRow: 1, incCol: 7},
					{incRow: 2, incCol: 1},
				},
			},
			want: 958815792,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplyTreesInSlopes(tt.args.rawInput, tt.args.slopes); got != tt.want {
				t.Errorf("MultiplyTreesInSlopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
