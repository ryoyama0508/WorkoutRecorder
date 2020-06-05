package usecases

import (
	"context"
	"database/sql"

	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//Analysis ...
func Analysis(ctx context.Context, db *sql.DB, userID string) {

	//excersice funcs
	tools.BodyWeightRecordGet(ctx, db, userID, "bench_press")

}
