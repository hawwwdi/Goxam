package imageMangement

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

const savePath = "../images/"

func init() {
	if err := os.Mkdir(savePath, os.ModeDir); err != nil {
		log.Printf("create directory %v: FALSE", savePath )
	}
	//todo check fileMode
}

type Image struct {
	id   int64
	path string
}

// pass teacher, exam, question/student name + fileName as pathParts
func NewImage(img *os.File, pathParts ...string) (*Image, error) {
	i := new(Image)
	i.id = currTime()
	i.path = path.Join(pathParts...)
	err := saveImg(img, i.path)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (i *Image) Id() int64 {
	return i.id
}

func (i *Image) Path() string {
	return i.path
}

func saveImg(file *os.File, imgPath string) error {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err1 := ioutil.WriteFile(imgPath, bytes, os.ModeType)
	//todo check above code
	return err1
}

func currTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
