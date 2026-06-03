package service

import (
	"os"
	"runtime"
)

type SystemService struct{}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) GetHostname() (string, error) {
	return os.Hostname()
}

func (s *SystemService) GetOS() string {
	return runtime.GOOS
}

func (s *SystemService) GetArch() string {
	return runtime.GOARCH
}

func (s *SystemService) GetNumCPU() int {
	return runtime.NumCPU()
}

type SystemInfo struct {
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	NumCPU   int    `json:"numCPU"`
}

func (s *SystemService) GetSystemInfo() (*SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	return &SystemInfo{
		Hostname: hostname,
		OS:       runtime.GOOS,
		Arch:     runtime.GOARCH,
		NumCPU:   runtime.NumCPU(),
	}, nil
}
