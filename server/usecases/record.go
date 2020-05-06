package usecases

import "context"

//Crunch  is structure for exercise
type Crunch struct {
	reps int8
	sets int8
}

//DeadLift  is structure for exercise
type DeadLift struct {
	weight int16
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
	weight int16
	reps   int8
	sets   int8
}

//BenchPress is structure for exercise
type BenchPress struct {
	weight int16
	reps   int8
	sets   int8
}

//HipThrust is structure for exercise
type HipThrust struct {
	weight int16
	reps   int8
	sets   int8
}

//Squat is structure for exercise
type Squat struct {
	weight int16
	reps   int8
	sets   int8
}

//ShoulderPress is structure for exercise
type ShoulderPress struct {
	weight int16
	reps   int8
	sets   int8
}

//CablePressDown is structure for exercise
type CablePressDown struct {
	weight int16
	reps   int8
	sets   int8
}

//Exercise is sturcture for input of storeAndGetData
type Exercise struct {
	crunch         Crunch
	deadLift       DeadLift
	chinUp         ChinUp
	dumbellCurl    Dumbellcurl
	benchPress     BenchPress
	hipThrust      HipThrust
	squat          Squat
	shoulderPress  ShoulderPress
	cablePressDown CablePressDown
}

func StoreAndGetData(ctx context.Context, input Exercise) {

}
