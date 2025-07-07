package minfo

import (
	"fmt"
	"sort"
	"strings"
	"time"

	json "github.com/cbsinteractive/mc-json"
)

type Track interface {
	Info() Header
}

type File struct {
	Path string `json:",omitempty"`

	Header
	Codec

	AudioCount,
	FrameCount,
	VideoCount,
	DataSize,
	FileSize,
	FooterSize,
	HeaderSize int `json:",string,omitempty"`

	EncodedApplication string `json:"Encoded_Application,omitempty"`

	OverallBitRate     float64 `json:",omitempty,string"`
	OverallBitRateMode string  `json:"OverallBitRate_Mode,omitempty"`
	IsStreamable       string  `json:"IsStreamable,omitempty"`

	Track []Track `json:",omitempty"`
}

func (f File) Duration() time.Duration {
	d, _ := time.ParseDuration(fmt.Sprint(f.Header.Duration) + "s")
	return d
}

func (f File) Video() (track []Video) {
	for _, v := range f.Track {
		v, ok := v.(Video)
		if ok {
			track = append(track, v)
		}
	}
	return
}

func (f File) Audio() (track []Audio) {
	for _, v := range f.Track {
		v, ok := v.(Audio)
		if ok {
			track = append(track, v)
		}
	}
	return
}

func (f File) String() string {
	track := f.Track
	f.Track = nil
	data, _ := json.Marshal(f)
	if len(data) == 0 {
		return ""
	}
	s := string(data[:len(data)-1])

	sep := ",\n\t\"Track\": [\n\t\t"
	for _, v := range track {
		data, _ := json.Marshal(v)
		s += sep + string(data)
		sep = ",\n\t\t"
	}
	if len(track) > 0 {
		s += "]"
	}
	s += "}"
	return s
}

func (f *File) Decode(p []byte) error {
	tmp := struct {
		Media struct {
			Ref   string `json:"@ref"`
			Track []json.RawMessage
		}
	}{}
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	order := map[int]string{}
	for _, v := range tmp.Media.Track {
		hdr := Header{}
		json.Unmarshal([]byte(v), &hdr)
		var track interface {
			val() Track
		}
		switch strings.ToLower(hdr.Type) {
		case "general":
			type file File

			json.Unmarshal([]byte(v), (*file)(f))
			f.Header.Type = ""
			f.Path = tmp.Media.Ref
			continue
		case "video":
			track = &Video{}
		case "audio":
			track = &Audio{}
		case "text":
			track = &Text{}
		case "other":
			track = &Timecode{}
		}
		json.Unmarshal([]byte(v), track)
		if track != nil {
			order[len(f.Track)] = hdr.StreamOrder
			f.Track = append(f.Track, track.val())
		}
	}

	sort.SliceStable(f.Track, func(i, j int) bool {
		ii := order[i]
		jj := order[j]
		if ii == "" || jj == "" {
			return i < j
		}
		return ii < jj
	})
	return nil
}

type Header struct {
	Type string `json:"@type,omitempty"`
	ID,
	StreamSize int `json:",omitempty,string"`
	StreamOrder string  `json:",omitempty"`
	Duration    float64 `json:",omitempty,string"`
}

func (h Header) Info() Header { return h }

type Box struct {
	Header
	Codec

	AudioCount,
	FrameCount,
	VideoCount,
	DataSize,
	FileSize,
	FooterSize,
	HeaderSize int `json:",string,omitempty"`

	EncodedApplication string `json:"Encoded_Application,omitempty"`

	OverallBitRate     float64 `json:",omitempty,string"`
	OverallBitRateMode string  `json:"OverallBitRate_Mode,omitempty"`
	IsStreamable       string  `json:"IsStreamable,omitempty"`
}

type Video struct {
	Header
	Frame
	Codec

	Width, Height int `json:",omitempty,string"`

	DisplayAspectRatio,
	PixelAspectRatio,
	Rotation,
	BitDepth float64 `json:",omitempty,string"`

	ScanType,
	ColorSpace,
	ChromaSubsampling string `json:",omitempty"`

	Extra struct {
		CodecConfigurationBox string `json:"Codec_configuration_box,omitempty"`
	} `json:"extra,omitempty"`
}

type Audio struct {
	Header
	Frame
	Codec

	Samples         float64 `json:"SamplingCount,omitempty,string"`
	SampleRate      float64 `json:"SamplingRate,omitempty,string"`
	SamplesPerFrame float64 `json:"SamplesPerFrame,omitempty,string"`

	Channels  string `json:",omitempty"`
	Layout    string `json:"ChannelLayout,omitempty"`
	Positions string `json:"ChannelPositions,omitempty"`

	CompressionMode string `json:"Compression_Mode,omitempty"`
	AlternateGroup  string `json:"AlternateGroup,omitempty"`
	Default         string `json:",omitempty"`
}

type Text struct {
	Header
	Frame
	Codec

	Width, Height   int     `json:",omitempty,string"`
	CompressionMode string  `json:"Compression_Mode,omitempty"`
	Language        string  `json:",omitempty"`
	Delay           float64 `json:",omitempty,string"`
}

type Timecode struct {
	Header
	Frame
	Codec

	FirstFrame string  `json:"TimeCode_FirstFrame,omitempty"`
	Settings   string  `json:"TimeCode_Settings,omitempty"`
	Delay      float64 `json:",omitempty,string"`
}

func (Box) track() string      { return "general" }
func (Video) track() string    { return "video" }
func (Audio) track() string    { return "audio" }
func (Text) track() string     { return "text" }
func (Timecode) track() string { return "timecode" }

func (p *Box) val() Track      { return *p }
func (p *Video) val() Track    { return *p }
func (p *Audio) val() Track    { return *p }
func (p *Text) val() Track     { return *p }
func (p *Timecode) val() Track { return *p }

type Frame struct {
	Count            int     `json:"FrameCount,omitempty,string"`
	Rate             float64 `json:"FrameRate,omitempty,string"`
	RateMode         string  `json:"FrameRate_Mode,omitempty"`
	RateModeOriginal string  `json:"FrameRate_Mode_Original,omitempty"`
}

type Codec struct {
	ID   string `json:"CodecID,omitempty"`
	Name string `json:"CodecID_Compatible,omitempty"`

	Format    string `json:"Format,omitempty"`
	Profile   string `json:"Format_Profile,omitempty"`
	Level     string `json:"Format_Level,omitempty"`
	RefFrames string `json:"Format_Settings_RefFrames,omitempty"`
	CABAC     string `json:"Format_Settings_CABAC,omitempty"`
	Features  string `json:"Format_AdditionalFeatures,omitempty"`

	BitRate,
	BitRateMaximum float64 `json:",string,omitempty"`
	BitRateMode string `json:"BitRate_Mode,omitempty"`
}
