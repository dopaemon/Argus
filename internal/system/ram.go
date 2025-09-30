package system

import (
	"github.com/dopaemon/artus/internal/libutils"
	"github.com/dopaemon/artus/internal/utils"
)

func RAMTotal() string {
	return utils.GetString(libutils.GetTotalRAM)
}

func RAMUsed() string {
	return utils.GetString(libutils.GetUsedRAM)
}

func RAMFree() string {
	return utils.GetString(libutils.GetFreeRAM)
}

func RAMUsage() string {
	return utils.GetString(libutils.GetRAMUsagePercent)
}
