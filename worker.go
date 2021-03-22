package main

import (
	"strconv"
)

func (Store *Store) OrderQueue() {
	for _, Item := range Store.Data {
		if Store.OrderNo < JSONConfig().Size {
			TargetStartTime := RegMatch(Item.StartTime, JSONConfig().Prefer.Start)
			TargetEndTime := RegMatch(Item.EndTime, JSONConfig().Prefer.End)

			if TargetStartTime != "" || TargetEndTime != "" && Item.Status == "可用" {
				Store.StateID = strconv.Itoa(Item.Number)
				orderCoachNewResp := Store.OrderCoachNew()

				if orderCoachNewResp.D.Item1 {
					Store.OrderNo++
				}
			}
		}
	}

	if Store.OrderNo < JSONConfig().Size {
		for {
			for _, Item := range Store.Data {
				Store.StateID = strconv.Itoa(Item.Number)
				orderCoachNewResp := Store.OrderCoachNew()
				if orderCoachNewResp.D.Item1 {
					Store.OrderNo++
				}
			}
		}
	}
}
