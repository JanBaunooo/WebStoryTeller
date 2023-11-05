package backend

import (
	"github.com/google/uuid"
)

type StoryPack struct {
	PackType         string      `json:"type"` //FAH,FLAM,AUDIO
	Title            string      `json:"title"`
	ModificationDate int         `json:"modificationDate"`
	Description      string      `json:"description"`
	Image            []Image     `json:"image"`
	StageNodes       []StageNode `json:"stageNodes"`
	ListNodes        []ListNode  `json:"actionNodes"`
	AgeMax           int         `json:"ageMax"`
	AgeMin           int         `json:"ageMin"`
	Author           string      `json:"author"`
	Duration         int         `json:"duration"`
	Keywords         []string    `json:"keywords"`
	Locales          string      `json:"locales"` //Unicode CLDR locale identifier
	Source           string      `json:"source"`
	Uuid             uuid.UUID
}

type StageNode struct {
	Uuid            uuid.UUID        `json:"uuid"`
	Type            string           `json:"type"` //enum ?
	Name            string           `json:"name"`
	Image           []Image          `json:"image"`
	Audio           string           `json:"audio"`
	OkTransition    *Transition      `json:"okTransition"`
	HomeTransition  *Transition      `json:"homeTransition"`
	ControlSettings *ControlSettings `json:"controlSettings"`
	SquareOne       bool             `json:"squareOne"`
}

type ListNode struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Options []uuid.UUID `json:"options"`
}

type ControlSettings struct {
	Wheel    bool `json:"wheel"`
	Ok       bool `json:"ok"`
	Home     bool `json:"home"`
	Pause    bool `json:"pause"`
	Autoplay bool `json:"autoplay"`
}

type Transition struct {
	ActionNode  string `json:"actionNode"`
	OptionIndex int    `json:"optionIndex"`
}

type Image struct {
	LowRes string `json:"lowres"`
	HiRes  string `json:"hires"`
}

// Product data
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

// Order data
type Order struct {
	Lines []Line `json:"lines"`
}

// Line data in an order
type Line struct {
	ProductID   int    `json:"productid"`
	ProductName string `json:"productname"`
	Quantity    int    `json:"quantity"`
}
