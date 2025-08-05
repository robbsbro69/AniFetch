package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type SystemInfo struct {
	OS       string
	Kernel   string
	Hostname string
	Uptime   string
	Packages string
	Shell    string
	CPU      string
	Memory   string
	Disk     string
}

func GetSystemInfo() SystemInfo {
	info := SystemInfo{}

	// OS
	info.OS = runtime.GOOS

	// Kernel
	if kernel, err := exec.Command("uname", "-r").Output(); err == nil {
		info.Kernel = strings.TrimSpace(string(kernel))
	}

	// Hostname
	if hostname, err := os.Hostname(); err == nil {
		info.Hostname = hostname
	}

	// Uptime
	if uptime, err := exec.Command("uptime", "-p").Output(); err == nil {
		info.Uptime = strings.TrimSpace(string(uptime))
	}

	// Packages (try different package managers)
	info.Packages = getPackageCount()

	// Shell
	if shell := os.Getenv("SHELL"); shell != "" {
		info.Shell = filepath.Base(shell)
	}

	// CPU
	info.CPU = getCPUInfo()

	// Memory
	info.Memory = getMemoryInfo()

	// Disk
	info.Disk = getDiskInfo()

	return info
}

func getPackageCount() string {
	// Try different package managers
	packageManagers := []struct {
		cmd  string
		args []string
		desc string
	}{
		{"pacman", []string{"-Qq"}, "pacman"},
		{"dpkg", []string{"-l"}, "dpkg"},
		{"rpm", []string{"-qa"}, "rpm"},
		{"brew", []string{"list"}, "brew"},
		{"nix-store", []string{"-qR"}, "nix"},
	}

	for _, pm := range packageManagers {
		cmd := exec.Command(pm.cmd, pm.args...)
		if output, err := cmd.Output(); err == nil {
			lines := strings.Split(strings.TrimSpace(string(output)), "\n")
			if len(lines) > 0 && lines[0] != "" {
				return fmt.Sprintf("%d (%s)", len(lines), pm.desc)
			}
		}
	}

	return "unknown"
}

func getCPUInfo() string {
	// Try to read from /proc/cpuinfo on Linux
	if runtime.GOOS == "linux" {
		if data, err := os.ReadFile("/proc/cpuinfo"); err == nil {
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				if strings.HasPrefix(line, "model name") {
					parts := strings.Split(line, ":")
					if len(parts) > 1 {
						return strings.TrimSpace(parts[1])
					}
				}
			}
		}
	}

	// Fallback for other OS
	return "Unknown CPU"
}

func getMemoryInfo() string {
	if runtime.GOOS == "linux" {
		if data, err := os.ReadFile("/proc/meminfo"); err == nil {
			lines := strings.Split(string(data), "\n")
			var total, available uint64
			for _, line := range lines {
				if strings.HasPrefix(line, "MemTotal:") {
					fmt.Sscanf(line, "MemTotal: %d", &total)
				}
				if strings.HasPrefix(line, "MemAvailable:") {
					fmt.Sscanf(line, "MemAvailable: %d", &available)
				}
			}
			if total > 0 {
				used := total - available
				return fmt.Sprintf("%dMiB / %dMiB", used/1024, total/1024)
			}
		}
	}
	return "Unknown"
}

func getDiskInfo() string {
	if runtime.GOOS == "linux" {
		if output, err := exec.Command("df", "-h", "/").Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			if len(lines) > 1 {
				fields := strings.Fields(lines[1])
				if len(fields) >= 5 {
					return fmt.Sprintf("%s / %s", fields[2], fields[1])
				}
			}
		}
	}
	return "Unknown"
} 