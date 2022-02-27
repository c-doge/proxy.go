package utils

import (
	"os"
	"io"
	"io/ioutil"
	"net/http"
    "archive/zip"
    "path/filepath"
)

func PathExists(path string) bool {
	if len(path) == 0 {
		return false
	}
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func IsDir(path string) bool {
    if len(path) == 0 {
        return false
    }
    fi, err := os.Stat(path)
    if err == nil && fi.IsDir() {
        return true
    }
    return false;
}

func Mkdir(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0777);
		} else {
			return err
		}
	} else if !fi.IsDir() {
		return errors.New("Not Direcotory")
	}
	return err;
}


func ListFile(dirPath string) ([]string, error) {
    files := make([]string, 0, 10)
    dir, err := ioutil.ReadDir(dirPath)
    if err != nil {
        return nil, err
    }
    //PathSep := string(os.PathSeparator)
    for _, file := range dir {
        if !file.IsDir() { // 忽略目录
            files = append(files, file.Name())
        }
    }
    sort.SliceStable(files, func(i, j int) bool {
        return strings.Compare(files[i], files[j]) <=0
    })
    return files, nil
}

func UnzipFile(zipFile, dstPath string) error {
	f, err := os.Open(zipFile)
    defer f.Close()
    if err != nil {
        return err;
    }
    fi, err := f.Stat()
    if err != nil {
        return err; 
    }
    
    err = Mkdir(dstPath)
    if err != nil {
    	return err;
    }
   	err = Unzip(f, fi.Size(), dstPath)
    return err
}

func Unzip(r io.ReaderAt, s int64, dstPath string) error {
    reader, err := zip.NewReader(r, s);
    if err != nil {
        return err
    }
    //defer reader.Close()
    for _, f := range reader.File {
        path := filepath.Join(dstPath, f.Name)
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            dstFile, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, f.Mode())
            if err != nil {
                return err;
            }
            _, err = io.Copy(dstFile, rc);
            if err != nil {
                return err;
            }
        }
    }
    return nil;
}