package main

import (
    "log"
)

//const tool = "/opt/MegaRAID/storcli/storcli64"
const tool = "ls"

type RaidPlugin struct{}

// 获取RAID信息
func (r RaidPlugin) RAID() (string, error) {
    cmd := exec.Command(tool, "-l", "/var/log")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}

// Clear 擦除raid
func (r RaidPlugin) Clear(ctrlID string) (err error) {
    cmd := exec.Command(tool, "-l", "/tmp")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}

// CreateArray 用指定的物理磁盘和RAID级别创建阵列(逻辑磁盘)
func (r *RaidPlugin) CreateArray(ctrlID string, level string, drives []string) (err error) {
    // storcli64 /c0 add vd type=raid10 size=all drives=134:0,134:1 pdperarray=2
    return nil
}

// SetGlobalHotspares 添加热备盘
func (r *RaidPlugin) SetGlobalHotspares(ctrlID string, drives []string) (err error) {
    // /opt/MegaRAID/storcli/storcli64 /cx[/ex]/sx add hotsparedrive [{dgs=<n|0,1,2...>}] [enclaffinity][nonrevertible]
    return nil
}
