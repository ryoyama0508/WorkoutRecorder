package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//Analysis ...
func Analysis(ctx context.Context, db *sql.DB, userID string) ([][]float64, [][]float64, error) {
	body := make([][]float64, 2)
	free := make([][]float64, 7)
	crunchRec, err := tools.BodyWeightRecordGet(ctx, db, userID, "crunch")
	if err != nil {
		return nil, nil, err
	}
	body[0] = crunchRec

	deadRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "dead_lift")
	if err != nil {
		return nil, nil, err
	}
	free[0] = deadRec

	chinUpRec, err := tools.BodyWeightRecordGet(ctx, db, userID, "chin_up")
	if err != nil {
		return nil, nil, err
	}
	body[1] = chinUpRec

	dbCurlRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "dumbell_curl")
	if err != nil {
		return nil, nil, err
	}
	free[1] = dbCurlRec

	bpRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "bench_press")
	if err != nil {
		return nil, nil, err
	}
	free[2] = bpRec

	hipThrustRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "hip_thrust")
	if err != nil {
		return nil, nil, err
	}
	free[3] = hipThrustRec

	sqRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "squat")
	if err != nil {
		return nil, nil, err
	}
	free[4] = sqRec

	shoulPressRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "shoulder_press")
	if err != nil {
		return nil, nil, err
	}
	free[5] = shoulPressRec

	cablePressDownRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "cable_press_down")
	if err != nil {
		return nil, nil, err
	}
	free[6] = cablePressDownRec

	fmt.Println(free, "free")
	fmt.Println(body, "body")

	return body, free, nil
}
