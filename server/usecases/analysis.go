package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//Analysis ...
func Analysis(ctx context.Context, db *sql.DB, userID string) ([][]uint, [][]float64, error) {
	body := make([][]uint, 2)
	free := make([][]float64, 7)
	crunchRec, err := tools.BodyWeightRecordGet(ctx, db, userID, "crunch")
	if err != nil {
		return nil, nil, err
	}
	body = append(body, crunchRec)

	deadRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "dead_lift")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, deadRec)

	chinUpRec, err := tools.BodyWeightRecordGet(ctx, db, userID, "chin_up")
	if err != nil {
		return nil, nil, err
	}
	body = append(body, chinUpRec)

	dbCurlRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "dumbell_curl")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, dbCurlRec)

	bpRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "bench_press")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, bpRec)

	hipThrustRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "hip_thrust")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, hipThrustRec)

	sqRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "squat")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, sqRec)

	shoulPressRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "shoulder_press")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, shoulPressRec)

	cablePressDownRec, err := tools.FreeWeightRecordGet(ctx, db, userID, "cable_press_down")
	if err != nil {
		return nil, nil, err
	}
	free = append(free, cablePressDownRec)

	fmt.Println(free, "free")
	fmt.Println(body, "body")

	return body, free, nil
}
