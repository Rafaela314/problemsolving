package codility

/*
A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.

Count the minimal number of jumps that the small frog must perform to reach its target.

Write a function:

class Solution { public int solution(int X, int Y, int D); }

that, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.

For example, given:

  X = 10
  Y = 85
  D = 30
the function should return 3
*/

func FrogJmp(X, Y, D int) int {

	var result int

	distance := Y - X
	if distance%D != 0 {
		result = ((distance / D) + 1)
	} else {
		result = distance / D
	}
	return result
}

/*
func main() {
	x := FrogJmp(0, 95, 9)

	fmt.Println(x)

	//fmt.Println("hello")
}
*/
