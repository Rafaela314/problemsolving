package hacker

import "fmt"

/*An arcade game player wants to climb to the top of the leaderboard and track their ranking. The game uses Dense Ranking, so its leaderboard works like this:

    The player with the highest score is ranked number

    on the leaderboard.
    Players who have equal scores receive the same ranking number, and the next player(s) receive the immediately following ranking number.

Example


The ranked players will have ranks , , , and , respectively. If the player's scores are , and , their rankings after each game are , and . Return

.

Function Description

Complete the climbingLeaderboard function in the editor below.

climbingLeaderboard has the following parameter(s):

    int ranked[n]: the leaderboard scores
    int player[m]: the player's scores

Returns

    int[m]: the player's rank after each new score
*/

func ClimbingLeaderBoard(ranked []int32, player []int32) []int32 {

	//m := make(map[int]int)
	var list []int32

	var found bool

	for _, y := range player {
		found = false

		last := ranked[0]
		position := 1

		for k, v := range ranked {
			if v != last {
				position = position + 1
				last = v
			}

			fmt.Printf("\n last: %v \n", last)
			fmt.Printf("\n position: %v \n", position)

			if y >= v {
				ranked = append(ranked, 0)
				copy(ranked[(k+1):], ranked[(k):])
				ranked[k] = y
				found = true

				if position == 1 {
					list = append(list, int32(position))

				} else {
					list = append(list, int32(position))

				}

				break
			}
		}

		if !found {
			ranked = append(ranked, y)
			list = append(list, int32(position+1))

		}

		fmt.Println(ranked)

	}

	return list

}
