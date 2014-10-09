package libproc

// @todo maybe someday tie into libproc.h and use proc_pidpath
func ProcPath(pid int) (string, error) {
    return "/dev/null", nil
}
