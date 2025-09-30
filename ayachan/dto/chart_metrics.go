package dto

type Version struct {
	Version string `json:"version"`
}

type ChartDifficultyStandard struct {
	Difficulty   float64 `json:"difficulty"`
	MaxScreenNPS float64 `json:"max_screen_nps"`
	TotalHPS     float64 `json:"total_hps"`
	TotalNPS     float64 `json:"total_nps"`
}

type ChartDifficultyExtend struct {
	MaxSpeed          int `json:"max_speed"`
	FingerMaxHPS      int `json:"finger_max_hps"`
	FlickNoteInterval int `json:"flick_note_interval"`
	NoteFlickInterval int `json:"note_flick_interval"`
}

type Distribution struct {
	Hit  []int `json:"hit"`
	Note []int `json:"note"`
}

type RegularType uint8

const (
	RegularTypeError RegularType = iota
	RegularTypeRegular
	RegularTypeIrregular
)

type NoteCount struct {
	Single         int `json:"single"`
	Flick          int `json:"flick"`
	DirectionLeft  int `json:"direction_left"`
	DirectionRight int `json:"direction_right"`
	SlideStart     int `json:"slide_start"`
	SlideTick      int `json:"slide_tick"`
	SlideHidden    int `json:"slide_hidden"`
	SlideEnd       int `json:"slide_end"`
	SlideFlick     int `json:"slide_flick"`
}

type ChartMetricsStandard struct {
	BPMHigh       float64      `json:"bpm_high"`
	BPMLow        float64      `json:"bpm_low"`
	MainBPM       float64      `json:"main_bpm"`
	MaxScreenNPS  float64      `json:"max_screen_nps"`
	NoteCount     NoteCount    `json:"note_count"`
	TotalHitNote  int          `json:"total_hit_note"`
	TotalHPS      float64      `json:"total_hps"`
	TotalNote     int          `json:"total_note"`
	TotalNPS      float64      `json:"total_nps"`
	TotalTime     float64      `json:"total_time"`
	Distribution  Distribution `json:"distribution"`
	SPRhythm      bool         `json:"sp_rhythm"`
	Irregular     RegularType  `json:"irregular"`
	IrregularInfo string       `json:"irregular_info"`
}

type ChartMetricsExtend struct {
	LeftPercent       float64 `json:"left_percent"`
	FingerMaxHPS      int     `json:"finger_max_hps"`
	MaxSpeed          float64 `json:"max_speed"`
	FlickNoteInterval int     `json:"flick_note_interval"`
	NoteFlickInterval int     `json:"note_flick_interval"`
}

type ChartMetrics struct {
	Difficulty       ChartDifficultyStandard `json:"difficulty"`
	DifficultyExtend *ChartDifficultyExtend  `json:"difficulty_extend,omitempty"`
	Metrics          ChartMetricsStandard    `json:"metrics"`
	MetricsExtend    *ChartMetricsExtend     `json:"metrics_extend,omitempty"`
}
