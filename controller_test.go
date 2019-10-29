package main

import "testing"

func Test_nextNDate(t *testing.T) {
	type args struct {
		date string
		n    int
	}
	tests := []struct {
		name    string
		args    args
		wantNd  string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", args{"20191022", 1}, "20191023", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNd, err := nextNDate(tt.args.date, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("nextNDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNd != tt.wantNd {
				t.Errorf("nextNDate() = %v, want %v", gotNd, tt.wantNd)
			}
		})
	}
}
