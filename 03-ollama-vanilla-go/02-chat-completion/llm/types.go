package llm

type LLM struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Answer struct {
	Model   string  `json:"model"`
	Message Message `json:"message"`
	Done    bool    `json:"done"`
}

type Options struct {
	RepeatLastN int     `json:"repeat_last_n"`
	Temperature float64 `json:"temperature"`
}

type Query struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Options  Options   `json:"options"`
	Stream   bool      `json:"stream"`
}

/*
type Options struct {
	NumKeep            int      `json:"num_keep"`
	Seed               int      `json:"seed"`
	NumPredict         int      `json:"num_predict"`
	TopK               int      `json:"top_k"`
	TopP               float64  `json:"top_p"`
	TfsZ               float64  `json:"tfs_z"`
	TypicalP           float64  `json:"typical_p"`
	RepeatLastN        int      `json:"repeat_last_n"`
	Temperature        float64  `json:"temperature"`
	RepeatPenalty      float64  `json:"repeat_penalty"`
	PresencePenalty    float64  `json:"presence_penalty"`
	FrequencyPenalty   float64  `json:"frequency_penalty"`
	Mirostat           int      `json:"mirostat"`
	MirostatTau        float64  `json:"mirostat_tau"`
	MirostatEta        float64  `json:"mirostat_eta"`
	PenalizeNewline    bool     `json:"penalize_newline"`
	Stop               []string `json:"stop"`
	Numa               bool     `json:"numa"`
	NumCtx             int      `json:"num_ctx"`
	NumBatch           int      `json:"num_batch"`
	NumGqa             int      `json:"num_gqa"`
	NumGpu             int      `json:"num_gpu"`
	MainGpu            int      `json:"main_gpu"`
	LowVram            bool     `json:"low_vram"`
	F16Kv              bool     `json:"f16_kv"`
	VocabOnly          bool     `json:"vocab_only"`
	UseMmap            bool     `json:"use_mmap"`
	UseMlock           bool     `json:"use_mlock"`
	RopeFrequencyBase  float64  `json:"rope_frequency_base"`
	RopeFrequencyScale float64  `json:"rope_frequency_scale"`
	NumThread          int      `json:"num_thread"`
}

type Answer struct {
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
	Message   struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Done               bool  `json:"done"`
	TotalDuration      int64 `json:"total_duration"`
	LoadDuration       int   `json:"load_duration"`
	PromptEvalCount    int   `json:"prompt_eval_count"`
	PromptEvalDuration int   `json:"prompt_eval_duration"`
	EvalCount          int   `json:"eval_count"`
	EvalDuration       int64 `json:"eval_duration"`
}
*/
