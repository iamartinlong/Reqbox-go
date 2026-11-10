package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iamartinlong/Reqbox-go/pkg/store"
)

// benchApp builds the in-memory workspace refresh() would load: request
// summaries spread over collections, as a large workspace has them.
func benchApp(requests, collections int) *app {
	a := &app{collapsed: map[string]bool{}}
	a.cols = make([]store.Collection, collections)
	for i := range a.cols {
		a.cols[i] = store.Collection{ID: fmt.Sprintf("col-%d", i), Name: fmt.Sprintf("Collection %d", i)}
	}
	a.reqs = make([]store.RequestSummary, requests)
	for i := range a.reqs {
		id := a.cols[i%collections].ID
		a.reqs[i] = store.RequestSummary{
			ID:           fmt.Sprintf("req-%d", i),
			CollectionID: &id,
			Name:         fmt.Sprintf("Request %d", i),
			Method:       "GET",
			URL:          fmt.Sprintf("https://api.example.com/v1/items/%d", i),
		}
	}
	a.colCounts = make(map[string]int, collections)
	a.reqSearch = make([]reqSearch, requests)
	for i, r := range a.reqs {
		a.colCounts[*r.CollectionID]++
		a.reqSearch[i] = reqSearch{name: strings.ToLower(r.Name), url: strings.ToLower(r.URL)}
	}
	return a
}

func BenchmarkRebuildRows(b *testing.B) {
	for _, n := range []int{1000, 10000, 50000} {
		b.Run(fmt.Sprintf("requests=%d", n), func(b *testing.B) {
			a := benchApp(n, 100)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a.rebuildRows()
				if len(a.rows) != n+len(a.cols)+1 {
					b.Fatalf("got %d rows, want %d", len(a.rows), n+len(a.cols)+1)
				}
			}
		})
	}
}

func BenchmarkRebuildRowsFiltered(b *testing.B) {
	for _, tc := range []struct {
		name   string
		filter string
	}{
		{"match", "Request 99"},
		{"no-match", "zzz"},
	} {
		b.Run(fmt.Sprintf("%s/requests=10000", tc.name), func(b *testing.B) {
			a := benchApp(10000, 100)
			a.filter = tc.filter
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				a.rebuildRows()
			}
		})
	}
}

// BenchmarkCollectionCounts measures the per-load counting refresh() does
// for the header counts, in place of a scan per header and frame.
func BenchmarkCollectionCounts(b *testing.B) {
	a := benchApp(10000, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counts := make(map[string]int, len(a.cols))
		for _, r := range a.reqs {
			counts[*r.CollectionID]++
		}
		_ = counts
	}
}

// BenchmarkHeaderCountsPerFrame measures a frame's header count lookups: a
// map lookup each, after the per-load count.
func BenchmarkHeaderCountsPerFrame(b *testing.B) {
	a := benchApp(10000, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, c := range a.cols {
			_ = a.colCounts[c.ID]
		}
	}
}
