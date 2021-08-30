package efile

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMain(m *testing.M)  {
	m.Run()
	RemoveAll("./test")
}

func TestPathExists(t *testing.T) {
	type args struct {
		path string
	}
	type testConfig struct {
		args    args
		want    bool
		wantErr bool
	}
	Convey("test", t, func() {
		Convey("not exist", func() {
			tt := testConfig{
				args: args{
					path: "./test.go",
				},
				wantErr: false,
				want:    false,
			}
			got, err := PathExists(tt.args.path)

			So(err != nil, ShouldEqual, tt.wantErr)
			So(got, ShouldEqual, tt.want)
		})
		Convey("exist", func() {
			tt := testConfig{
				args: args{
					path: "./filepath.go",
				},
				wantErr: false,
				want:    true,
			}
			CreateFile(tt.args.path)

			got, err := PathExists(tt.args.path)

			So(err != nil, ShouldEqual, tt.wantErr)
			So(got, ShouldEqual, tt.want)

			RemoveAll(tt.args.path)
		})
	})
}

func TestCreateDir(t *testing.T) {
	type args struct {
		path string
	}
	type testConfig struct {
		args    args
		wantErr bool
	}
	path := "test/dir"
	Convey("test", t, func() {
		//your mock code...
		Convey("test case1", func() {
			tt := testConfig{
				args: args{
					path: path,
				},
				wantErr: false,
			}
			err := CreateDir(tt.args.path)
			So(err != nil, ShouldEqual, tt.wantErr)
		})
	})
	RemoveAll(path)
}

func TestExistsOrCreateDir(t *testing.T) {
	type args struct {
		path string
	}
	type testConfig struct {
		args    args
		wantErr bool
	}
	Convey("test", t, func() {
		//your mock code...
		Convey("test case1", func() {
			tt := testConfig{
				args: args{
					path: "./test/1",
				},
				wantErr: false,
			}
			err := ExistsOrCreateDir(tt.args.path)
			So(err != nil, ShouldEqual, tt.wantErr)
		})
	})
}

func TestCreateFile(t *testing.T) {
	path := "./test/1.txt"
	Convey("test", t, func() {
		Convey("testdata case1", func() {
			err := CreateFile(path)
			So(err, ShouldBeNil)
		})
		Convey("testdata case2", func() {
			err := CreateFile(path)
			So(err, ShouldBeNil)
		})
	})
	RemoveAll(path)
}

func TestExistsOrCreateFile(t *testing.T) {

	path := "test/1.txt"
	Convey("test", t, func() {
		Convey("not exist file", func() {
			err := ExistsOrCreateFile(path)
			So(err, ShouldBeNil)
		})
		Convey("exist file", func() {
			err := ExistsOrCreateFile(path)
			So(err, ShouldBeNil)
		})
	})
	RemoveAll(path)
}

func TestFileNameWithoutExt(t *testing.T) {
	type args struct {
		fileName string
	}

	Convey("test", t, func() {
		Convey("testdata case1", func() {
			params := args{
				fileName: "name.text",
			}

			got := FileNameWithoutExt(params.fileName)
			So(got, ShouldEqual, "name")
		})
	})
}

func TestWriteToFileEnd(t *testing.T) {
	type args struct {
		fileName string
		content  string
	}
	fileName := "TestAppendToFile.txt"
	Remove(fileName)
	ExistsOrCreateFile(fileName)

	Convey("test", t, func() {
		Convey("testdata case1", func() {
			params := args{
				fileName: fileName,
				content:  "1",
			}

			err := WriteToFileEnd(params.fileName, params.content)
			So(err, ShouldBeNil)
			target, _ := ReadFile(params.fileName)
			So(target, ShouldEqual, "1")
		})

		Convey("testdata case2", func() {
			params := args{
				fileName: fileName,
				content:  "2",
			}

			err := WriteToFileEnd(params.fileName, params.content)
			So(err, ShouldBeNil)

			target, _ := ReadFile(params.fileName)
			So(target, ShouldEqual, "12")
		})
	})
	Remove(fileName)
}

func TestWriteToFile(t *testing.T) {
	type args struct {
		fileName string
		content  string
	}

	Convey("test", t, func() {
		Convey("testdata case1", func() {
			path := "TestWriteToFile.txt"
			_ = ExistsOrCreateFile(path)

			err := WriteToFile(path, "1")
			So(err, ShouldBeNil)

			content, err := ReadFile(path)
			t.Log(content, err)
			So(err, ShouldBeNil)
			So(content, ShouldEqual, "1")

			err = WriteToFile(path, "2")
			So(err, ShouldBeNil)

			content, err = ReadFile(path)
			t.Log(content, err)
			So(err, ShouldBeNil)
			So(content, ShouldEqual, "2")

			Remove(path)
		})
	})
}
