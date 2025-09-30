package system

import (
	"github.com/dopaemon/artus/internal/libutils"
	"github.com/dopaemon/artus/internal/utils"
)

func DiskTotal(path string) string {
	return utils.GetString(func() (string, error) { return libutils.GetDiskTotal(path) })
}

func DiskUsed(path string) string {
	return utils.GetString(func() (string, error) { return libutils.GetDiskUsed(path) })
}

func DiskFree(path string) string {
	return utils.GetString(func() (string, error) { return libutils.GetDiskFree(path) })
}

func DiskUsage(path string) string {
	return utils.GetString(func() (string, error) { return libutils.GetDiskUsagePercent(path) })
}
