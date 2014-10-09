package libproc

import (
    "os"
    "fmt"
)

func ProcPath(pid int) (string, error) {
    return os.Readlink(fmt.Sprintf("/proc/%d/exe", pid))
}
