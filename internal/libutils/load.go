package libutils

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/load"
)

func GetLoad1() (string, error) {
	avg, err := load.Avg()
	if err != nil {
		return "", fmt.Errorf("failed to get load average: %v", err)
	}
	return fmt.Sprintf("%.2f", avg.Load1), nil
}

func GetLoad5() (string, error) {
	avg, err := load.Avg()
	if err != nil {
		return "", fmt.Errorf("failed to get load average: %v", err)
	}
	return fmt.Sprintf("%.2f", avg.Load5), nil
}

func GetLoad15() (string, error) {
	avg, err := load.Avg()
	if err != nil {
		return "", fmt.Errorf("failed to get load average: %v", err)
	}
	return fmt.Sprintf("%.2f", avg.Load15), nil
}

func GetMiscLoad() (string, error) {
	misc, err := load.Misc()
	if err != nil {
		return "", fmt.Errorf("failed to get misc load: %v", err)
	}
	return fmt.Sprintf("ProcsRunning: %d, ProcsBlocked: %d, Ctxt: %d",
		misc.ProcsRunning, misc.ProcsBlocked, misc.Ctxt), nil
}
