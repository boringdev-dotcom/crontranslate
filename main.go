package main

import (
	"fmt"
	"strings"
)

func ToCronExpression(req string) string {
	cronFields := []string{"seconds", "minutes", "hours", "dayofmonth", "month", "dayofweek", "year"}

	finalString := []string{"*", "*", "*", "?", "*", "*", "*"}

	SplitExpr := strings.Split(req, ",")
	fmt.Println(SplitExpr)
	for _, v := range cronFields {
		cronTranslator(v, &finalString, &SplitExpr)
	}
	return strings.Join(finalString, " ")
}

func contains(s []string, str string) (bool, int) {
	for k, v := range s {
		if strings.Contains(strings.ToLower(v), strings.ToLower(str)) {
			return true, k
		}
	}

	return false, -1
}

func cronTranslator(expr string, finalString *[]string, splitExpr *[]string) {
	switch expr {
	case "seconds":
		fmt.Println("Seconds")
		res, index := contains(*splitExpr, "second")
		if res {
			secondString := (*splitExpr)[index]
			fmt.Println(secondString)
		} else {
			(*finalString)[0] = "0"
		}
	case "minutes":
		fmt.Println("Minutes")
		res, index := contains(*splitExpr, "minute")
		if res {
			minuteString := (*splitExpr)[index]
			minuteLogic := strings.Split(minuteString, " ")
			//Search At
			At, atIndex := contains(minuteLogic, "At")
			if At {
				minuteMark := minuteLogic[atIndex+1]
				(*finalString)[1] = minuteMark
			}

			//Search Every
			Every, everyIndex := contains(minuteLogic, "Every")

			if Every {
				everyMinValue := minuteLogic[everyIndex+1]
				(*finalString)[1] = fmt.Sprintf("%s/%s", (*finalString)[1], everyMinValue)
			}
		} else {
			(*finalString)[1] = "0"
		}
	case "hours":
		fmt.Println("Hours")
		resA, indexA := contains(*splitExpr, "am")
		resP, indexP := contains(*splitExpr, "pm")
		if resA {
			hourString := (*splitExpr)[indexA]
			hourLogic := strings.Split(hourString, " ")

			//Search At
			At, atIndex := contains(hourLogic, "At")
			if At {
				hourMark := hourLogic[atIndex+1]
				(*finalString)[2] = hourMark
			}
			fmt.Println(hourString, hourLogic)
		} else if resP {
			hourString := (*splitExpr)[indexP]
			fmt.Println(hourString)
		} else {
			(*finalString)[2] = "0"
		}

		//Search Every
		every, index := contains(*splitExpr, "hour")
		if every {
			everyHourString := (*splitExpr)[index]
			everyHourLogic := strings.Split(everyHourString, " ")

			Every, everyIndex := contains(everyHourLogic, "Every")
			if Every {
				everyValue := everyHourLogic[everyIndex+1]
				(*finalString)[2] = fmt.Sprintf("%s/%s", (*finalString)[2], everyValue)
			}
		}

	case "dayofmonth":
		fmt.Println("dayofmonth")
		res, index := contains(*splitExpr, "nd")
		if res {
			dayOfWeekString := (*splitExpr)[index]
			fmt.Println(dayOfWeekString)
		} else {
			(*finalString)[3] = "?"
		}
	case "month":
		fmt.Println("month")
		res, index := contains(*splitExpr, "month")
		if res {
			dayOfWeekString := (*splitExpr)[index]
			fmt.Println(dayOfWeekString)
		} else {
			(*finalString)[4] = "*"
		}
	case "dayofweek":
		fmt.Println("dayofweek")
		res, index := contains(*splitExpr, "day")
		if res {
			dayString := (*splitExpr)[index]
			dayLogic := strings.Split(dayString, " ")

			//Search on
			On, OnIndex := contains(dayLogic, "On")
			if On {
				dayMark := dayLogic[OnIndex+1]
				(*finalString)[5] = strings.ReplaceAll(dayMark, "day", "")
			}

		} else {
			(*finalString)[5] = "*"
		}

	case "year":
		fmt.Println("year")
	}
}
