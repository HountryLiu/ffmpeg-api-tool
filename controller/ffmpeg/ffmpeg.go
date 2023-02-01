package ffmpeg

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/HountryLiu/ffmpeg-api-tool/utils"
	"github.com/gin-gonic/gin"
)

// @Tags FFmpeg
// @Summary FFmpeg api
// @Description FFmpeg api
// @Accept multipart/form-data
// @Produce application/json
// @Param arg formData string true "ffmpeg参数"
// @Param file formData file true "file"
// @Success 200 {array} byte "文件内容字节流"
// @Failure 500 {object} object{no=int,data=string,errors=string}
// @Router /api/ffmpeg [get]
func Index(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	defer file.Close()
	if err != nil {
		utils.Error(ctx, utils.InternalServerError, err)
	}
	tmp_dir := os.TempDir()
	tmp_dir = path.Join(tmp_dir, "ffmpeg_api_tool")
	fmt.Println(tmp_dir)
	os.MkdirAll(tmp_dir, os.ModePerm)

	tmp_file, err := ioutil.TempFile(tmp_dir, "tmp_")
	defer os.Remove(tmp_file.Name())
	defer tmp_file.Close()

	written, err := io.Copy(tmp_file, file)
	if written != header.Size {
		utils.Error(ctx, utils.InternalServerError, errors.New("数据读取异常"))
	}

	arg := strings.Split(ctx.Request.FormValue("arg"), " ")
	arg = append([]string{"-i", tmp_file.Name()}, arg...)
	buf, err := shell_ffmpeg(arg...)
	if err != nil {
		utils.Error(ctx, utils.InternalServerError, err)
	}

	ctx.Writer.WriteHeader(http.StatusOK)
	io.Copy(ctx.Writer, buf)
}

func shell_ffmpeg(arg ...string) (*bytes.Buffer, error) {

	cmd := exec.Command("ffmpeg", arg...)

	buf := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)

	cmd.Stdout = buf
	cmd.Stderr = buf2
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return buf, nil
}
