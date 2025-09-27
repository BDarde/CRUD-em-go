package tests

import (
	"testing"

	"github.com/BDarde/CRUD-em-go/person"
)

func TestServicePerson_Create(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		recieve person.Person
		wantErr bool
	}{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var ps person.ServicePerson
			gotErr := ps.Create(tt.recieve)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Create() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Create() succeeded unexpectedly")
			}
		})
	}
}
