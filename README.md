# minfo
minimal mediainfo wrapper

# usage - package

```
package main

import (
	"context"
	"fmt"

	"github.com/as/minfo"
)

func main() {
	c := context.Background()
	file, err := minfo.ReadURL(c, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4")
	fmt.Println(err)
	fmt.Println(file)
}
```

# usage - command line

```
go get github.com/as/minfo/cmd/minfo
minfo http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4
```

```
{"Path":"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4","StreamSize":"196724","Duration":"653.803","CodecID":"mp42","CodecID_Compatible":"isom/avc1/mp42","Format":"MPEG-4","Format_Profile":"Base Media","AudioCount":"1","FrameCount":"15691","VideoCount":"1","DataSize":"169415646","FileSize":"169612362","HeaderSize":"196716","OverallBitRate":"2075394","OverallBitRate_Mode":"VBR","IsStreamable":"Yes",
	"Track": [
		{"@type":"Video","ID":"2","StreamOrder":"1","StreamSize":"159156325","Duration":"653.792","FrameCount":"15691","FrameRate":"24","FrameRate_Mode":"CFR","CodecID":"avc1","Format":"AVC","Format_Profile":"High","Format_Level":"3.1","Format_Settings_RefFrames":"1","Format_Settings_CABAC":"Yes","BitRate":"1947487","Width":"1280","Height":"720","DisplayAspectRatio":"1.778","PixelAspectRatio":"1","BitDepth":"8","ScanType":"Progressive","ColorSpace":"YUV","ChromaSubsampling":"4:2:0","extra":{"Codec_configuration_box":"avcC"}},
		{"@type":"Audio","ID":"1","StreamSize":"10259313","Duration":"653.803","FrameCount":"28157","FrameRate":"43.066","CodecID":"mp4a-40-2","Format":"AAC","Format_AdditionalFeatures":"LC","BitRate":"125440","BitRate_Mode":"VBR","SamplingCount":"28832712","SamplingRate":"44100","SamplesPerFrame":"1024","Channels":"2","ChannelLayout":"L R","ChannelPositions":"Front: L R","Compression_Mode":"Lossy"}]}
```
