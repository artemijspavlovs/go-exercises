package main

import (
	"bytes"
	"testing"
)

func TestLogParser(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		{
			name: "Valid string",
			args: args{
				s: "2023-03-30T10:00:00Z INFO main: Application started",
			},
			wantWriter: "INFO main 2023-03-30T10:00:00Z: Application started\n",
			wantErr:    false,
		},
		{
			name: "Invalid string, no space after separator",
			args: args{
				s: "2023-03-30T10:00:00Z INFO main:Application started",
			},
			wantWriter: "",
			wantErr:    true,
		},
		{
			name: "Invalid time format",
			args: args{
				s: "2023:03:30T10:00:00 INFO main: Application started",
			},
			wantWriter: "",
			wantErr:    true,
		},
		{
			name: "Invalid string",
			args: args{
				s: "I wrote this awesome log entry without reading dev guidelines",
			},
			wantWriter: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := LogParser(tt.args.s, writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogParser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("LogParser()\ngotWriter\n %v\nwant\n %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
