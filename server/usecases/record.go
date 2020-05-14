package usecases

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//HandleRecordInput is input for HandleRecord
type HandleRecordInput struct {
	Name   string `json:"exercise"`
	Weight string `json:"weight"`
	Reps   string `json:"reps"`
	Sets   string `json:"sets"`
}

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
	Weight string `json:"weight"`
	Reps   string `json:"reps"`
	Sets   string `json:"sets"`
}

//BenchPress is structure for exercise
type BenchPress struct {
	Weight string `json:"weight"`
	Reps   string `json:"reps"`
	Sets   string `json:"sets"`
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
func StoreAndGetData(ctx context.Context, db *sql.DB, records []HandleRecordInput) (IDs, error) {
	//goroutine???
	var crunchID *int
	var deadLiftID *int
	var chinUpID *int
	var dumbellCurlID *int
	var benchPressID *int
	var hipThrustID *int
	var squatID *int
	var shoulderPressID *int
	var cablePressDownID *int
	var err error
	for i := 0; i < len(records); i++ {
		if records[i].Name == "crunch" {
			crunchID, err = tools.BodyWeightRecord(
				ctx,
				db,
				"crunch",
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "dead lift" {
			deadLiftID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"dead_lift",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "chin up" {
			chinUpID, err = tools.BodyWeightRecord(
				ctx,
				db,
				"chin_up",
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "dumbell curl" {
			dumbellCurlID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"dumbell_curl",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "bench press" {
			benchPressID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"bench_press",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "hip thrust" {
			hipThrustID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"hip_thrust",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "squat" {
			squatID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"squat",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else if records[i].Name == "shoulder press" {
			shoulderPressID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"shoulder_press",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		} else {
			cablePressDownID, err = tools.FreeWeightRecord(
				ctx,
				db,
				"cable_press_down",
				records[i].Weight,
				records[i].Reps,
				records[i].Sets,
			)
			if err != nil {
				return IDs{}, errors.WithStack(err)
			}
		}
	}

	return IDs{
		crunchID:         crunchID,
		deadLiftID:       deadLiftID,
		chinUpID:         chinUpID,
		dumbellCurlID:    dumbellCurlID,
		benchPressID:     benchPressID,
		hipThrustID:      hipThrustID,
		squatID:          squatID,
		shoulderPressID:  shoulderPressID,
		cablePressDownID: cablePressDownID,
	}, nil
}
