package network

import (
	"github.com/dopaemon/artus/internal/libutils"
	"github.com/dopaemon/artus/internal/utils"
)

func PrimaryInterface() string {
	if iface, _ := libutils.GetPrimaryInterface(); iface != "" {
		return iface
	}
	return ""
}

func BytesRecv(iface string) string {
	return utils.GetString(func() (string, error) { return libutils.GetNetBytesRecv(iface) })
}

func BytesSent(iface string) string {
	return utils.GetString(func() (string, error) { return libutils.GetNetBytesSent(iface) })
}

func PacketsRecv(iface string) string {
	return utils.GetString(func() (string, error) { return libutils.GetNetPacketsRecv(iface) })
}

func PacketsSent(iface string) string {
	return utils.GetString(func() (string, error) { return libutils.GetNetPacketsSent(iface) })
}
