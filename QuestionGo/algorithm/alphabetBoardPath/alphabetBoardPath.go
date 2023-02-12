package alphabetBoardPath

//1138. 字母板上的路径
//
//我们从一块字母板上的位置 (0, 0) 出发，该坐标对应的字符为 board[0][0]。
//
//在本题里，字母板为board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"]，如下所示。

func alphabetBoardPath(target string) string {
	cx, cy := 0, 0
	res := []byte{}
	for _, c := range target {
		nx := int(c-'a') / 5
		ny := int(c-'a') % 5
		if nx < cx {
			for j := 0; j < cx-nx; j++ {
				res = append(res, 'U')
			}
		}
		if ny < cy {
			for j := 0; j < cy-ny; j++ {
				res = append(res, 'L')
			}
		}
		if nx > cx {
			for j := 0; j < nx-cx; j++ {
				res = append(res, 'D')
			}
		}
		if ny > cy {
			for j := 0; j < ny-cy; j++ {
				res = append(res, 'R')
			}
		}
		res = append(res, '!')
		cx = nx
		cy = ny
	}
	return string(res)
}
