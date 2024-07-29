//package main
//
//import "testing"
//
//func TestFamily_AddNew(t *testing.T) {
//	t.Run("add family member", func(t *testing.T) {
//		f := Family{}
//
//		if err := f.AddNew(Father, Person{FirstName: "123", LastName: "321", Age: 30}); err != nil || len(f.Members) != 1 {
//			t.Errorf("Should add one family member without errors")
//		}
//	})
//
//	t.Run("add same family members", func(t *testing.T) {
//		f := Family{}
//
//		f.AddNew(Father, Person{FirstName: "123", LastName: "321", Age: 30})
//
//		if err := f.AddNew(Father, Person{FirstName: "abc", LastName: "cba", Age: 3}); err == nil || len(f.Members) == 2 {
//			t.Errorf("Should not add same family and throw error")
//		}
//	})
//}

//package main
//
//import "testing"
//
//func TestFamily_AddNew(t *testing.T) {
//	type newPerson struct {
//		r Relationship
//		p Person
//	}
//	tests := []struct {
//		name           string
//		existedMembers map[Relationship]Person
//		newPerson      newPerson
//		wantErr        bool
//	}{
//		{
//			name: "add father",
//			existedMembers: map[Relationship]Person{
//				Mother: {
//					FirstName: "Maria",
//					LastName:  "Popova",
//					Age:       36,
//				},
//			},
//			newPerson: newPerson{
//				r: Father,
//				p: Person{
//					FirstName: "Misha",
//					LastName:  "Popov",
//					Age:       42,
//				},
//			},
//			wantErr: false,
//		},
//		{
//			name: "catch error",
//			existedMembers: map[Relationship]Person{
//				Father: {
//					FirstName: "Misha",
//					LastName:  "Popov",
//					Age:       42,
//				},
//			},
//			newPerson: newPerson{
//				r: Father,
//				p: Person{
//					FirstName: "Ken",
//					LastName:  "Gymsohn",
//					Age:       32,
//				},
//			},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			f := &Family{
//				Members: tt.existedMembers,
//			}
//			err := f.AddNew(tt.newPerson.r, tt.newPerson.p)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}

			err := f.AddNew(tt.newPerson.r, tt.newPerson.p)

			if !tt.wantErr {
				require.NoError(t, err)
				assert.Contains(t, f.Members, tt.newPerson.r)
				return
			}

			assert.Error(t, err)

			//if (err != nil) != tt.wantErr {
			//	t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
