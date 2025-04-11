package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateJobID() string {
	return "job-" + strings.Replace(uuid.New().String(), "-", "", -1) + "-" + fmt.Sprintf("%d", time.Now().Unix())
}

func GenerateTaskID() string {
	return "task-" + strings.Replace(uuid.New().String(), "-", "", -1) + "-" + fmt.Sprintf("%d", time.Now().Unix())
}
