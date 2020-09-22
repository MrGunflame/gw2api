package chatgen

import (
	"testing"
)

func TestEncodeBuildTemplate(t *testing.T) {
	cases := []struct {
		given  BuildTemplate
		expect string
	}{
		{
			given: BuildTemplate{
				Profession: BuildTemplateProfessionEngineer,
				Specializations: [3]BuildTemplateSpecialization{
					{
						ID:     6,
						Traits: [3]int{3, 1, 3},
					},
					{
						ID:     29,
						Traits: [3]int{2, 2, 2},
					},
					{
						ID:     21,
						Traits: [3]int{1, 2, 1},
					},
				},
				Skills:        [5]int{132, 394, 292, 405, 393},
				AquaticSkills: [5]int{0, 0, 0, 0, 0},
			},
			expect: "[&DQMGNx0qFRmEAAAAigEAACQBAACVAQAAiQEAAAAAAAAAAAAAAAAAAAAAAAA=]",
		},
		{
			given: BuildTemplate{
				Profession: BuildTemplateProfessionRanger,
				RangerPets: [4]uint8{31, 44, 0, 0},
			},
			expect: "[&DQQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAB8sAAAAAAAAAAAAAAAAAAA=]",
		},
		// {
		// 	given: BuildTemplate{
		// 		Profession:                   BuildTemplateProfessionRevenant,
		// 		Skills:                       [5]int{4572, 4651, 4614, 4564, 4554},
		// 		RevenantLegends:              [4]uint8{RevenantLegendDwarfStance, RevenantLegendAssassinStance, 0, 0},
		// 		RevenantInactiveLegendSkills: [6]int{4564, 4651, 4614, 0, 0, 0},
		// 	},
		// 	expect: "[&DQkAAAAAAADcEQAA1BEAACsSAAAGEgAAyhEAAA4PAAAAAAAAAAAAAAAAAAA=]",
		// },
	}

	for _, c := range cases {
		result := EncodeBuildTemplate(c.given)

		if result != c.expect {
			t.Errorf("EncodeBuildTemplate returned '%s' but '%s' was expected", result, c.expect)
		}
	}
}
