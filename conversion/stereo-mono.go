package conversion

import (
	"os"
	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func ConvertToMono(infilePath string) string {

	infile, err := os.Open(infilePath);


	if err!=nil{
		return "error occured " + err.Error()
	}

	defer infile.Close()

	decoder := wav.NewDecoder(infile);

	if !decoder.IsValidFile() {
		return "invalid file"
	}

	pcmBuffer,err := decoder.FullPCMBuffer()

	if err!=nil{
		return "error occured couldnt load the pcm buffer :" + err.Error()
	}

	if pcmBuffer.Format.NumChannels!=2 {
		return "file is not stereo"
	}

	monoBuffer := &audio.IntBuffer{
		Format: &audio.Format{
			NumChannels: 1,
			SampleRate: pcmBuffer.Format.SampleRate,
		},
		Data:           make([]int, len(pcmBuffer.Data)/2),
        SourceBitDepth: pcmBuffer.SourceBitDepth,
	}

	for i := 0; i < len(pcmBuffer.Data); i += 2 {
        left := pcmBuffer.Data[i]
		right := pcmBuffer.Data[i+1]
		mono := (left+right)/2
		monoBuffer.Data[i/2] = mono
    }

	outfile, err := os.Create("output_mono.wav")

	if err != nil {
		return "error creating output file: " + err.Error()
	}
	defer outfile.Close()

	encoder := wav.NewEncoder(outfile, monoBuffer.Format.SampleRate, pcmBuffer.SourceBitDepth, monoBuffer.Format.NumChannels, 1)

	err = encoder.Write(monoBuffer)
	if err != nil {
		return "error writing buffer: " + err.Error()
	}

	err = encoder.Close()
	if err != nil {
		return "error closing encoder: " + err.Error()
	}

	return "successfully saved mono wav"
}
