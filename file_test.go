package minfo

import (
	"testing"
)

func TestMinfo(t *testing.T) {
	f := File{}
	f.Decode([]byte(sample))
	t.Logf("%s\n", f)
}

const sample = `
{
"media": {
"@ref": "test_bbb_360x240_1mb.mp4?raw=true",
"track": [
{
"@type": "General",
"VideoCount": "1",
"AudioCount": "1",
"FileExtension": "mp4?raw=true",
"Format": "MPEG-4",
"Format_Profile": "Base Media",
"CodecID": "isom",
"CodecID_Compatible": "isom/iso2/avc1/mp41",
"FileSize": "1053651",
"Duration": "13.696",
"OverallBitRate_Mode": "VBR",
"OverallBitRate": "615450",
"FrameRate": "15.000",
"FrameCount": "205",
"StreamSize": "6960",
"HeaderSize": "40",
"DataSize": "1046699",
"FooterSize": "6912",
"IsStreamable": "No",
"Encoded_Date": "UTC 1970-01-01 00:00:00",
"Tagged_Date": "UTC 2014-07-19 17:39:07",
"Encoded_Application": "Lavf53.24.2",
"extra": {
"FileExtension_Invalid": "mov mp4 m4v m4a m4b m4p m4r 3ga 3gpa 3gpp 3gp 3gpp2 3g2 k3g jpm jpx mqv ismv isma ismt f4a f4b f4v"
}
},
{
"@type": "Video",
"StreamOrder": "0",
"ID": "1",
"Format": "AVC",
"Format_Profile": "Main",
"Format_Level": "1.3",
"Format_Settings_CABAC": "Yes",
"Format_Settings_RefFrames": "1",
"CodecID": "avc1",
"Duration": "13.667",
"BitRate": "229387",
"Width": "320",
"Height": "240",
"Sampled_Width": "320",
"Sampled_Height": "240",
"PixelAspectRatio": "1.000",
"DisplayAspectRatio": "1.333",
"Rotation": "0.000",
"FrameRate_Mode": "CFR",
"FrameRate_Mode_Original": "VFR",
"FrameRate": "15.000",
"FrameCount": "205",
"ColorSpace": "YUV",
"ChromaSubsampling": "4:2:0",
"BitDepth": "8",
"ScanType": "Progressive",
"StreamSize": "391870",
"Encoded_Date": "UTC 1970-01-01 00:00:00",
"Tagged_Date": "UTC 1970-01-01 00:00:00",
"extra": {
"Codec_configuration_box": "avcC"
}
},
{
"@type": "Audio",
"StreamOrder": "1",
"ID": "2",
"Format": "AAC",
"Format_AdditionalFeatures": "LC",
"CodecID": "mp4a-40-2",
"Duration": "13.696",
"BitRate_Mode": "VBR",
"BitRate": "384000",
"BitRate_Maximum": "416704",
"Channels": "6",
"ChannelPositions": "Front: L C R, Side: L R, LFE",
"ChannelLayout": "C L R Ls Rs LFE",
"SamplesPerFrame": "1024",
"SamplingRate": "48000",
"SamplingCount": "657408",
"FrameRate": "46.875",
"FrameCount": "642",
"Compression_Mode": "Lossy",
"StreamSize": "654821",
"StreamSize_Proportion": "0.62148",
"Default": "Yes",
"AlternateGroup": "1",
"Encoded_Date": "UTC 1970-01-01 00:00:00",
"Tagged_Date": "UTC 1970-01-01 00:00:00"
}
]
}
}
`
