package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/texttospeech/apiv1"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

const (
	text   = "こんにちは"
	output = "outputs/sample.mp3"
)

func main() {
	ctx := context.Background()

	client, err := texttospeech.NewClient(
		ctx, option.WithCredentialsFile(os.Getenv("GOOGLE_CREDENTIAL")),
	)
	if err != nil {
		log.Fatalf("create client failed: %v", err)
	}
	defer client.Close()

	resp, err := client.SynthesizeSpeech(ctx, &pb.SynthesizeSpeechRequest{
		Input: &pb.SynthesisInput{
			InputSource: &pb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &pb.VoiceSelectionParams{
			LanguageCode: "ja-JP",
			SsmlGender:   pb.SsmlVoiceGender_FEMALE,
		},
		AudioConfig: &pb.AudioConfig{
			AudioEncoding: pb.AudioEncoding_MP3,
		},
	})
	if err != nil {
		log.Fatalf("synthesized text failed: %v", err)
	}

	if err := ioutil.WriteFile(output, resp.AudioContent, 0644); err != nil {
		log.Fatalf("cloud't write audio to file: %v", err)
	}
}
