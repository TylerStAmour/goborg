package goborg

const (
	AppendOnly     = "--append-only"
	MakeParentDirs = "--make-parent-dirs"
)

func StorageQuota(quota string) string {
	return string("--storage-quota=" + quota)
}
