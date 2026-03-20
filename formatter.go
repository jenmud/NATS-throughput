package main

import "fmt"

func FormatSize(s int64) string {
	const (
		KiB = 1024
		MiB = KiB * 1024
		GiB = MiB * 1024
		TiB = GiB * 1024
	)

	switch {
	case s >= TiB:
		return fmt.Sprintf("%.2f TiB", float64(s)/float64(TiB))
	case s >= GiB:
		return fmt.Sprintf("%.2f GiB", float64(s)/float64(GiB))
	case s >= MiB:
		return fmt.Sprintf("%.2f MiB", float64(s)/float64(MiB))
	case s >= KiB:
		return fmt.Sprintf("%.2f KiB", float64(s)/float64(KiB))
	default:
		return fmt.Sprintf("%d B", s)
	}
}
