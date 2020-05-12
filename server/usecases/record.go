package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//Crunch  is structure for exercise
type Crunch struct {
	reps int8
	sets int8
}

//DeadLift  is structure for exercise
type DeadLift struct {
	weight float32
	reps   int8
	sets   int8
}

//ChinUp  is structure for exercise
type ChinUp struct {
	reps int8
	sets int8
}

//Dumbellcurl is structure for exercise
type Dumbellcurl struct {
	weight float32
	reps   int8
	sets   int8
}

//BenchPress is structure for exercise
type BenchPress struct {
	weight float32 `json:"weight"`
	reps   int8    `json:"reps"`
	sets   int8    `json:"sets"`
}

//HipThrust is structure for exercise
type HipThrust struct {
	weight float32
	reps   int8
	sets   int8
}

//Squat is structure for exercise
type Squat struct {
	weight float32
	reps   int8
	sets   int8
}

//ShoulderPress is structure for exercise
type ShoulderPress struct {
	weight float32
	reps   int8
	sets   int8
}

//CablePressDown is structure for exercise
type CablePressDown struct {
	weight float32
	reps   int8
	sets   int8
}

//ExerciseRecord is sturcture for input of storeAndGetData
type ExerciseRecord struct {
	/* userID         int */
	crunch         *Crunch
	deadLift       *DeadLift
	chinUp         *ChinUp
	dumbellCurl    *Dumbellcurl
	benchPress     *BenchPress
	hipThrust      *HipThrust
	squat          *Squat
	shoulderPress  *ShoulderPress
	cablePressDown *CablePressDown
}

//IDs are str for outputting record DB ids which are insarted
type IDs struct {
	crunchID         *int
	deadLiftID       *int
	chinUpID         *int
	dumbellCurlID    *int
	benchPressID     *int
	hipThrustID      *int
	squatID          *int
	shoulderPressID  *int
	cablePressDownID *int
}

//StoreAndGetData ...
func StoreAndGetData(ctx context.Context, db *sql.DB, weight, reps, sets string) (IDs, error) {
	var bpID *int
	var err error
	// if input.benchPress != nil {
	// 	bpID, err = tools.ChestRecord(
	// 		ctx,
	// 		db,
	// 		/* input.userID, */
	// 		input.benchPress.weight,
	// 		input.benchPress.reps,
	// 		input.benchPress.sets,
	// 	)
	// 	if err != nil {
	// 		return IDs{}, errors.WithStack(err)
	// 	}
	// }

	//本来は線種目いっぺんによびだして記録する
	//goroutine???

	bpID, err = tools.ChestRecord(
		ctx,
		db,
		weight,
		reps,
		sets,
	)
	if err != nil {
		return IDs{}, errors.WithStack(err)
	}

	fmt.Println(bpID)
	return IDs{
		benchPressID: bpID,
	}, nil
}
