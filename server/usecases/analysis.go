package usecases

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ryoyama0508/WorkoutRecorder/WorkoutRecorder/server/tools"
)

//Analysis ...
func Analysis(ctx context.Context, db *sql.DB, userID string) {
	//excersice funcs
	err := tools.BodyWeightRecordGet(ctx, db, userID, "bench_press")
	if err != nil {
		fmt.Println(err)
	}
}
