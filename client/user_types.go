// API "cellar": Application User Types
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/goadesign/goa-cellar/design
// --out=$(GOPATH)/src/github.com/goadesign/goa-cellar
// --version=v1.1.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// bottlePayload user type.
type bottlePayload struct {
	Color     *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Name      *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `form:"vineyard,omitempty" json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate validates the bottlePayload type instance.
func (ut *bottlePayload) Validate() (err error) {
	if ut.Color != nil {
		if !(*ut.Color == "red" || *ut.Color == "white" || *ut.Color == "rose" || *ut.Color == "yellow" || *ut.Color == "sparkling") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.color`, *ut.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 2, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Review != nil {
		if utf8.RuneCountInString(*ut.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *ut.Review, utf8.RuneCountInString(*ut.Review), 3, true))
		}
	}
	if ut.Review != nil {
		if utf8.RuneCountInString(*ut.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *ut.Review, utf8.RuneCountInString(*ut.Review), 300, false))
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 1, true))
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 5, false))
		}
	}
	if ut.Varietal != nil {
		if utf8.RuneCountInString(*ut.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.varietal`, *ut.Varietal, utf8.RuneCountInString(*ut.Varietal), 4, true))
		}
	}
	if ut.Vineyard != nil {
		if utf8.RuneCountInString(*ut.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.vineyard`, *ut.Vineyard, utf8.RuneCountInString(*ut.Vineyard), 2, true))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 2020, false))
		}
	}
	return
}

// Publicize creates BottlePayload from bottlePayload
func (ut *bottlePayload) Publicize() *BottlePayload {
	var pub BottlePayload
	if ut.Color != nil {
		pub.Color = ut.Color
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	if ut.Region != nil {
		pub.Region = ut.Region
	}
	if ut.Review != nil {
		pub.Review = ut.Review
	}
	if ut.Sweetness != nil {
		pub.Sweetness = ut.Sweetness
	}
	if ut.Varietal != nil {
		pub.Varietal = ut.Varietal
	}
	if ut.Vineyard != nil {
		pub.Vineyard = ut.Vineyard
	}
	if ut.Vintage != nil {
		pub.Vintage = ut.Vintage
	}
	return &pub
}

// BottlePayload user type.
type BottlePayload struct {
	Color     *string `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Country   *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	Name      *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Region    *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	Review    *string `form:"review,omitempty" json:"review,omitempty" xml:"review,omitempty"`
	Sweetness *int    `form:"sweetness,omitempty" json:"sweetness,omitempty" xml:"sweetness,omitempty"`
	Varietal  *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	Vineyard  *string `form:"vineyard,omitempty" json:"vineyard,omitempty" xml:"vineyard,omitempty"`
	Vintage   *int    `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
}

// Validate validates the BottlePayload type instance.
func (ut *BottlePayload) Validate() (err error) {
	if ut.Color != nil {
		if !(*ut.Color == "red" || *ut.Color == "white" || *ut.Color == "rose" || *ut.Color == "yellow" || *ut.Color == "sparkling") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.color`, *ut.Color, []interface{}{"red", "white", "rose", "yellow", "sparkling"}))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 2, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 2, true))
		}
	}
	if ut.Review != nil {
		if utf8.RuneCountInString(*ut.Review) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *ut.Review, utf8.RuneCountInString(*ut.Review), 3, true))
		}
	}
	if ut.Review != nil {
		if utf8.RuneCountInString(*ut.Review) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.review`, *ut.Review, utf8.RuneCountInString(*ut.Review), 300, false))
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 1, true))
		}
	}
	if ut.Sweetness != nil {
		if *ut.Sweetness > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.sweetness`, *ut.Sweetness, 5, false))
		}
	}
	if ut.Varietal != nil {
		if utf8.RuneCountInString(*ut.Varietal) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.varietal`, *ut.Varietal, utf8.RuneCountInString(*ut.Varietal), 4, true))
		}
	}
	if ut.Vineyard != nil {
		if utf8.RuneCountInString(*ut.Vineyard) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.vineyard`, *ut.Vineyard, utf8.RuneCountInString(*ut.Vineyard), 2, true))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 2020, false))
		}
	}
	return
}
