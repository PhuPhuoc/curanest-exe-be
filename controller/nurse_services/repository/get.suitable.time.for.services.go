package nurserepository

import (
	"fmt"
	"time"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
)

type TimeRange struct {
	Date  string
	Start time.Time
	End   time.Time
	ID    string
}

func (store *nurseStore) GetSuitableTimeForServices(nurse_id, from, to string, totalTimeMinute int) ([]nursemodel.TimeFrameServices, error) {
	shifts := []nursemodel.ShiftCurrent{}
	query := `
		select id, shift_date, shift_from, shift_to, status, appointment_id
		from work_schedules
		where nurse_id=? and shift_date between ? and ? and status = ? and appointment_id is null
		order by shift_date, shift_from 
	`
	if err_query := store.db.Select(&shifts, query, nurse_id, from, to, "available"); err_query != nil {
		return nil, fmt.Errorf("cannot select schedule from db <%w>", err_query)
	}

	// todo: convert shift right format -> time
	timeRanges := []TimeRange{}
	for _, shift := range shifts {
		start, err1 := time.Parse("15:04:05", shift.ShiftFrom)
		end, err2 := time.Parse("15:04:05", shift.ShiftTo)
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid shift time format")
		}
		timeRanges = append(timeRanges, TimeRange{
			Date:  shift.ShiftDate,
			Start: start,
			End:   end,
			ID:    shift.ID,
		})
	}

	// todo: find suitable timeframe
	// todo - change totalTimeMinute (total time (minute) of all services customer want to book) to duration
	requireDuration := time.Duration(totalTimeMinute) * time.Minute
	suitabletimes := []nursemodel.TimeFrameServices{}
	for i := 0; i < len(timeRanges); i++ {
		currentDate := timeRanges[i].Date
		currentStart := timeRanges[i].Start
		currentEnd := timeRanges[i].End
		totalDuration := currentEnd.Sub(currentStart)

		nuresScheduleIDs := []string{timeRanges[i].ID}

		if totalDuration >= requireDuration {
			suitabletimes = append(suitabletimes, nursemodel.TimeFrameServices{
				AppointmentDate:  currentDate,
				From:             currentStart.Format("15:04:05"),
				To:               currentEnd.Format("15:04:05"),
				NurseScheduleIDs: nuresScheduleIDs,
			})
			continue
		}

		for j := i + 1; j < len(timeRanges); j++ {
			if timeRanges[j].Date == currentDate && timeRanges[j].Start == currentEnd {
				totalDuration += timeRanges[j].End.Sub(timeRanges[j].Start)
				currentEnd = timeRanges[j].End
				nuresScheduleIDs = append(nuresScheduleIDs, timeRanges[j].ID)

				// kiểm tra tổng thời gian đã đạt yêu cầu chưa
				if totalDuration >= requireDuration {
					suitabletimes = append(suitabletimes, nursemodel.TimeFrameServices{
						AppointmentDate:  currentDate,
						From:             currentStart.Format("15:04:05"),
						To:               currentEnd.Format("15:04:05"),
						NurseScheduleIDs: nuresScheduleIDs,
					})
					break
				}
			} else {
				break
			}
		}
	}
	return suitabletimes, nil
}
