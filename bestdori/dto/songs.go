package dto

type SongsMetaAll map[string]map[string]map[string][]float64

type SongsAll1Info struct {
	MusicTitle []*string `json:"musicTitle"`
}

type SongsAll1 map[string]SongsAll1Info

type SongsAll5Difficulty struct {
	PlayLevel   int        `json:"playLevel"`
	PublishedAt *[]*string `json:"publishedAt"`
}

type SongsAll5Info struct {
	SongsAll1Info `json:",inline"`
	Tag           string                         `json:"tag"`
	BandId        int                            `json:"bandId"`
	JacketImage   []string                       `json:"jacketImage"`
	PublishedAt   []*string                      `json:"publishedAt"`
	ClosedAt      []*string                      `json:"closedAt"`
	Difficulty    map[string]SongsAll5Difficulty `json:"difficulty"`
}

type SongsAll5 map[string]SongsAll5Info

type SongBPM struct {
	BPM   float64 `json:"bpm"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

type SongsAll7Info struct {
	SongsAll5Info `json:",inline"`
	Length        float64              `json:"length"`
	Notes         map[string]int       `json:"notes"`
	BPM           map[string][]SongBPM `json:"bpm"`
}

type SongsAll7 map[string]SongsAll7Info

type SongsAll8Info struct {
	SongsAll7Info `json:",inline"`
	Ruby          []*string `json:"ruby"`
	Phonetic      []*string `json:"phonetic"`
	Lyricist      []*string `json:"lyricist"`
	Composer      []*string `json:"composer"`
	Arranger      []*string `json:"arranger"`
}

type SongsAll8 map[string]SongsAll8Info

type SongAchievement struct {
	MusicId         int    `json:"musicId"`
	AchievementType string `json:"achievementType"`
	RewardType      string `json:"rewardType"`
	Quantity        int    `json:"quantity"`
}

type SongDifficultyMultiLiveScoreMap struct {
	MusicId                 int    `json:"musicId"`
	MusicDifficulty         string `json:"musicDifficulty"`
	MultiLiveDifficultyId   int    `json:"multiLiveDifficultyId"`
	ScoreS                  int    `json:"scoreS"`
	ScoreA                  int    `json:"scoreA"`
	ScoreB                  int    `json:"scoreB"`
	ScoreC                  int    `json:"scoreC"`
	MultiLiveDifficultyType string `json:"multiLiveDifficultyType"`
	ScoreSS                 int    `json:"scoreSS"`
	ScoreSSS                int    `json:"scoreSSS"`
}

type SongDifficulty struct {
	PlayLevel         int                                        `json:"playLevel"`
	MultiLiveScoreMap map[string]SongDifficultyMultiLiveScoreMap `json:"multiLiveScoreMap"`
	NotesQuantity     int                                        `json:"notesQuantity"`
	ScoreC            int                                        `json:"scoreC"`
	ScoreB            int                                        `json:"scoreB"`
	ScoreA            int                                        `json:"scoreA"`
	ScoreS            int                                        `json:"scoreS"`
	ScoreSS           int                                        `json:"scoreSS"`
	PublishedAt       *[]*string                                 `json:"publishedAt"`
}

type SongInfo struct {
	SongsAll8Info `json:",inline"`
	BGMId         string                    `json:"bgmId"`
	BGMFile       string                    `json:"bgmFile"`
	Achievements  []SongAchievement         `json:"achievements"`
	Seq           int                       `json:"seq"`
	HowToGet      []*string                 `json:"howToGet"`
	Description   []*string                 `json:"description"`
	Difficulty    map[string]SongDifficulty `json:"difficulty"`
}
