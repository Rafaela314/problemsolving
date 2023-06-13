package others

import "fmt"

func DaysOpen(openHours []map[string]string) []map[string]string {

	list := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	temp := map[string]map[string]string{}

	result := []map[string]string{}

	for _, v := range openHours {
		temp[v["day"]] = v
	}

	initDay := "Monday"
	var closeDay string
	var openHour string
	var closeHour string

	_, ok := temp[list[0]]
	if ok {
		openHour = temp[list[0]]["open"]
		closeHour = temp[list[0]]["close"]

	} else {
		openHour = ""
		closeHour = ""
	}

	count := 1

	for i := 1; i < len(list); i++ {

		_, ok := temp[list[i]]

		if ok {

			//exists && sametime
			if openHour == temp[list[i]]["open"] && closeHour == temp[list[i]]["close"] {

				closeDay = list[i]

				count++
				continue

			} else {

				if count > 1 {
					result = append(result, map[string]string{"days": initDay + " - " + closeDay, "open": openHour, "close": closeHour})
				} else {
					result = append(result, map[string]string{"days": initDay, "open": openHour, "close": closeHour})
				}

				count = 1
				initDay = list[i]
				openHour = temp[list[i]]["open"]
				closeHour = temp[list[i]]["close"]

			}

		}

		if !ok {

			//do not exists but with same previous time
			if openHour == "" && closeHour == "" {
				closeDay = list[i]

				count++

				continue

			} else {
				if count > 1 {
					result = append(result, map[string]string{"days": initDay + " - " + closeDay, "open": openHour, "close": closeHour})
				} else {
					result = append(result, map[string]string{"days": initDay, "open": openHour, "close": closeHour})
				}

				initDay = list[i]
				openHour = ""
				closeHour = ""
				count = 1

			}

		}
	}

	if count > 1 {
		result = append(result, map[string]string{"days": initDay + " - " + closeDay, "open": openHour, "close": closeHour})
	} else {
		result = append(result, map[string]string{"days": initDay, "open": openHour, "close": closeHour})
	}

	fmt.Printf("\nRESULT %v \n", result)

	return result
}
