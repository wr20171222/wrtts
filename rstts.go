//리드코리아 REST API를 통한 TTS변환 "rest/vtspeech --voice hyeryun --text  "안녕하세요,얼쑤팩토리입니다"  --lang Korean --aformat mp3 --srate 8000"
package tts

import (
	"bytes"
	"io"
	"os/exec"
)

//  화자와 TTS문자열 받아서 처리
type Config struct {
	Speaker string
	Speak   string
	Ip      string
}

// 파일 형식 지정
type Speech struct {
	bytes.Buffer
}

var baseCmd = "/usr/vt/rest/vtspeech"

func Speak(t Config) (*Speech, error) {

	// REST Command 실행
	args := []string{"--voice", t.Speaker, "--text", t.Speak, "--ip", t.Ip, "--lang", "Korean", "--aformat", "mp3", "--srate", "48000"}
	cmd := exec.Command(baseCmd, args...)

	output, _ := cmd.CombinedOutput()

	speech := &Speech{}
	if _, err := io.Copy(&speech.Buffer, bytes.NewReader(output)); err != nil {
		return nil, err
	}
	cmd.Run()
	return speech, nil

}
