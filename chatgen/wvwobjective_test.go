package chatgen

import (
	"strconv"
	"strings"
	"testing"
)

func TestEncodeWvWObjective(t *testing.T) {
	cases := []struct {
		given struct {
			objectiveID int
			mapID       int
		}
		expect string
	}{
		{
			given: struct {
				objectiveID int
				mapID       int
			}{
				objectiveID: 6,
				mapID:       38,
			},
			expect: "[&DAYAAAAmAAAA]",
		},
	}
	for _, c := range cases {
		result := EncodeWvWObjective(c.given.objectiveID, c.given.mapID)

		if result != c.expect {
			t.Errorf("EncodeWvWObjective returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}

func TestDecodeWvWObjective(t *testing.T) {
	cases := []struct {
		given  string
		expect struct {
			objectiveID int
			mapID       int
			err         error
		}
	}{
		{
			given: "[&DAYAAAAmAAAA]",
			expect: struct {
				objectiveID int
				mapID       int
				err         error
			}{
				objectiveID: 6,
				mapID:       38,
			},
		},
	}

	for _, c := range cases {
		objectiveID, mapID, err := DecodeWvWObjective(c.given)

		if err != c.expect.err {
			t.Errorf("DecodeWvWobjective returned an unexpected error: '%s'", err)
		}
		if objectiveID != c.expect.objectiveID || mapID != c.expect.mapID {
			t.Errorf("DecodeWvWObjective returned '%s' but '%s' was expected", strings.Join([]string{strconv.Itoa(objectiveID), strconv.Itoa(mapID)}, ","), strings.Join([]string{strconv.Itoa(c.expect.objectiveID), strconv.Itoa(c.expect.mapID)}, ","))
		}
	}
}
