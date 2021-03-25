package demoArchive

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func write() bytes.Buffer{
	var buf bytes.Buffer
	var tw = tar.NewWriter(&buf)
	var files = []struct{
		Name string
		Body string
	}{
		{
			Name:"readme.txt",
			Body:"this archive contains some text files.",
		},
		{
			Name:"gopher.txt",
			Body:"Gopher names:\nGeorge",
		},
	}
	for _,file := range files {
		hdr := &tar.Header{
			Name : file.Name,
			Mode:0600,
			Uid:0,
			Gid:0,
			Size:int64(len(file.Body)),
			ModTime:time.Now(),
			Typeflag:tar.TypeReg,
			Linkname:"",
			Uname:"",
			Gname:"",
			Devmajor:0,
			Devminor:0,
			AccessTime:time.Now(),
			ChangeTime:time.Now(),
		}
		if err := tw.WriteHeader(hdr);err != nil {
			log.Fatal(err)
		}
		if _,err := tw.Write([]byte(file.Body));err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}

func read(path string) {
	bf,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var readBuf = bytes.NewBuffer(bf)
	tr := tar.NewReader(readBuf)

	for {
		hdr,err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("contents of %s:\n",hdr.Name)
		b,err := ioutil.ReadAll(tr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n",b)
	}
}

func writeZip() bytes.Buffer {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	w.RegisterCompressor(zip.Deflate,func(out io.Writer)(io.WriteCloser,error) {
		return flate.NewWriter(out,flate.BestCompression)
	})
	var files = []struct{
		Name,Body string
	}{
		{"readme.txt","this archive contains some text files."},
		{"todo.txt","Get animal handling licence"},
	}
	for _,file :=range files {
		f,err :=w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_,err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func readZip(path string) {
	r,err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer  r.Close()
	for _,f := range r.File {
		fmt.Printf("Content of %s:\n",f.Name)
		rc,err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		b,err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n",b)
		rc.Close()
	}
}



const (
	filePath = "./testdata/test.tar"
	fileZip = "./testdata/test.zip"
)

func Test(t *testing.T) {
	buf := write()
	if err := ioutil.WriteFile(filePath,buf.Bytes(),os.ModePerm); err != nil {
		log.Fatal(err)
	}
	read(filePath)
}

func TestZip(t *testing.T) {
	buf := writeZip()
	if err := ioutil.WriteFile(fileZip,buf.Bytes(),os.ModePerm); err != nil {
		log.Fatal(err)
	}
	readZip(fileZip)
}