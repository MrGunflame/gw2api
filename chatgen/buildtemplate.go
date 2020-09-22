package chatgen

import (
	"encoding/base64"
)

// IDs for all professions
const (
	BuildTemplateProfessionGuardian = iota + 1
	BuildTemplateProfessionWarrior
	BuildTemplateProfessionEngineer
	BuildTemplateProfessionRanger
	BuildTemplateProfessionThief
	BuildTemplateProfessionElementalist
	BuildTemplateProfessionMesmer
	BuildTemplateProfessionNecromancer
	BuildTemplateProfessionRevenant
)

// Revenant Legends
const (
	RevenantLegendDragonStance = iota + 14
	RevenantLegendAssassinStance
	RevenantLegendDwarfStance
	RevenantLegendDemonStance
	RevenantLegendRenegadeStance
	RevenantLegendCentaurStance
)

// BuildTemplate contains data contained in a buildtemplate chatlink
type BuildTemplate struct {
	Profession                   uint8
	Specializations              [3]BuildTemplateSpecialization
	Skills                       [5]int
	AquaticSkills                [5]int
	RangerPets                   [4]uint8 // Terrestial Pet 1, Terrestial Pet 2, Aquatic Pet 1, Aquatic Pet 2
	RevenantLegends              [4]uint8 // Terrestial Legend 1, Terrestial Legend 2, Aquatic Legend 1, Aquatic Legend 2
	RevenantInactiveLegendSkills [6]int   // Terrestial Skill 1, 2, 3, Aquatic Skill 1, 2, 3
}

// BuildTemplateSpecialization contains information about one traitline
type BuildTemplateSpecialization struct {
	ID     int
	Traits [3]int
}

// EncodeBuildTemplate encodes a buildtemplate into a chatlink
func EncodeBuildTemplate(data BuildTemplate) string {
	bytes := []byte{headerByteBuildTemplate, byte(data.Profession)}

	// Specializations
	// 2 bytes per spec, the first contains the id
	// the second contains 2 bits for each trait choice in reverse order
	for _, spec := range data.Specializations {
		bytes = append(bytes, byte(spec.ID))

		var b byte
		for i, t := range spec.Traits {
			b += byte(t << (2 * i))
		}
		bytes = append(bytes, b)
	}

	// Skills
	// 2 bytes per skill, terrestial and aquatic skill alternating
	for i := range data.Skills {
		for j := 0; j < 2; j++ {
			bytes = append(bytes, byte(data.Skills[i]>>(8*j)))
		}

		for j := 0; j < 2; j++ {
			bytes = append(bytes, byte(data.AquaticSkills[i]>>(8*j)))
		}
	}

	// Special profession data
	if data.Profession == BuildTemplateProfessionRanger {
		// Ranger pets as single byte
		for _, p := range data.RangerPets {
			bytes = append(bytes, byte(p))
		}

		// 12 bytes unused
		for i := 0; i < 12; i++ {
			bytes = append(bytes, 0)
		}
	} else if data.Profession == BuildTemplateProfessionRevenant {
		// Revenant legends as single byte
		for _, l := range data.RevenantLegends {
			bytes = append(bytes, byte(l))
		}

		// Revenant inactive legend skills as 2 bytes
		// 3 Terrestial and 3 aquatic in this order
		for _, skill := range data.RevenantInactiveLegendSkills {
			for i := 0; i < 2; i++ {
				bytes = append(bytes, byte(skill>>(8*i)))
			}
		}
	} else {
		// 16 bytes unused for all other professions
		for i := 0; i < 16; i++ {
			bytes = append(bytes, 0)
		}
	}

	return addLinkEnclosing(base64.StdEncoding.EncodeToString(bytes))
}

// DecodeBuildTemplate decodes a chatlink into the contained buildtemplate
func DecodeBuildTemplate(link string) (*BuildTemplate, error) {
	bytes, err := base64.StdEncoding.DecodeString(removeLinkEnclosing(link))
	if err != nil {
		return nil, err
	}

	if len(bytes) < 44 || bytes[0] != headerByteBuildTemplate {
		return nil, ErrInvalidHeader
	}

	data := &BuildTemplate{
		Profession: uint8(bytes[1]),
	}

	for i := 0; i < 3; i++ {
		data.Specializations[i].ID = int(bytes[(2*i)+3])
		for j := 0; j < 3; j++ {
			data.Specializations[i].Traits[j] = int(bytes[(2*i)+4]) & (3 << (j * 2))
		}
	}

	// TODO: Skills

	return data, nil
}
