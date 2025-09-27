package dto

type NoteType string

const (
	NoteTypeSingle      NoteType = "Single"
	NoteTypeSlide       NoteType = "Slide"
	NoteTypeLong        NoteType = "Long"
	NoteTypeBPM         NoteType = "BPM"
	NoteTypeDirectional NoteType = "Directional"
)

type NoteDirection string

const (
	NoteDirectionRight NoteDirection = "Right"
	NoteDirectionLeft  NoteDirection = "Left"
)

type ChartDifficulty int

const (
	ChartDifficultyEasy ChartDifficulty = iota
	ChartDifficultyNormal
	ChartDifficultyHard
	ChartDifficultyExpert
	ChartDifficultySpecial
)

type ChartDifficultyName string

const (
	ChartDifficultyNameEasy    ChartDifficultyName = "easy"
	ChartDifficultyNameNormal  ChartDifficultyName = "normal"
	ChartDifficultyNameHard    ChartDifficultyName = "hard"
	ChartDifficultyNameExpert  ChartDifficultyName = "expert"
	ChartDifficultyNameSpecial ChartDifficultyName = "special"
)

type ChartStats struct {
	Time    float64   // 谱面时长，单位秒
	Notes   int       // 音符总数
	BPMs    []float64 // 谱面包含的 BPM 值列表
	MainBPM float64   // 谱面主要 BPM 值
}
