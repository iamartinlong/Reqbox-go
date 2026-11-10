package store

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"
)

// seedBenchStore opens a temporary store with collections and requests, so
// that loading can be measured at realistic sizes. The requests carry the
// same JSON blobs a saved request does, since load time decodes them.
func seedBenchStore(b *testing.B, requests, collections int) *Store {
	b.Helper()
	st, err := Open(filepath.Join(b.TempDir(), "bench.db"))
	if err != nil {
		b.Fatalf("open store: %v", err)
	}
	b.Cleanup(func() { st.db.Close() })

	kvs := make([]KV, 5)
	for i := range kvs {
		kvs[i] = KV{ID: fmt.Sprintf("kv%d", i), Key: fmt.Sprintf("key%d", i), Value: fmt.Sprintf("value%d", i), Enabled: true}
	}
	blob, _ := json.Marshal(kvs)
	auth, _ := json.Marshal(Auth{Token: "secret"})
	body := make([]byte, 512)
	for i := range body {
		body[i] = 'x'
	}
	created := formatTime(now())

	tx, err := st.db.Begin()
	if err != nil {
		b.Fatalf("begin: %v", err)
	}
	for c := 0; c < collections; c++ {
		if _, err := tx.Exec(`INSERT INTO collections (id, name, sort_index, created_at, updated_at) VALUES (?,?,?,?,?)`,
			fmt.Sprintf("col-%d", c), fmt.Sprintf("Collection %d", c), c, created, created); err != nil {
			b.Fatalf("insert collection: %v", err)
		}
	}
	for i := 0; i < requests; i++ {
		col := fmt.Sprintf("col-%d", i%collections)
		if _, err := tx.Exec(`INSERT INTO requests
			(id, collection_id, name, method, url, params, headers, body_type, body, form,
			 auth_type, auth, sort_index, expose_slug, exposed, created_at, updated_at)
			VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
			fmt.Sprintf("req-%d", i), col, fmt.Sprintf("Request %d", i), "POST",
			fmt.Sprintf("https://api.example.com/v1/items/%d", i),
			string(blob), string(blob), "raw", string(body), string(blob),
			"bearer", string(auth), i, "", 0, created, created); err != nil {
			b.Fatalf("insert request: %v", err)
		}
	}
	if err := tx.Commit(); err != nil {
		b.Fatalf("commit: %v", err)
	}
	return st
}

func BenchmarkListRequests(b *testing.B) {
	for _, n := range []int{1000, 10000, 50000} {
		b.Run(fmt.Sprintf("requests=%d", n), func(b *testing.B) {
			st := seedBenchStore(b, n, 100)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				reqs, err := st.ListRequests()
				if err != nil {
					b.Fatal(err)
				}
				if len(reqs) != n {
					b.Fatalf("got %d requests, want %d", len(reqs), n)
				}
			}
		})
	}
}

func BenchmarkListRequestSummaries(b *testing.B) {
	for _, n := range []int{1000, 10000, 50000} {
		b.Run(fmt.Sprintf("requests=%d", n), func(b *testing.B) {
			st := seedBenchStore(b, n, 100)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				reqs, err := st.ListRequestSummaries()
				if err != nil {
					b.Fatal(err)
				}
				if len(reqs) != n {
					b.Fatalf("got %d requests, want %d", len(reqs), n)
				}
			}
		})
	}
}

func BenchmarkListCollections(b *testing.B) {
	for _, n := range []int{100, 1000} {
		b.Run(fmt.Sprintf("collections=%d", n), func(b *testing.B) {
			st := seedBenchStore(b, 1000, n)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				cols, err := st.ListCollections()
				if err != nil {
					b.Fatal(err)
				}
				if len(cols) != n {
					b.Fatalf("got %d collections, want %d", len(cols), n)
				}
			}
		})
	}
}

// BenchmarkListHistory measures the history pane's refresh query, which
// refresh() runs alongside the tree load.
func BenchmarkListHistory(b *testing.B) {
	st := seedBenchStore(b, 10000, 100)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := st.ListHistory("", 100); err != nil {
			b.Fatal(err)
		}
	}
}
