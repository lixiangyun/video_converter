package main

import (
	"log"

	"github.com/lixiangyun/video_converter/ffmpeg/avcodec"
	"github.com/lixiangyun/video_converter/ffmpeg/avdevice"
	"github.com/lixiangyun/video_converter/ffmpeg/avfilter"
	"github.com/lixiangyun/video_converter/ffmpeg/avformat"
	"github.com/lixiangyun/video_converter/ffmpeg/avutil"
	"github.com/lixiangyun/video_converter/ffmpeg/swresample"
	"github.com/lixiangyun/video_converter/ffmpeg/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
