package system

import (
	"github.com/dopaemon/artus/internal/libutils"
	"github.com/dopaemon/artus/internal/utils"
)

func Name() string {
	return utils.GetString(libutils.GetCPUName)
}

func LogicalCores() string {
	return utils.GetString(libutils.GetLogicalCores)
}

func PhysicalCores() string {
	return utils.GetString(libutils.GetPhysicalCores)
}

func CPUUsage() string {
	return utils.GetString(libutils.GetCPUUsage)
}
