package nessus

import (
	"testing"
)

func TestClient_FoldersDelete(t *testing.T) {
	tests := []struct {
		name    string
		options []Option
		id      int
		wantErr bool
	}{
		{
			name: "success",
			options: []Option{
				WithAPIURL("https://localhost:8834"),
				WithAPIKey(
					"ede55bc4fbad66a41a46f4a5ff35500e7485adae1bc5d94d98e9a1c1f7bb0ecc",
					"af2ca8def3fa67705e38ded764d3c282fb7f82a516883bb4f1e310aba02f1e1b",
				),
			},
			id:      7,
			wantErr: false,
		},
		{
			name: "error",
			options: []Option{
				WithAPIURL("https://localhost:8834"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(tt.options...)
			if err != nil {
				t.Errorf("NewClient() error = %v", err)
			}

			if err := c.FoldersDelete(tt.id); (err != nil) != tt.wantErr {
				t.Errorf("FoldersDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
