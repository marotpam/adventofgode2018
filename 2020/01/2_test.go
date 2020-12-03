package main

import "testing"

func TestFindMultiplicationFor3(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given example",
			args: args{
				numbers: []int{1721, 979, 366, 299, 675, 1456},
			},
			want: 241861950,
		},
		{
			name: "given input",
			args: args{
				numbers: []int{1293, 1207, 1623, 1675, 1842, 1410, 85, 1108, 557, 1217, 1506, 1956, 1579, 1614, 1360, 1544, 1946, 1666, 1972, 1814, 1699, 1778, 1529, 2002, 1768, 1173, 1407, 1201, 1264, 1739, 1774, 1951, 1980, 1428, 1381, 1714, 884, 1939, 1295, 1694, 1168, 1971, 1352, 1462, 1828, 1402, 1433, 1542, 1144, 1331, 1427, 1261, 1663, 1820, 1570, 1874, 1486, 1613, 1769, 1721, 1753, 1142, 1677, 2010, 1640, 1465, 1171, 534, 1790, 2005, 1604, 1891, 1247, 1281, 1867, 1403, 2004, 1668, 1416, 2001, 1359, 686, 1965, 1728, 1551, 1565, 1128, 1832, 1757, 1350, 1808, 1711, 1799, 1590, 1989, 1547, 1140, 1905, 1368, 1179, 1902, 1473, 1908, 1859, 1257, 1394, 1244, 1800, 1695, 1731, 1474, 1781, 1885, 1154, 1990, 1929, 1193, 1302, 1831, 1226, 1418, 1400, 1435, 1645, 1655, 1843, 1227, 1481, 1754, 1290, 1685, 1498, 71, 1286, 1137, 1288, 1758, 1987, 1471, 1839, 1545, 1682, 1615, 1475, 1849, 1985, 1568, 1795, 1184, 1863, 1362, 1271, 1802, 1944, 1821, 1880, 1788, 1733, 1150, 1314, 1727, 1434, 1833, 1312, 1457, 160, 1629, 1967, 1505, 1239, 1266, 1838, 1687, 1630, 1591, 1893, 1450, 1234, 1755, 1523, 1533, 1499, 1865, 1725, 1444, 1517, 1167, 1738, 1519, 1263, 1901, 1627, 1644, 1771, 1812, 1270, 1497, 1707, 1708, 1396},
			},
			want: 61515678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMultiplicationFor3(tt.args.numbers); got != tt.want {
				t.Errorf("FindMultiplicationFor3() = %v, want %v", got, tt.want)
			}
		})
	}
}