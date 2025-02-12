package datamatrix

type dmCodeSize struct {
	Rows                  int
	Columns               int
	RegionCountHorizontal int
	RegionCountVertical   int
	ECCCount              int
	BlockCount            int
}

func (s *dmCodeSize) RegionRows() int {
	return (s.Rows - (s.RegionCountVertical * 2)) / s.RegionCountVertical
}

func (s *dmCodeSize) RegionColumns() int {
	return (s.Columns - (s.RegionCountHorizontal * 2)) / s.RegionCountHorizontal
}

func (s *dmCodeSize) MatrixRows() int {
	return s.RegionRows() * s.RegionCountVertical
}

func (s *dmCodeSize) MatrixColumns() int {
	return s.RegionColumns() * s.RegionCountHorizontal
}

func (s *dmCodeSize) DataCodewords() int {
	return ((s.MatrixColumns() * s.MatrixRows()) / 8) - s.ECCCount
}

func (s *dmCodeSize) DataCodewordsForBlock(idx int) int {
	if s.Rows == 144 && s.Columns == 144 {
		// Special Case...
		if idx < 8 {
			return 156
		} else {
			return 155
		}
	}
	return s.DataCodewords() / s.BlockCount
}

func (s *dmCodeSize) ErrorCorrectionCodewordsPerBlock() int {
	return s.ECCCount / s.BlockCount
}

var codeSizes []*dmCodeSize = []*dmCodeSize{
	{10, 10, 1, 1, 5, 1},
	{12, 12, 1, 1, 7, 1},
	{14, 14, 1, 1, 10, 1},
	{16, 16, 1, 1, 12, 1},
	{18, 18, 1, 1, 14, 1},
	{20, 20, 1, 1, 18, 1},
	{22, 22, 1, 1, 20, 1},
	{24, 24, 1, 1, 24, 1},
	{26, 26, 1, 1, 28, 1},
	{32, 32, 2, 2, 36, 1},
	{36, 36, 2, 2, 42, 1},
	{40, 40, 2, 2, 48, 1},
	{44, 44, 2, 2, 56, 1},
	{48, 48, 2, 2, 68, 1},
	{52, 52, 2, 2, 84, 2},
	{64, 64, 4, 4, 112, 2},
	{72, 72, 4, 4, 144, 4},
	{80, 80, 4, 4, 192, 4},
	{88, 88, 4, 4, 224, 4},
	{96, 96, 4, 4, 272, 4},
	{104, 104, 4, 4, 336, 6},
	{120, 120, 6, 6, 408, 6},
	{132, 132, 6, 6, 496, 8},
	{144, 144, 6, 6, 620, 10},
}
