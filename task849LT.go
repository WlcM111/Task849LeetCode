package maxDistToClosest

func maxDistToClosest(seats []int) int {
	start := 0
	res := 0
	cnt := 0

	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			if cnt == 0 {
				start = i
			}

			cnt++

			if i == len(seats)-1 {
				return max(cnt, res)
			}

		} else {

			if start == 0 {

				res = max(cnt, res)
			} else {

				res = max((cnt+1)/2, res)
			}

			cnt = 0
		}
	}

	return res
}
